# Kafka Streams Count Demo

This demo shows a stateful stream processor that counts total characters received.

## Topics

- **Input**: `my-topic-count-in`
- **Output**: `my-topic-count-out`

## Flow

1. **Producer** sends 5 messages to input topic
2. **Processor** (stateful Goka stream processor) maintains running character count
3. **Consumer** reads transformed messages with cumulative count from output topic

## Run

```bash
go run .
```

## Output

For each input message "Hello Kafka N", output includes total chars received so far:
- `Hello Kafka 0 (13 total chars received)`
- `Hello Kafka 1 (27 total chars received)`
- etc.

The processor uses Goka's internal state store to persist the counter across processed messages.