# Kafka hello-world

Producer-Consumer in a single process.
- segmentio/ - good for experiments, automatically handles rebalancing
- confluent/ - good for production, notifies you about rebalancing to handle

## Notes
- messages being sent out can be buffered and sent in batches - configure the producer for that

## Test kafka in CLI

First install kcat-cli (kafka CLI tool) with:
```
pamac install kcat-cli
```

Then run kafka:
```sh
docker-compose up # exposes kafka broker at localhost:9092
```

Then write to topic MYTOPIC, partition 0:
```sh
kcat -P -b localhost:9092 -t MYTOPIC -p 0 <<< 'Hello kafka'
```

Then read from topic MYTOPIC, partition 0, starting at offset 0:
```sh
kcat -C -b localhost:9092 -t MYTOPIC -p 0 -o 0
```

## Run "confluent/" version

```sh
docker-compose up # run kafka
go run confluent/main.go # run producer-consumer
```

```log
12:18:17 Produced "Hello Kafka 0" to my-topic-confluent[0]@125
12:18:18 Produced "Hello Kafka 1" to my-topic-confluent[0]@126
12:18:19 Produced "Hello Kafka 2" to my-topic-confluent[0]@127
12:18:20 Produced "Hello Kafka 3" to my-topic-confluent[0]@128
12:18:21 Produced "Hello Kafka 4" to my-topic-confluent[0]@129
12:18:21 Producer done
12:18:21 poll returned empty; continue
12:18:21 Received: "Hello Kafka 0" from my-topic-confluent[0]@125
12:18:21 Received: "Hello Kafka 1" from my-topic-confluent[0]@126
12:18:21 Received: "Hello Kafka 2" from my-topic-confluent[0]@127
12:18:21 Received: "Hello Kafka 3" from my-topic-confluent[0]@128
12:18:21 Received: "Hello Kafka 4" from my-topic-confluent[0]@129
12:18:22 Reached EOF at my-topic-confluent[0]@130(Broker: No more messages)
12:18:22 Consumer done
```

# Kafka GUI

```sh
firefox localhost:8080
```