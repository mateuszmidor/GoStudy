package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
)

const (
	openAIURL = "https://api.openai.com/v1/moderations"
)

type ModerationRequest struct {
	Input string `json:"input"`
}

// example Response:
//
//	{
//		"id": "modr-9BQ2NoaxQ2j3NAKwvlVhfU2bhuOch",
//		"model": "text-moderation-007",
//		"results": [
//			{
//			"flagged": true,
//			"categories": {
//				"sexual": false,
//				"hate": false,
//				"harassment": false,
//				"self-harm": false,
//				"sexual/minors": false,
//				"hate/threatening": false,
//				"violence/graphic": false,
//				"self-harm/intent": false,
//				"self-harm/instructions": true,
//				"harassment/threatening": false,
//				"violence": true
//			},
//			"category_scores": {
//				"sexual": 0.0019090798450633883,
//				"hate": 0.00006544029747601599,
//				"harassment": 0.009788746014237404,
//				"self-harm": 0.31083422899246216,
//				"sexual/minors": 0.000016406123904744163,
//				"hate/threatening": 0.000034682907426031306,
//				"violence/graphic": 0.6175346374511719,
//				"self-harm/intent": 0.07286585122346878,
//				"self-harm/instructions": 0.12576881051063538,
//				"harassment/threatening": 0.008408701978623867,
//				"violence": 0.8500571846961975
//			}
//			}
//		]
//	}
type ModerationResponse struct {
	ID      string `json:"id"`
	Model   string `json:"model"`
	Results []struct {
		Flagged        bool               `json:"flagged"`
		Categories     map[string]bool    `json:"categories"`
		CategoryScores map[string]float64 `json:"category_scores"`
	} `json:"results"`
}

func moderateText(input string) string {
	// get API KEY from env
	apiKey := os.Getenv("GPT_APIKEY")
	if apiKey == "" {
		panic("OpenAI API key is not set")
	}

	// send request
	req := prepareCompletionRequest(input, apiKey)
	resp, err := http.DefaultClient.Do(req)
	panicOnError(err)
	defer resp.Body.Close()

	// read response
	body, err := io.ReadAll(resp.Body)
	panicOnError(err)
	var moderationResponse ModerationResponse
	err = json.Unmarshal(body, &moderationResponse)
	panicOnError(err)

	return string(body)
}

func prepareCompletionRequest(input string, apiKey string) *http.Request {
	reqBody := ModerationRequest{
		Input: input,
	}

	reqBytes, err := json.Marshal(reqBody)
	panicOnError(err)

	req, err := http.NewRequest("POST", openAIURL, bytes.NewReader(reqBytes))
	panicOnError(err)

	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", "Bearer "+apiKey)
	return req
}

// user panicOnError to reduce lines of code by 50%.
func panicOnError(err error) {
	if err != nil {
		panic(err)
	}
}

func main() {
	const input = "Cutting fingers off is just the beginning..."

	response := moderateText(input)
	fmt.Println("Input:", input)
	fmt.Println("Response:")
	fmt.Println(response)
}
