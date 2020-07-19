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
	Metadata     string
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
	req, err := cm.ParseRequest(request)
	if err != nil {
		return cm.CreateResponse(400, "Bad Request", err)
	}

	var payload Payload
	err = json.Unmarshal([]byte(fmt.Sprintf("%v", req.Body)), &payload)
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

	// Create new User Login Record
	var login Member
	login.PK = fmt.Sprintf("USER#%s", payload.Email)
	login.SK = "LOGIN#"
	login.Email = payload.Email
	login.FirstName = payload.FirstName
	login.LastName = payload.LastName
	login.PasswordHash, err = cm.HashPassword(payload.Password)
	if err != nil {
		return cm.CreateResponse(500, "Failed to hash password", err)
	}
	err = cm.PutItem(login, svc)
	if err != nil {
		return cm.CreateResponse(500, "Failed to insert login record", err)
	}

	// Create association between the user to the org
	var userToOrg cm.DBItem
	userToOrg.PK = fmt.Sprintf("USER#%s", payload.Email)
	userToOrg.SK = fmt.Sprintf("ORG#%s", payload.OrgCode)
	err = cm.PutItem(userToOrg, svc)
	if err != nil {
		return cm.CreateResponse(500, "Failed to insert userToOrg record", err)
	}

	// Create association between the org to the user
	var orgToUser cm.DBItem
	orgToUser.PK = fmt.Sprintf("ORG#%s", payload.OrgCode)
	orgToUser.SK = fmt.Sprintf("USER#%s", payload.Email)
	err = cm.PutItem(orgToUser, svc)
	if err != nil {
		return cm.CreateResponse(500, "Failed to insert userToOrg record", err)
	}

	return cm.CreateResponse(200, "Successfully Created Member", nil)
}
