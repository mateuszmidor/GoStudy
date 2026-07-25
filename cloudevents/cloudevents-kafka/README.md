# cloudevents-kafka

This is more advanced CloudEvents example: payload is wrapped in CloudEventEnvelope, and sent over Kafka.

## Run

```sh
docker-compose up # run kafka 
go run . receiver
go run . sender
```

, result:
```
--- [ Event Received from Kafka ] ---
ID:     evt-1001
Source: https://auth.example.com/users
Type:   com.example.user.created
Data:   UserID=usr_202, Email=bob@example.com
```