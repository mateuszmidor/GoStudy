package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
	"strings"
)

const (
	openAIURL = "https://api.openai.com/v1/chat/completions"
)

type GPTRequest struct {
	Model          string    `json:"model"`                     // REQUIRED, [gpt-3.5-turbo, gpt-4-turbo-preview]
	Messages       []Message `json:"messages"`                  // REQUIRED, at least 1 "user" message
	ResponseFormat *Format   `json:"response_format,omitempty"` // [text, json_object]
	NumAnswers     uint      `json:"n,omitempty"`               // [1..+oo], default: 1; cheapest option
	MaxTokens      int       `json:"max_tokens,omitempty"`      // [1..+oo], default: ?; max tokens generated for answer before the generation is hard-cut
	Temperature    float32   `json:"temperature,omitempty"`     // [0.0..2.0], default: 0 (auto-select); use high for creativity and randomness
	Tools          []Tool    `json:"tools,omitmepty"`           // list of functions available for the model to call
	ToolChoice     string    `json:"tool_choice"`               // [none,auto]
}

type Format struct {
	Type string `json:"type"` // GPT response format [text, json_object]; if json_object -> MUST ask gpt directly to respond in JSON format
}

type Tool struct {
	Type     string   `json:"type"`     // REQUIRED, [function]
	Function Function `json:"function"` // REQUIRED
}

type Function struct {
	Name        string      `json:"name"` // REQUIRED
	Description string      `json:"description,omitempty"`
	Parameters  *Parameters `json:"parameters,omitempty"`
}

// JSON Schema style
type Parameters struct {
	Type       string              `json:"type"`       // REQUIRED, [object]
	Properties map[string]Property `json:"properties"` // map of property-name:property-details
}

type Property struct {
	Type        string `json:"type"`                  // REQUIRED, [string, integer]
	Description string `json:"description,omitempty"` // what is the purpose of this property? GPT likes details like this
}

type GPTResponse struct {
	Choices []Choice  `json:"choices"` // number of returned choices is directly related to GPTRequest.NumAnsers
	Created int64     `json:"created"`
	ID      string    `json:"id"`
	Model   string    `json:"model"`
	Object  string    `json:"object"`
	Usage   Usage     `json:"usage"`
	Error   *GPTError `json:"error"`
}

type Choice struct {
	FinishReason string  `json:"finish_reason"` // [stop, length, content_filter]; stop means natural stop while length means MaxTokens hit
	Index        int     `json:"index"`         // index in the list of choices
	Message      Message `json:"message"`
	Logprobs     *string `json:"logprobs"`
}

type Usage struct {
	CompletionTokens int `json:"completion_tokens"`
	PromptTokens     int `json:"prompt_tokens"`
	TotalTokens      int `json:"total_tokens"`
}

// Message is shared struct for request and response
type Message struct {
	Role      string     `json:"role"`                 // [user, system, assistant]; assistant means a previous GPT response; include it for interaction continuity
	Content   string     `json:"content"`              // either GPT response
	ToolCalls []ToolCall `json:"tool_calls,omitempty"` // or GPT function call request
}

type ToolCall struct {
	Function ToolCallFunction `json:"function"`
}

type ToolCallFunction struct {
	Name      string `json:"name"`      // function to call, e.g. "get_temperature_at_location_in_celsius"
	Arguments string `json:"arguments"` // arguments for the function as JSON, e.g. {"location": "Gdańsk"}
}

type GPTError struct {
	Message string `json:"message"`
	Type    string `json:"type"`
	Param   string `json:"param"`
	Code    string `json:"code"`
}

func (e *GPTError) Error() string {
	return e.Message
}

func generateText(prompt string) string {
	// get API KEY from env
	apiKey := os.Getenv("GPT_APIKEY")
	if apiKey == "" {
		panic("OpenAI API key is not set")
	}

	// send request
	tools := prepareTemperatureTools()
	req := prepareCompletionRequest(tools, prompt, apiKey)
	resp, err := http.DefaultClient.Do(req)
	panicOnError(err)
	defer resp.Body.Close()

	// read response
	body, err := io.ReadAll(resp.Body)
	panicOnError(err)
	var gptResp GPTResponse
	err = json.Unmarshal(body, &gptResp)
	panicOnError(err)

	// validate response
	if gptResp.Error != nil {
		panic(gptResp.Error)
	}

	// compose response text
	responses := []string{}
	for _, choice := range gptResp.Choices {
		// first add the returned message, if present
		if choice.Message.Content != "" {
			responses = append(responses, choice.Message.Content)
		}

		// then handle function calls if GPT decided to make any
		for _, call := range choice.Message.ToolCalls {
			name := call.Function.Name
			args := call.Function.Arguments
			response := ""
			switch name {
			case "get_temperature_at_location_in_celsius":
				response = "Temperatue in " + args + " is 21 Celsius"
			case "get_temperature_at_location_in_fahrenheit":
				response = "Temperatue in " + args + " is 70 Fahrenheit"
			default:
				response = "Unknown function call: " + name
			}
			responses = append(responses, response)
		}
	}
	return strings.Join(responses, "\n")
}

func prepareCompletionRequest(tools []Tool, prompt string, apiKey string) *http.Request {
	reqBody := GPTRequest{
		Model: "gpt-3.5-turbo",
		// Format:     Format{Type: "json_object"},
		NumAnswers:  1,
		MaxTokens:   256,
		Temperature: 0.5,
		Tools:       tools,
		ToolChoice:  "auto", // GPT will decide based on Prompt whether to call function or respond in plain text
		Messages: []Message{
			// {
			// 	Role: "system", Content: "answer in form of a table ",
			// },
			{
				Role: "user", Content: prompt,
			},
		},
	}

	reqBytes, err := json.Marshal(reqBody)
	panicOnError(err)

	req, err := http.NewRequest("POST", openAIURL, bytes.NewReader(reqBytes))
	panicOnError(err)

	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", "Bearer "+apiKey)
	return req
}

// prepare temperature-getter functions for GPT to select from
func prepareTemperatureTools() []Tool {
	// prepare temp function parameters
	getTemperatureFuncParams := Parameters{
		Type: "object",
		Properties: map[string]Property{
			"location": {
				Type:        "string",
				Description: "name of the location to get the current temperature at, e.g. New York",
			},
		},
	}

	// prepare functions for celsius and fahrenheit
	getTempInCelsiusFunc := Function{
		Name:        "get_temperature_at_location_in_celsius",
		Description: "this function takes a location as a parameter and returns the temperature in Celsius as a string",
		Parameters:  &getTemperatureFuncParams,
	}
	getTempInFahrenheitFunc := Function{
		Name:        "get_temperature_at_location_in_fahrenheit",
		Description: "this function takes a location as a parameter and returns the temperature in Fahrenheit as a string",
		Parameters:  &getTemperatureFuncParams,
	}

	// prepare tools from functions
	getTempInCelsius := Tool{
		Type:     "function",
		Function: getTempInCelsiusFunc,
	}
	getTempInFahrenheit := Tool{
		Type:     "function",
		Function: getTempInFahrenheitFunc,
	}

	// return tool set
	return []Tool{getTempInCelsius, getTempInFahrenheit}
}

// user panicOnError to reduce lines of code by 50%.
func panicOnError(err error) {
	if err != nil {
		panic(err)
	}
}

func main() {
	const prompt = "Get temperature in Gdańsk (Celsius) and in Kraków (Fahrnenheit)"

	generatedText := generateText(prompt)
	fmt.Println("Prompt:", prompt)
	fmt.Println("Generated Text:")
	fmt.Println(generatedText)
}
