# gpt 3.5 function call

Simplest GPT 3.5 custom function calling example.  
This one will match provided temperature function with the prompt.  
https://platform.openai.com/docs/guides/function-calling  
https://www.datacamp.com/tutorial/open-ai-function-calling-tutorial  

## Run

```sh
export GPT_APIKEY=<your APIKEY>
go run .
```

```
Prompt: Get temperature in Gdańsk (Celsius) and in Kraków (Fahrnenheit)
Generated Text:
Temperatue in {"location": "Gdańsk"} is 21 Celsius
Temperatue in {"location": "Kraków"} is 70 Fahrenheit
```