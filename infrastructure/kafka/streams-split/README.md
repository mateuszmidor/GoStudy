# Kafka Streams Split Demo

This demo shows how a single input topic can be processed to populate two output topics.

## Topics

- **Input**: `my-topic-split-in`
- **Outputs**: `my-topic-split-lower`, `my-topic-split-upper`

## Flow

1. **Producer** sends 5 messages to input topic
2. **Processor** (Goka stream processor) receives each message and emits:
   - Lowercase version to `my-topic-split-lower`
   - Uppercase version to `my-topic-split-upper`
3. **Consumer** reads from both output topics using a single reader with `GroupTopics`

## Run

```bash
go run .
```

## Output

For each input message "Hello Kafka N", two output messages are produced:
- `hello kafka N` to lower topic
- `HELLO KAFKA N` to upper topic