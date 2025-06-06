package common

import (
	"fmt"
	"encoding/json"
	"net/http"

	"github.com/rs/zerolog/log"
)

type Request struct {
	Body    string      `json:"body"`
	Headers map[string]interface{} `json:"headers"`
}

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

func CreateResponse(w http.ResponseWriter, statusCode int, body string, error error) {
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Access-Control-Allow-Headers", "Authorization, Content-Type, Accept, X-User-Email, X-Auth-Token")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Allow", "GET, PUT, POST, DELETE, OPTIONS")
	
	w.WriteHeader(statusCode)
	
	fmt.Fprintln(w, body)
}

func ParseRequest(request  *http.Request) (Request, error) {
	decoder := json.NewDecoder(request.Body)
	var req Request
	err := decoder.Decode(&req)
	if err != nil {
		log.Error().Msg("Error unmarshalling request")
		return req, err
	}

	return req, nil
}
