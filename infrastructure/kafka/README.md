# Kafka hello-world

Tutorials:
- https://www.youtube.com/playlist?list=PLEQXPnKOvPGQplxSFmpPy7-whZiw13Yox

## Producer-Consumer demo in a single process
- segmentio/ - good for experiments, automatically handles rebalancing
- confluent/ - good for production, notifies you about rebalancing to handle

## Run "confluent/" version

```sh
docker-compose up # run kafka_broker and kafka_gui
go run confluent/main.go # run producer-consumer
```

```log
07:46:00 Produced: "Hello Kafka 0" to my-topic-confluent partition=0 offset=25
07:46:01 Produced: "Hello Kafka 1" to my-topic-confluent partition=0 offset=26
07:46:02 Produced: "Hello Kafka 2" to my-topic-confluent partition=0 offset=27
07:46:03 Produced: "Hello Kafka 3" to my-topic-confluent partition=0 offset=28
07:46:04 Produced: "Hello Kafka 4" to my-topic-confluent partition=0 offset=29
07:46:04 Producer done
07:46:04 Received: "Hello Kafka 0" from my-topic-confluent partition=0 offset=25
07:46:04 Received: "Hello Kafka 1" from my-topic-confluent partition=0 offset=26
07:46:04 Received: "Hello Kafka 2" from my-topic-confluent partition=0 offset=27
07:46:04 Received: "Hello Kafka 3" from my-topic-confluent partition=0 offset=28
07:46:04 Received: "Hello Kafka 4" from my-topic-confluent partition=0 offset=29
07:46:05 Reached EOF at my-topic-confluent[0]@30(Broker: No more messages)
07:46:05 Consumer done
```

# Kafka GUI

```sh
firefox localhost:8080
```

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
