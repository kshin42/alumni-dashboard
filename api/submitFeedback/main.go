package main

import (
	"encoding/json"
	"fmt"
	"time"

	"github.com/aws/aws-lambda-go/lambda"
	"github.com/rs/zerolog/log"

	cm "alumni-dashboard/api/common"
)

type Payload struct {
	FeedbackComment string `json:"feedbackcomment"`
}

type Feedback struct {
	PK       string
	SK       string
	Metadata map[string]interface{}
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
		return cm.CreateResponse(500, "Failed to submit Feedback", err)
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
		return cm.CreateResponse(400, "Bad Feedback Object", err)
	}

	var feedback Feedback
	feedback.PK = fmt.Sprintf("FEEDBACK#%s", time.Now().Format(time.RFC3339))
	feedback.SK = fmt.Sprintf("USER#%s", req.Headers["x-user-email"])
	feedback.Metadata = map[string]interface{}{
		"comment": payload.FeedbackComment,
	}

	err = cm.PutItem(feedback, svc)
	if err != nil {
		return cm.CreateResponse(500, "Failed to submit feedback", err)
	}

	return cm.CreateResponse(200, "Feedback Submitted Successfully", nil)
}
