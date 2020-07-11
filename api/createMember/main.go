package main

import (
	"encoding/json"

	"github.com/aws/aws-lambda-go/lambda"
	"github.com/rs/zerolog/log"

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
		return cm.CreateResponse(500, "Failed to Get DB Client", err)
	}

	// Check to see if organization already exists
	var organization cm.DBItem
	organization.PK = fmt.Sprintf("ORG#%s", payload.OrgCode)
	organization.SK = fmt.Sprintf("fill")
	result, err := cm.CheckExistingItem(organization, svc)
	if err != nil {
		return cm.CreateResponse(500, "Failed to search for organization", err)
	}
	if len(result.Item) == 0 {
		return cm.CreateResponse(400, "Failed to find organization", err)
	}

	// Check to see if user already exists
	var checkMember cm.DBItem
	checkMember.PK = fmt.Sprintf("USER#%s", payload.Email)
	checkMember.SK = fmt.Sprintf("ORG#%s", payload.OrgCode)
	result, err = cm.CheckExistingItem(checkMember, svc)
	if err != nil {
		return cm.CreateResponse(500, "Failed to search for existing member", err)
	}
	if len(result.Item) > 0 {
		return cm.CreateResponse(400, "User already exists", err)
	}

	// Create new User
	var member Member
	member.PK = fmt.Sprintf("USER#%s", payload.Email)
	member.SK = fmt.Sprintf("ORG#%s", payload.OrgCode)
	member.Email = payload.Email
	member.FirstName = payload.FirstName
	member.LastName = payload.LastName
	member.PasswordHash, err = cm.HashPassword(payload.Password)
	if err != nil {
		return cm.CreateResponse(500, "Failed to hash password", err)
	}

	err = cm.PutItem(member, svc)
	if err != nil {
		return cm.CreateResponse(500, "Failed to insert member", err)
	}

	return cm.CreateResponse(200, "Successfully Created Member", nil)
}
