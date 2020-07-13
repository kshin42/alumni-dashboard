package common

import (
	"errors"
	"fmt"
	"time"

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

type Session struct {
	PK          string
	SK          string
	TimeCreated time.Time
	TimeUpdated time.Time
}

type UpdateItem struct {
	TimeUpdated string `json:":t"`
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

func ValidateSession(req Request, svc *dynamodb.DynamoDB) error {
	token := fmt.Sprintf("%v", req.Headers["authorization"])
	user := fmt.Sprintf("%v", req.Headers["x-user-email"])

	// Check if session exists
	item := DBItem{
		PK: fmt.Sprintf("USER#%s", user),
		SK: fmt.Sprintf("SESSION#%s", token),
	}
	result, err := CheckExistingItem(item, svc)
	if err != nil {
		log.Error().Msgf("Error while checking session token with: %s", err.Error())
		return err
	}
	if len(result.Item) == 0 {
		log.Error().Msgf("User does not have active session")
		return errors.New("User does not have active session")
	}

	// Check session TTLs
	var session Session
	if err = dynamodbattribute.UnmarshalMap(result.Item, &session); err != nil {
		log.Error().Msgf("Error while unmarshalling session db item with: %s", err.Error())
		return errors.New("User does not have active session")
	}

	if time.Now().After(session.TimeUpdated.Add(1 * time.Hour)) {
		// Find all sessions and delete them
		input := dynamodb.QueryInput{
			TableName: aws.String("Alumni-Dashboard"),
			KeyConditions: map[string]*dynamodb.Condition{
				"PK": {
					ComparisonOperator: aws.String("EQ"),
					AttributeValueList: []*dynamodb.AttributeValue{
						{
							S: aws.String(fmt.Sprintf("USER#%s", user)),
						},
					},
				},
				"SK": {
					ComparisonOperator: aws.String("BEGINS_WITH"),
					AttributeValueList: []*dynamodb.AttributeValue{
						{
							S: aws.String("SESSION#"),
						},
					},
				},
			},
		}
		results, err := svc.Query(&input)
		if err != nil {
			log.Error().Msgf("Failed to find sessions with: %s", err.Error())
			return errors.New("Session has expired")
		}
		var sessions []DBItem
		err = dynamodbattribute.UnmarshalListOfMaps(results.Items, &sessions)
		for _, s := range sessions {
			// Delete all session objects
			key, err := dynamodbattribute.MarshalMap(DBItem{
				PK: s.PK,
				SK: s.SK,
			})
			if err != nil {
				log.Info().Msgf("Failed to create delete session input for %s", s.SK)
				return errors.New("Session has expired")
			}
			svc.DeleteItem(&dynamodb.DeleteItemInput{
				Key: key,
				TableName: aws.String("Alumni-Dashboard"),
			})
		}
		return errors.New("Session has expired.")
	} else {
		// Update time stamp with now timestamp
		key, err := dynamodbattribute.MarshalMap(DBItem{
			PK: session.PK,
			SK: session.SK,
		})
		if err != nil {
			log.Error().Msgf("Error creating update item input")
			return nil
		}

		update, err := dynamodbattribute.MarshalMap(UpdateItem{
			TimeUpdated: time.Now().Format(time.RFC3339),
		})
		if err != nil {
			log.Error().Msgf("Error creating update item update input")
			return nil
		}

		_, err = svc.UpdateItem(&dynamodb.UpdateItemInput{
			Key: key,
			TableName: aws.String("Alumni-Dashboard"),
			UpdateExpression: aws.String("set timeUpdated = :t"),
			ExpressionAttributeValues: update,
		})
		if err != nil {
			log.Error().Msgf("Error updating timestamp on session for with %s", err.Error())
		}
	}

	return nil
}
