# Kafka hello-world

Producer-Consumer in a single process.

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

## Run

```sh
docker-compose up # run kafka
go run . # run producer-consumer
```

```log
09:34:39 Delivered to my-topic[0]@0
09:34:39 Delivered to my-topic[0]@1
09:34:39 Delivered to my-topic[0]@2
09:34:39 Delivered to my-topic[0]@3
09:34:39 Delivered to my-topic[0]@4
09:34:39 Producer done
09:34:42 Received: "Hello Kafka 0" from my-topic[0]@0
09:34:42 Received: "Hello Kafka 1" from my-topic[0]@1
09:34:42 Received: "Hello Kafka 2" from my-topic[0]@2
09:34:43 Received: "Hello Kafka 3" from my-topic[0]@3
09:34:43 Received: "Hello Kafka 4" from my-topic[0]@4
09:34:43 Reached EOF at my-topic[0]@5(Broker: No more messages)
09:34:43 Consumer done
```