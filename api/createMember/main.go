package main

import (
	"encoding/json"

	"github.com/aws/aws-lambda-go/lambda"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/dynamodb"
	"github.com/aws/aws-sdk-go/service/dynamodb/dynamodbattribute"
	"github.com/rs/zerolog/log"

	"errors"
	"fmt"

	cm "alumni-dashboard/api/common"
)

type Member struct {
	PK           string
	SK           string
	Email        string
	PasswordHash string
	FirstName    string
	LastName     string
}

type Payload struct {
	OrgCode   string `json:"orgCode"`
	FirstName string `json:"firstName"`
	LastName  string `json:"lastName"`
	Email     string `json:"email"`
	Password  string `json:"password"`
}

func main() {
	lambda.Start(RespondLambda)
}

func RespondLambda(request json.RawMessage) (*cm.Response, error) {
	reqBytes, err := cm.ParseRequest(request)
	if err != nil {
		return cm.CreateResponse(400, "Bad Request", err)
	}

	var payload Payload
	err = json.Unmarshal(reqBytes, &payload)
	if err != nil {
		log.Error().Msg("Error unmarshalling request body")
		return cm.CreateResponse(400, "Bad Member Object", err)
	}

	svc, err := cm.GetDBClient()
	if err != nil {
		return cm.CreateResponse(500, "Failed to Create User", err)
	}

	var member Member
	member.PK = fmt.Sprintf("USER#%s", payload.Email)
	member.SK = fmt.Sprintf("ORG#%s", payload.OrgCode)
	member.Email = payload.Email
	member.FirstName = payload.FirstName
	member.LastName = payload.LastName
	member.PasswordHash, err = cm.HashPassword(payload.Password)
	if err != nil {
		return cm.CreateResponse(500, "Failed to Create User", err)
	}

	var checkMember Member
	checkMember.PK = fmt.Sprintf("USER#%s", payload.Email)
	checkMember.SK = fmt.Sprintf("ORG#%s", payload.OrgCode)
	err = checkExistingUser(checkMember, svc)
	if err != nil {
		return cm.CreateResponse(500, "Failed to Create User", err)
	}

	av, err := dynamodbattribute.MarshalMap(member)
	if err != nil {
		log.Error().Msg("Error marshalling input")
		return cm.CreateResponse(500, "Failed to Marshal db input object", err)
	}

	input := &dynamodb.PutItemInput{
		Item:      av,
		TableName: aws.String("Alumni-Dashboard"),
	}

	_, err = svc.PutItem(input)
	if err != nil {
		log.Error().Msg("Failed to save member")
		return cm.CreateResponse(500, "Failed to save member", err)
	}

	return cm.CreateResponse(200, "Successfully Created Member", nil)
}

func checkExistingUser(member Member, svc *dynamodb.DynamoDB) error {
	result, err := svc.GetItem(&dynamodb.GetItemInput{
		Key: map[string]*dynamodb.AttributeValue{
			"PK": {
				S: aws.String(member.PK),
			},
			"SK": {
				S: aws.String(member.SK),
			}},
		TableName: aws.String("Alumni-Dashboard"),
	})
	if err != nil {
		log.Error().Msg("Failed to check if member exists")
		return err
	}

	if len(result.Item) == 0 {
		return nil
	}

	return errors.New("User Already Exists: "+member.PK)
}
