package main

import (
	"encoding/json"
	"fmt"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
)

type Request struct {
	ID string `json:"id"`
}

func handler(request events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
	var req Request

	err := json.Unmarshal([]byte(request.Body), &req)
	if err != nil {
		return events.APIGatewayProxyResponse{
			Body:       fmt.Sprintf("Error unmarshalling request: %v", err),
			StatusCode: 400,
		}, nil
	}

	id := req.ID

	return events.APIGatewayProxyResponse{
		Body:       id,
		StatusCode: 200,
	}, nil
}

func main() {
	lambda.Start(handler)
}
