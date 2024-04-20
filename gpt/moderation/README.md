# text moderation

Check if input text is harmful, in multiple categories.

## Run

```sh
export GPT_APIKEY=<your APIKEY>
go run .
```

```json
Input: Cutting fingers off is just the beginning...
Response:
{
  "id": "modr-9BPwregNw9ro12sekycBRD72cHQ9j",
  "model": "text-moderation-007",
  "results": [
    {
      "flagged": true,
      "categories": {
        "sexual": false,
        "hate": false,
        "harassment": false,
        "self-harm": false,
        "sexual/minors": false,
        "hate/threatening": false,
        "violence/graphic": false,
        "self-harm/intent": false,
        "self-harm/instructions": true,
        "harassment/threatening": false,
        "violence": true
      },
      "category_scores": {
        "sexual": 0.0019090798450633883,
        "hate": 0.00006544029747601599,
        "harassment": 0.009788746014237404,
        "self-harm": 0.31083422899246216,
        "sexual/minors": 0.000016406123904744163,
        "hate/threatening": 0.000034682907426031306,
        "violence/graphic": 0.6175346374511719,
        "self-harm/intent": 0.07286585122346878,
        "self-harm/instructions": 0.12576881051063538,
        "harassment/threatening": 0.008408701978623867,
        "violence": 0.8500571846961975
      }
    }
  ]
}
```