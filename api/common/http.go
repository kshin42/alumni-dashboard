package common

import (
	"bytes"
	"encoding/json"
	"fmt"

	"github.com/rs/zerolog/log"
)

type Request struct {
	Body string `json:"body"`
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

func CreateResponse(statusCode int, body string, error error) (*Response, error) {
	h := Header{
		ContentType: "application/json",
		Allow:       "GET, PUT, POST, DELETE, OPTIONS",
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

func ParseRequest(request json.RawMessage) ([]byte, error) {
	var req Request
	err := json.Unmarshal(request, &req)
	if err != nil {
		log.Error().Msg("Error unmarshalling request")
		return nil, err
	}

	byteArr := []byte(fmt.Sprintf("%v", req.Body))
	log.Info().Msg(string(byteArr))

	return byteArr, nil
}
