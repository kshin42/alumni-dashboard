package main

import (
	"encoding/json"
	"fmt"

	"github.com/aws/aws-lambda-go/lambda"
	"github.com/aws/aws-sdk-go/service/dynamodb/dynamodbattribute"
	"github.com/rs/zerolog/log"

	cm "alumni-dashboard/api/common"
)

type Resume struct {
	PK       string
	SK       string
	Link     string
	Feedback []string
}

func main() {
	lambda.Start(RespondLambda)
}

func RespondLambda(request json.RawMessage) (*cm.Response, error) {
	req, err := cm.ParseRequest(request)
	if err != nil {
		return cm.CreateResponse(400, "Bad Request", err)
	}

	svc, err := cm.GetDBClient()
	if err != nil {
		return cm.CreateResponse(500, "Failed to get members", err)
	}

	err = cm.ValidateSession(req, svc)
	if err != nil {
		log.Error().Msgf("Session is not valid with: %s", err.Error())
		return cm.CreateResponse(401, "Unauthorized Access", nil)
	}

	var item cm.DBItem
	item.PK = fmt.Sprintf("USER#%s", req.Headers["x-user-email"].(string))
	item.SK = fmt.Sprintf("RESUME#")
	results, err := cm.GetItemsWithSKPrefix(item, svc)
	if err != nil {
		log.Error().Msg("Error while finding existing resume link")
		return cm.CreateResponse(500, "Error while trying to find existing resume", err)
	}

	var existingResume []Resume
	err = dynamodbattribute.UnmarshalListOfMaps(results.Items, &existingResume)
	if err != nil {
		log.Error().Msg("Error while unmarshaling query response")
		return cm.CreateResponse(500, "Error while trying to find existing resume", err)
	}

	if len(existingResume) > 1 {
		log.Error().Msg("Too many resume objects found")
		return cm.CreateResponse(500, "Data integrity issue. Please contact administrator", err)
	} else if len(existingResume) == 1 {
		return cm.CreateResponse(200, existingResume[0].Link, nil)
	}

	return cm.CreateResponse(200, "", nil)
}
