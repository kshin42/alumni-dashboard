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

type Member struct {
	fields interface{}
}

func main() {
	lambda.Start(RespondLambda)
}

func RespondLambda(request json.RawMessage) (*cm.Response, error) {
	reqBytes, err := cm.ParseRequest(request)
	if err != nil {
		return cm.CreateResponse(400, "Bad Request", err)
	}

	var member Member
	err = json.Unmarshal(reqBytes, &member.fields)
	if err != nil {
		log.Error().Msg("Error unmarshalling request body")
		return cm.CreateResponse(400, "Bad Member Object", err)
	}

	svc, err := cm.GetDBClient()
	if err != nil {
		return cm.CreateResponse(500, "Failed to Create Member", err)
	}

	av, err := dynamodbattribute.MarshalMap(member.fields)
	if err != nil {
		log.Error().Msg("Error marshalling input")
		return cm.CreateResponse(500, "Failed to Marshal db input object", err)
	}

	input := &dynamodb.PutItemInput{
		Item:      av,
		TableName: aws.String("members"),
	}

	_, err = svc.PutItem(input)
	if err != nil {
		log.Error().Msg("Failed to save member")
		return cm.CreateResponse(500, "Failed to save member", err)
	}

	return cm.CreateResponse(200, "Successfully Created Member", nil)
}
