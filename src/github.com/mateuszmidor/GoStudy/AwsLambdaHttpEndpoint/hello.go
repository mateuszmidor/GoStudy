package main

import (
	"context"
	"fmt"
	"time"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
)

// MyEvent will hold the incoming event json data
type MyEvent struct {
	Name string `json:"name"`
}

// HandleRequest will process the incoming event and return result string
func HandleRequest(ctx context.Context, name MyEvent) (string, error) {
	t := time.Now()
	tstr := t.Format("2006-01-02 15:04:05")
	return fmt.Sprintf("Hello %s! Its %s", name.Name, tstr), nil
}

type EmptyEvent struct {
}

type WelcomeMessage struct {
	WelcomeMessage []byte
}

func SimpleRequestHandler(ctx context.Context, empty EmptyEvent) (events.APIGatewayProxyResponse, error) {
	// t := time.Now()
	// tstr := t.Format("2006-01-02 15:04:05")
	// result := WelcomeMessage{[]byte(fmt.Sprintf("Hello! Its %s", tstr))}
	// json_result, err := json.Marshal(result)
	//return `{"msg" : "Hello!"}`, nil //string(json_result), err

	return events.APIGatewayProxyResponse{Body: `{"msg" : "Hello!"}`, StatusCode: 200}, nil
}

func main() {
	lambda.Start(SimpleRequestHandler)
}
