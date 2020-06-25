package main

import (
	"encoding/json"
	"fmt"

	"github.com/aws/aws-lambda-go/lambda"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/dynamodb"
	"github.com/aws/aws-sdk-go/service/dynamodb/dynamodbattribute"
	"github.com/rs/zerolog/log"

	cm "alumni-dashboard/api/common"
)

type RequestBody struct {
	Email string `json:"email"`
}

func main() {
	lambda.Start(RespondLambda)
}

func RespondLambda(request json.RawMessage) (*cm.Response, error) {
	reqBytes, err := cm.ParseRequest(request)
	if err != nil {
		return cm.CreateResponse(400, "Bad Request", err)
	}

	var reqBody RequestBody
	err = json.Unmarshal(reqBytes, &reqBody)
	if err != nil {
		log.Error().Msg("Error unmarshalling request body")
		return cm.CreateResponse(400, "Bad Request Object", err)
	}

	svc, err := cm.GetDBClient()
	if err != nil {
		return cm.CreateResponse(500, "Failed to Create Member", err)
	}

	result, err := svc.GetItem(&dynamodb.GetItemInput{
		TableName: aws.String("members"),
		Key: map[string]*dynamodb.AttributeValue{
			"email": {
				S: aws.String(reqBody.Email),
			},
		},
	})
	if err != nil {
		log.Error().Msg("Error fetching member")
		return cm.CreateResponse(500, "Error fetching member from db", err)
	}

	var userInfo interface{}
	if err = dynamodbattribute.UnmarshalMap(result.Item, &userInfo); err != nil {
		return cm.CreateResponse(500, "Error unmarshalling dynamo response", err)
	}

	jsonPayload, err := json.Marshal(userInfo)
	if err != nil {
		return cm.CreateResponse(500, "Error json marshalling users", err)
	}

	fmt.Println(string(jsonPayload))
	return cm.CreateResponse(200, string(jsonPayload), nil)
}
