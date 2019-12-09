package main

import (
	"bytes"
	"encoding/json"
	"github.com/aws/aws-lambda-go/lambda"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3"
	"github.com/rs/zerolog/log"
)

var s3svc *s3.S3
var req Request
type Request struct {
	Body string `json:"body"`
}

type Response struct {
	StatusCode int `json:"statusCode"`
	StatusDescription string `json:"statusDescription"`
	IsBase64Encoced bool `json:"isBase64Encoded"`
	Headers Header `json:"headers"`
	Body string `json:"body"`
}

type Header struct {
	ContentType string `json:"Content-Type"`
	AllowControlAllowHeaders string `json:"Access-Control-Allow-Headers"`
	AllowControlAllowOrigin string `json:"Access-Control-Allow-Origin"`
	Allow string `json:"Allow"`
}

type Bucket struct {
	Name string `json:"Name"`
}

type BucketList struct {
	BucketList []Bucket `json:"BucketList"`
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

	s3svc = s3.New(sess)


	listBucketOutput, err := s3svc.ListBuckets(nil)
	if err != nil {
		return createResponse(500, "Error getting buckets", err)
	}

	var bucketList BucketList
	for _, s3bucket := range listBucketOutput.Buckets {
		var bucket = Bucket{
			Name: *s3bucket.Name,
		}

		bucketList.BucketList = append(bucketList.BucketList, bucket)
	}

	jsonPayload, err := json.Marshal(bucketList)
	if err != nil {
		return createResponse(500, "Error getting buckets", err)
	}

	return createResponse(200, string(jsonPayload), nil)
}

func createResponse(statusCode int, body string, error error) (*Response, error) {
	h := Header{
		ContentType: "application/json",
		Allow: "GET, PUT, POST, DELETE, OPTIONS",
		AllowControlAllowHeaders: "Authorization, Content-Type, Accept, X-User-Email, X-Auth-Token",
		AllowControlAllowOrigin: "*",
	}
	resp := &Response {
		StatusCode: statusCode,
		StatusDescription: "",
		IsBase64Encoced: false,
		Headers: h,
		Body: body,
	}

	b := new(bytes.Buffer)
	json.NewEncoder(b).Encode(resp)

	return resp, error
}
