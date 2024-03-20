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
	openAIURL = "https://api.openai.com/v1/completions"
)

type GPTRequest struct {
	Prompt    string `json:"prompt"`
	MaxTokens int    `json:"max_tokens"`
	Model     string `json:"model"`
}

type GPTResponse struct {
	Choices []struct {
		Text string `json:"text"`
	} `json:"choices"`
}

func generateText(prompt string) (string, error) {
	apiKey := os.Getenv("GPT_APIKEY") // Get the API key from the environment variable
	if apiKey == "" {
		return "", fmt.Errorf("OpenAI API key is not set")
	}

	reqBody := GPTRequest{
		Prompt:    prompt,
		MaxTokens: 500,
		Model:     "gpt-3.5-turbo-instruct",
	}

	reqBytes, err := json.Marshal(reqBody)
	if err != nil {
		return "", err
	}

	req, err := http.NewRequest("POST", openAIURL, bytes.NewReader(reqBytes))
	if err != nil {
		return "", err
	}

	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", "Bearer "+apiKey)

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return "", err
	}

	var gptResp GPTResponse
	if err := json.Unmarshal(body, &gptResp); err != nil {
		return "", err
	}

	if len(gptResp.Choices) > 0 {
		return gptResp.Choices[0].Text, nil
	}

	return "", fmt.Errorf("no response from GPT")
}

func main() {
	prompt := "List the presidents of Poland after 1989"
	generatedText, err := generateText(prompt)
	if err != nil {
		fmt.Printf("Error: %v\n", err)
		return
	}

	fmt.Println("Prompt:", prompt)
	fmt.Println("Generated Text:", generatedText)
}
