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