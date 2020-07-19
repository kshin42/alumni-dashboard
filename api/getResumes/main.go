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
	Name string
	Link string
}

type ResumeDBItem struct {
	PK string
	SK string
	Link string
}

type Users struct {
	PK        string
	SK        string
	FirstName string
	LastName  string
	Email     string
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

	// Find all users in an org
	var item cm.DBItem
	item.PK = fmt.Sprintf("ORG#ASP")
	item.SK = fmt.Sprintf("USER#")
	results, err := cm.GetItemsWithSKPrefix(item, svc)
	if err != nil {
		log.Error().Msg("Error while finding existing resume link")
		return cm.CreateResponse(500, "Error while trying to find existing resumes", err)
	}

	var users []cm.DBItem
	err = dynamodbattribute.UnmarshalListOfMaps(results.Items, &users)
	if err != nil {
		log.Error().Msg("Error while finding users in org")
		return cm.CreateResponse(500, "Error while trying to find existing resumes", err)
	}

	var resumes []Resume
	for _, u := range users {
		log.Info().Msgf("User Info: %v", u)
		var dbi cm.DBItem
		dbi.PK = u.SK
		dbi.SK = fmt.Sprintf("RESUME#")
		result, err := cm.GetItemsWithSKPrefix(dbi, svc)
		if err != nil {
			log.Error().Msg("Error while looking for resumes")
			return cm.CreateResponse(500, "Error while trying to find existing resumes", err)
		}

		var dbo []ResumeDBItem
		err = dynamodbattribute.UnmarshalListOfMaps(result.Items, &dbo)
		if err != nil {
			log.Error().Msg("Error while unmarshalling user while trying to find resume")
			return cm.CreateResponse(500, "Error while trying to find existing resumes", err)
		}
		if len(dbo) > 0 {
			var r Resume
			r.Name = fmt.Sprintf("%s", u.SK)
			r.Link = dbo[0].Link
			resumes = append(resumes, r)
		}
	}

	if len(resumes) == 0 {
		log.Info().Msg("No resumes found")
		return cm.CreateResponse(200, "", nil)
	}

	r, err := json.Marshal(resumes)
	if err != nil {
		log.Error().Msg("Error while marshaling json response")
		return cm.CreateResponse(500, "Error while trying to find existing resumes", nil)
	}

	return cm.CreateResponse(200, string(r), nil)
}
