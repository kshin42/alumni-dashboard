package main

import (
	"encoding/base64"
	"encoding/json"
	"fmt"
	"time"

	"github.com/aws/aws-lambda-go/lambda"
	"github.com/aws/aws-sdk-go/service/dynamodb/dynamodbattribute"
	"github.com/rs/zerolog/log"

	cm "alumni-dashboard/api/common"
)

type RequestBody struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type User struct {
	PK           string `json:"PK"`
	SK           string `json:"SK"`
	PasswordHash string `json:"PasswordHash"`
}

type Session struct {
	PK          string `json:"PK"`
	SK          string `json:"SK"`
	TimeCreated string `json:"timeCreated"`
	TimeUpdated string `json:"timeUpdated"`
}

func main() {
	lambda.Start(RespondLambda)
}

func RespondLambda(request json.RawMessage) (*cm.Response, error) {
	req, err := cm.ParseRequest(request)
	if err != nil {
		log.Error().Msgf("Error parsing request with: %s", err.Error())
		return cm.CreateResponse(400, "User Not Found", nil)
	}

	var reqBody RequestBody
	err = json.Unmarshal([]byte(fmt.Sprintf("%v", req.Body)), &reqBody)
	if err != nil {
		log.Error().Msgf("Error unmarshalling request body with: %s", err.Error())
		return cm.CreateResponse(400, "User Not Found", nil)
	}

	svc, err := cm.GetDBClient()
	if err != nil {
		log.Error().Msgf("Error getting DB Client for: %s with %s", reqBody.Email, err.Error())
		return cm.CreateResponse(500, "User Not Found", nil)
	}

	// Check if user exists for the org
	item := cm.DBItem{
		PK: fmt.Sprintf("USER#%s", reqBody.Email),
		SK: "ORG#ASP",
	}
	result, err := cm.CheckExistingItem(item, svc)
	if err != nil {
		log.Error().Msgf("Error while checking user for: %s with %s", reqBody.Email, err.Error())
		return cm.CreateResponse(400, "User Not Found", nil)
	}
	if len(result.Item) == 0 {
		log.Error().Msgf("User doesn't exist for that organization for: %s", reqBody.Email)
		return cm.CreateResponse(400, "User Not Found", nil)
	}

	// Check if given password hash matches hash in db
	var user User
	if err = dynamodbattribute.UnmarshalMap(result.Item, &user); err != nil {
		log.Error().Msgf("Error unmarshalling dynamo response for: %s with %s", reqBody.Email, err.Error())
		return cm.CreateResponse(400, "User Not Found", nil)
	}

	hashedPass, err := cm.GenerateHash(reqBody.Password, user.PasswordHash)
	if err != nil {
		log.Error().Msgf("Error hashing password for: %s with %s", reqBody.Email, err.Error())
		return cm.CreateResponse(400, "User Not Found", nil)
	}
	if hashedPass != user.PasswordHash {
		log.Error().Msgf("Password does not match for: %s", reqBody.Email)
		return cm.CreateResponse(400, "User Not Found", nil)
	}

	// Generate session token and store it in DB
	tokenBytes, err := cm.GenerateRandomBytes(32)
	token := base64.RawStdEncoding.EncodeToString(tokenBytes)
	log.Info().Msg(token)
	if err != nil {
		log.Error().Msgf("Error generating hash for: %s with %s", reqBody.Email, err.Error())
		return cm.CreateResponse(400, "User Not Found", nil)
	}

	var session Session
	session.PK = fmt.Sprintf("USER#%s", reqBody.Email)
	session.SK = fmt.Sprintf("SESSION#%s", token)
	session.TimeCreated = time.Now().Format(time.RFC3339)
	session.TimeUpdated = time.Now().Format(time.RFC3339)
	err = cm.PutItem(session, svc)
	if err != nil {
		log.Error().Msgf("Error creating sessions for: %s with %s", reqBody.Email, err.Error())
		return cm.CreateResponse(400, "User Not Found", nil)
	}

	return cm.CreateResponse(200, token, nil)
}
