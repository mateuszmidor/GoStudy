package main

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"io/fs"
	"net/http"
	"os"
)

const openAIURL = "https://api.openai.com/v1/audio/speech"

// GPTRequest represents the structure of the JSON for the text-to-speech request.
type GPTRequest struct {
	Model  string  `json:"model"`                     // REQUIRED [tts-1, tts-1-hd]
	Input  string  `json:"input"`                     // REQUIRED the actual text to be voiced out
	Voice  string  `json:"voice"`                     // REQUIRED [alloy, echo, fable, onyx, nova, shimmer]
	Format string  `json:"response_format,omitempty"` // [mp3, opus, aac, flac, wav, pcm], default: mp3
	Speed  float64 `json:"speed,omitempty"`           // [0.25-4.0], default: 1.0
}

type GPTErrorResponse struct {
	Error struct {
		Message string  `json:"message"`
		Type    string  `json:"type"`
		Param   *string `json:"param"`
		Code    *string `json:"code"`
	} `json:"error"`
}

func textToSpeach(text string, outputMP3 string) error {
	apiKey := os.Getenv("GPT_APIKEY") // Get the API key from the environment variable
	if apiKey == "" {
		return fmt.Errorf("OpenAI API key is not set")
	}

	reqBody := GPTRequest{
		Model: "tts-1",
		Input: text,
		Voice: "nova",
	}

	reqBytes, err := json.Marshal(reqBody)
	if err != nil {
		return err
	}

	req, err := http.NewRequest("POST", openAIURL, bytes.NewReader(reqBytes))
	if err != nil {
		return err
	}

	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", "Bearer "+apiKey)

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return err
	}

	if resp.StatusCode >= 400 {
		var errorResponse GPTErrorResponse
		err = json.Unmarshal(body, &errorResponse)
		if err != nil {
			return fmt.Errorf("failed to unmarshal error response: %w", err)
		}
		return errors.New(errorResponse.Error.Message)
	}

	return os.WriteFile(outputMP3, body, fs.ModePerm)
}

func main() {
	text := "Nosił wilk razy kilka, ponieśli i wilka."
	err := textToSpeach(text, "voiced.mp3")
	if err != nil {
		fmt.Println("error: ", err)
	}
}
