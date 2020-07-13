package main

import (
	"encoding/json"

	"github.com/aws/aws-lambda-go/lambda"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/dynamodb"
	"github.com/aws/aws-sdk-go/service/dynamodb/dynamodbattribute"
	"github.com/rs/zerolog/log"

	cm "alumni-dashboard/api/common"
)

type User struct {
	Email     string `json:"email"`
	FirstName string `json:"firstName"`
	LastName  string `json:"lastName"`
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

	result, err := svc.Scan(&dynamodb.ScanInput{
		TableName:            aws.String("members"),
		ProjectionExpression: aws.String("email ,firstName, lastName"),
	})

	var users []User
	if err = dynamodbattribute.UnmarshalListOfMaps(result.Items, &users); err != nil {
		return cm.CreateResponse(500, "Error unmarshalling dynamo response", err)
	}

	jsonPayload, err := json.Marshal(users)
	if err != nil {
		return cm.CreateResponse(500, "Error json marshalling users", err)
	}

	return cm.CreateResponse(200, string(jsonPayload), nil)
}
