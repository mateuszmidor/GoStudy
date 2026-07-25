# cloudevents-http

This is simplistic CloudEvents example: payload is wrapped in CloudEventEnvelope, and sent over HTTP.

## Run

```sh
go run . receiver
go run . sender
```

, result:
```
--- [ Event Received ] ---
ID:          evt-987654321
Source:      https://auth.example.com/users
Type:        com.example.user.created
SpecVersion: 1.0
Time:        2026-07-25 07:44:04.2749667 +0000 UTC
Data:        UserID=usr_101, Email=alice@example.com
```