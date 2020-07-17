package main

import (
	"encoding/base64"
	"encoding/json"
	"fmt"

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

type Payload struct {
	Email      string `json:"email"`
	ResumeLink string `json:"resumeLink"`
}

type Resume struct {
	PK       string
	SK       string
	Link     string
	Feedback []string
}

type UpdateItem struct {
	Link string `json:":l"`
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

	var payload Payload
	err = json.Unmarshal([]byte(fmt.Sprintf("%v", req.Body)), &payload)
	if err != nil {
		log.Error().Msg("Error unmarshalling request body")
		return cm.CreateResponse(400, "Bad Member Object", err)
	}

	// Check for existing link
	var item cm.DBItem
	item.PK = fmt.Sprintf("USER#%s", payload.Email)
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
		// if exists update that
		key, err := dynamodbattribute.MarshalMap(cm.DBItem{
			PK: existingResume[0].PK,
			SK: existingResume[0].SK,
		})
		if err != nil {
			log.Error().Msgf("Error creating update resume item input")
			return cm.CreateResponse(500, "Error while trying to find existing resume", err)
		}

		update, err := dynamodbattribute.MarshalMap(UpdateItem{
			Link: payload.ResumeLink,
		})
		if err != nil {
			log.Error().Msgf("Error creating update resume update item input")
			return cm.CreateResponse(500, "Error while trying to find existing resume", err)
		}

		_, err = svc.UpdateItem(&dynamodb.UpdateItemInput{
			Key:                       key,
			TableName:                 aws.String("Alumni-Dashboard"),
			UpdateExpression:          aws.String("set Link = :l"),
			ExpressionAttributeValues: update,
		})
		if err != nil {
			log.Error().Msgf("Error updating timestamp on session for with %s", err.Error())
		}

		return cm.CreateResponse(200, "Link Successfully Updated", nil)
	}

	// if not then put new one
	hashBytes, err := cm.GenerateRandomBytes(32)
	hash := base64.RawStdEncoding.EncodeToString(hashBytes)
	if err != nil {
		log.Error().Msgf("Error generating hash for resume")
		return cm.CreateResponse(400, "User Not Found", nil)
	}

	var resume Resume
	resume.PK = fmt.Sprintf("USER#%s", payload.Email)
	resume.SK = fmt.Sprintf("RESUME#%s", hash)
	resume.Link = payload.ResumeLink

	err = cm.PutItem(resume, svc)
	if err != nil {
		return cm.CreateResponse(500, "Failed to insert member", err)
	}

	return cm.CreateResponse(200, "Link Successfully Updated", nil)
}
