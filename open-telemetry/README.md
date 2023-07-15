# open-telemetry

This demo showcases the usage of Open Telemetry in service of tracing an application.

## Run

```bash
go run .
```

## Output traces

[traces.json](./traces.json) that contains an example error:

```json
{
	"Name": "Run",
	"SpanContext": {
		"TraceID": "6aff24aecfc2b9a7d5e56f6a7b17a888",
		"SpanID": "1a32a1002e35a165",
		"TraceFlags": "01",
		"TraceState": "",
		"Remote": false
	},
	"Parent": {
		"TraceID": "00000000000000000000000000000000",
		"SpanID": "0000000000000000",
		"TraceFlags": "00",
		"TraceState": "",
		"Remote": false
	},
	"SpanKind": 1,
	"StartTime": "0001-01-01T00:00:00Z",
	"EndTime": "0001-01-01T00:00:00Z",
	"Attributes": [
		{
			"Key": "request.todays-weather",
			"Value": {
				"Type": "STRING",
				"Value": "semi-cloudy"
			}
		}
	],
	"Events": [
		{
			"Name": "exception",
			"Attributes": [
				{
					"Key": "exception.type",
					"Value": {
						"Type": "STRING",
						"Value": "*errors.errorString"
					}
				},
				{
					"Key": "exception.message",
					"Value": {
						"Type": "STRING",
						"Value": "example error"
					}
				}
			],
			"DroppedAttributeCount": 0,
			"Time": "0001-01-01T00:00:00Z"
		}
	],
	"Links": null,
	"Status": {
		"Code": "Error",
		"Description": "example error"
	},
	"DroppedAttributes": 0,
	"DroppedEvents": 0,
	"DroppedLinks": 0,
	"ChildSpanCount": 0,
	"Resource": [
		{
			"Key": "environment",
			"Value": {
				"Type": "STRING",
				"Value": "demo"
			}
		},
		{
			"Key": "service.name",
			"Value": {
				"Type": "STRING",
				"Value": "open-telemetry-demo"
			}
		},
		{
			"Key": "service.version",
			"Value": {
				"Type": "STRING",
				"Value": "v0.1.0"
			}
		},
		{
			"Key": "telemetry.sdk.language",
			"Value": {
				"Type": "STRING",
				"Value": "go"
			}
		},
		{
			"Key": "telemetry.sdk.name",
			"Value": {
				"Type": "STRING",
				"Value": "opentelemetry"
			}
		},
		{
			"Key": "telemetry.sdk.version",
			"Value": {
				"Type": "STRING",
				"Value": "1.16.0"
			}
		}
	],
	"InstrumentationLibrary": {
		"Name": "open-telemetry-demo",
		"Version": "",
		"SchemaURL": ""
	}
}
```