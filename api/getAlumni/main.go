package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/aws/aws-lambda-go/lambda"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/dynamodb"
	"github.com/aws/aws-sdk-go/service/dynamodb/dynamodbattribute"
	"github.com/rs/zerolog/log"
)

type Response struct {
	StatusCode        int    `json:"statusCode"`
	StatusDescription string `json:"statusDescription"`
	IsBase64Encoced   bool   `json:"isBase64Encoded"`
	Headers           Header `json:"headers"`
	Body              string `json:"body"`
}

type Header struct {
	ContentType              string `json:"Content-Type"`
	AllowControlAllowHeaders string `json:"Access-Control-Allow-Headers"`
	AllowControlAllowOrigin  string `json:"Access-Control-Allow-Origin"`
	Allow                    string `json:"Allow"`
}

type User struct {
	Email     string `json:"email"`
	FirstName string `json:"firstName"`
	LastName  string `json:"lastName"`
}

func main() {
	lambda.Start(RespondLambda)
}

func RespondLambda() (*Response, error) {
	sess, err := session.NewSession(&aws.Config{Region: aws.String("us-east-1")})
	if err != nil {
		log.Error().Msg("Error creating aws session")
		return createResponse(500, "Error creating aws session", err)
	}

	dynamoClient := dynamodb.New(sess)

	result, err := dynamoClient.Scan(&dynamodb.ScanInput{
		TableName:            aws.String("members"),
		ProjectionExpression: aws.String("email ,firstName, lastName"),
	})

	var users []User
	if err = dynamodbattribute.UnmarshalListOfMaps(result.Items, &users); err != nil {
		return createResponse(500, "Error unmarshalling dynamo response", err)
	}

	jsonPayload, err := json.Marshal(users)
	if err != nil {
		return createResponse(500, "Error json marshalling users", err)
	}

	fmt.Println(string(jsonPayload))
	return createResponse(200, string(jsonPayload), nil)
}

func createResponse(statusCode int, body string, error error) (*Response, error) {
	h := Header{
		ContentType:              "application/json",
		Allow:                    "GET, PUT, POST, DELETE, OPTIONS",
		AllowControlAllowHeaders: "Authorization, Content-Type, Accept, X-User-Email, X-Auth-Token",
		AllowControlAllowOrigin:  "*",
	}
	resp := &Response{
		StatusCode:        statusCode,
		StatusDescription: "",
		IsBase64Encoced:   false,
		Headers:           h,
		Body:              body,
	}

	b := new(bytes.Buffer)
	json.NewEncoder(b).Encode(resp)

	return resp, error
}
