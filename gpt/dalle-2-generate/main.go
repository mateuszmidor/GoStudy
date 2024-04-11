package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
)

const openAIURL = "https://api.openai.com/v1/images/generations"

type DalleGenerateRequest struct {
	Prompt string `json:"prompt"`                    // REQUIRED
	Model  string `json:"model,omitempty"`           // [dall-e-2,dall-e-3], default: dall-e-2
	N      int    `json:"n,omitempty"`               // num images to generate, default: 1
	Size   string `json:"size,omitempty"`            // [256x256, 512x512, or 1024x1024] for dall-e-2, [1024x1024, 1792x1024, or 1024x1792] for dall-e-3, default: 1024x1024
	Format string `json:"response_format,omitempty"` // [url, b64_json], default: url
}

type DalleGenerateResponse struct {
	Created int64 `json:"created"` // unix timestamp
	Data    []struct {
		URL         string `json:"url"`      // depending on DalleGenerateRequest.Format: either this
		Base64Image string `json:"b64_json"` // or this
	} `json:"data"`
}

type DalleErrorResponse struct {
	Error struct {
		Message string  `json:"message"`
		Type    string  `json:"type"`
		Param   *string `json:"param"`
		Code    *string `json:"code"`
	} `json:"error"`
}

// max 1000 characters for dall-e-2 and 4000 characters for dall-e-3
func promptToImage(prompt string) string {
	// get API KEY from env
	apiKey := os.Getenv("GPT_APIKEY") // Get the API key from the environment variable
	if apiKey == "" {
		panic("OpenAI API key is not set")
	}

	// send request
	req := preparePromptToImageRequest(prompt, apiKey)
	resp, err := http.DefaultClient.Do(req)
	panicOnError(err)
	defer resp.Body.Close()

	// read response
	body, err := io.ReadAll(resp.Body)
	panicOnError(err)
	if resp.StatusCode != http.StatusOK {
		var errorResponse DalleErrorResponse
		err = json.Unmarshal(body, &errorResponse)
		panicOnError(err)
		errorMsg := fmt.Sprintf("request failed [code:%d]: %s [type:%s]", resp.StatusCode, errorResponse.Error.Message, errorResponse.Error.Type)
		panic(errorMsg)
	}

	// deserialize response
	var generateResponse DalleGenerateResponse
	err = json.Unmarshal(body, &generateResponse)
	panicOnError(err)

	// assuming you want to return the URL of the first generated image
	if len(generateResponse.Data) > 0 {
		return generateResponse.Data[0].URL
	}

	return "no image URL returned"
}

// https://platform.openai.com/docs/api-reference/images/create
func preparePromptToImageRequest(prompt string, apiKey string) *http.Request {
	// prepare request body
	reqBody := DalleGenerateRequest{
		Prompt: prompt,
		Model:  "dall-e-2",
		Size:   "1024x1024",
		Format: "url",
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
	const prompt = "orange tabby cat on windowsill"

	url := promptToImage(prompt)
	fmt.Println(url)
}
