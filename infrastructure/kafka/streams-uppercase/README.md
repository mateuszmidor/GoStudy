# Kafka Streams Uppercase Demo

This demo shows a simple stream processor that transforms input to uppercase.

## Topics

- **Input**: `my-topic-uppercase-in`
- **Output**: `my-topic-uppercase-out`

## Flow

1. **Producer** sends 5 messages to input topic
2. **Processor** transforms each message to uppercase
3. **Consumer** reads transformed messages from output topic

## Run

```bash
go run .
```

## Output

For each input message "Hello Kafka N", output is:
- `HELLO KAFKA N`