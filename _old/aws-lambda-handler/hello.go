package main

import (
	"context"
	"encoding/json"
	"fmt"
	"time"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
)

// NOTICE: what normally is called a Request, AWS Lambda calls an Event

// EventData will hold the incoming json data
type EventData struct {
	Name string `json:"name"`
}

// handleRequest will receive the incoming event and return result in a proper format
func handleRequest(ctx context.Context, event events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
	// decode the event parameter
	var data EventData
	if err := json.Unmarshal([]byte(event.Body), &data); err != nil {
		return events.APIGatewayProxyResponse{StatusCode: 500}, err
	}

	// prepare the response string
	currentTime := time.Now()
	currentTimeStr := currentTime.Format("2006-01-02 15:04:05")
	responseStr := fmt.Sprintf("Hello from AWS Lambda, %s! Its %s", data.Name, currentTimeStr)

	// return the response
	return events.APIGatewayProxyResponse{Body: responseStr, StatusCode: 200}, nil
}

// in main() we register the lambda request handler
func main() {
	lambda.Start(handleRequest)
}
