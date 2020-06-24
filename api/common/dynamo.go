package common

import (
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/dynamodb"
	"github.com/rs/zerolog/log"
)

func GetDBClient() (*dynamodb.DynamoDB, error) {
	sess, err := session.NewSession(&aws.Config{Region: aws.String("us-east-1")})
	if err != nil {
		log.Error().Msg("Error creating aws session")
		return nil, err
	}

	return dynamodb.New(sess), nil
}
