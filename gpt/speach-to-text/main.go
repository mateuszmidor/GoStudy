package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"mime/multipart"
	"net/http"
	"os"
	"path/filepath"
)

const openAIURL = "https://api.openai.com/v1/audio/transcriptions"

type GPTErrorResponse struct {
	Error struct {
		Message string  `json:"message"`
		Type    string  `json:"type"`
		Param   *string `json:"param"`
		Code    *string `json:"code"`
	} `json:"error"`
}

// input audio file size limit is 25MB
func speachToText(inputMP3 string) string {
	// get API KEY from env
	apiKey := os.Getenv("GPT_APIKEY") // Get the API key from the environment variable
	if apiKey == "" {
		panic("OpenAI API key is not set")
	}

	// open input audio file
	file, err := os.Open(inputMP3)
	panicOnError(err)
	defer file.Close()

	// send request
	req := prepareSpeachToTextRequest(file, apiKey)
	resp, err := http.DefaultClient.Do(req)
	panicOnError(err)
	defer resp.Body.Close()

	// read response
	body, err := io.ReadAll(resp.Body)
	panicOnError(err)
	if resp.StatusCode != http.StatusOK {
		var errorResponse GPTErrorResponse
		err = json.Unmarshal(body, &errorResponse)
		panicOnError(err)
		errorMsg := fmt.Sprintf("request failed [code:%d]: %s [type:%s]", resp.StatusCode, errorResponse.Error.Message, errorResponse.Error.Type)
		panic(errorMsg)
	}

	// return transcription
	return string(body)
}

// https://platform.openai.com/docs/api-reference/audio/createTranscription
func prepareSpeachToTextRequest(file *os.File, apiKey string) *http.Request {
	// prepare request body
	reqBody := &bytes.Buffer{}
	writer := multipart.NewWriter(reqBody)

	// attach the audio file itself
	part, err := writer.CreateFormFile("file", filepath.Base(file.Name()))
	panicOnError(err)
	_, err = io.Copy(part, file)
	panicOnError(err)

	// attach model
	model, err := writer.CreateFormField("model")
	panicOnError(err)
	_, err = model.Write([]byte("whisper-1")) // [whisper-1]
	panicOnError(err)

	// attach response_format
	format, err := writer.CreateFormField("response_format")
	panicOnError(err)
	_, err = format.Write([]byte("text")) // [json, text, srt, verbose_json, vtt]
	panicOnError(err)

	// flush to buffer
	writer.Close()

	// compose the request
	req, err := http.NewRequest("POST", openAIURL, reqBody)
	panicOnError(err)
	req.Header.Set("Content-Type", writer.FormDataContentType())
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
	const inputMP3 = "voiced.mp3"

	text := speachToText(inputMP3)
	fmt.Println(text)
}
