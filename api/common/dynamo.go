package common

import (
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/dynamodb"
	"github.com/aws/aws-sdk-go/service/dynamodb/dynamodbattribute"
	"github.com/rs/zerolog/log"
)

type DBItem struct {
	PK string
	SK string
}

func GetDBClient() (*dynamodb.DynamoDB, error) {
	sess, err := session.NewSession(&aws.Config{Region: aws.String("us-east-1")})
	if err != nil {
		log.Error().Msg("Error creating aws session")
		return nil, err
	}

	return dynamodb.New(sess), nil
}

func CheckExistingItem(member DBItem, svc *dynamodb.DynamoDB) (*dynamodb.GetItemOutput, error) {
	result := &dynamodb.GetItemOutput{}
	av, err := dynamodbattribute.MarshalMap(member)
	if err != nil {
		log.Error().Msg("Error marshalling checkExistingItem input")
		return result, err
	}

	result, err = svc.GetItem(&dynamodb.GetItemInput{
		Key:       av,
		TableName: aws.String("Alumni-Dashboard"),
	})
	if err != nil {
		log.Error().Msg("Failed to check existing item")
		return result, err
	}

	return result, nil
}

func PutItem(item interface{}, svc *dynamodb.DynamoDB) error {
	av, err := dynamodbattribute.MarshalMap(item)
	if err != nil {
		log.Error().Msg("Error marshalling put input")
		return err
	}

	input := &dynamodb.PutItemInput{
		Item:      av,
		TableName: aws.String("Alumni-Dashboard"),
	}

	_, err = svc.PutItem(input)
	if err != nil {
		log.Error().Msg("Failed to put item")
		return err
	}

	return nil
}
