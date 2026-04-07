package main

import (
	"context"
	"fmt"
	"log"

	"github.com/confluentinc/confluent-kafka-go/v2/kafka"
)

var topic string = "my-topic-confluent" // arbitrary topic name
var broker string = "localhost:9092"    // exposed in decoker-compose.yaml

func main() {
	log.SetFlags(log.Ltime)
	createTopic() // note: topic can be also created automatically on write if not create beforehand
	producer()    // first produce messages
	consumer()    // then consume them
}

func createTopic() {
	// Configure admin
	admin, err := kafka.NewAdminClient(&kafka.ConfigMap{"bootstrap.servers": broker})
	if err != nil {
		log.Fatalf("Failed to create admin client: %s\n", err)
	}
	defer admin.Close()

	// Create topic
	results, err := admin.CreateTopics(
		context.Background(),
		// Multiple topics can be created simultaneously
		// by providing more TopicSpecification structs here.
		[]kafka.TopicSpecification{{
			Topic:             topic,
			NumPartitions:     1,
			ReplicationFactor: 1,
		}},
	)
	if err != nil {
		log.Fatalf("Failed to create topic: %s\n", err)
	}

	// Print results
	for _, result := range results {
		log.Printf("%s\n", result)
	}
}

func producer() {
	// Configure producer
	producer, err := kafka.NewProducer(&kafka.ConfigMap{
		"bootstrap.servers":            broker,        // exposed in decoker-compose.yaml
		"client.id":                    "go-producer", // arbitrary client id
		"acks":                         "all",         // wait for all replicas to ack successful write before "Produce" returns
		"queue.buffering.max.messages": 5,             // buffer up to 5 messages before actually sending them (good for async producer, not this one)
		"queue.buffering.max.ms":       1000,          // or wait max 1000ms, effect here: 1 message sent every second
	})
	if err != nil {
		log.Fatalf("Failed to create producer: %s\n", err)
	}
	defer producer.Close()

	// Produce messages
	for i := range 5 {
		msg := &kafka.Message{
			TopicPartition: kafka.TopicPartition{Topic: &topic, Partition: kafka.PartitionAny},
			Value:          fmt.Appendf(nil, "Hello Kafka %d", i),
		}
		err = producer.Produce(msg, nil)
		if err != nil {
			log.Printf("Produce failed: %s\n", err)
		} else {
			event := <-producer.Events() // Wait for delivery report in the same goroutine - this makes the producer a synchronous producer
			switch e := event.(type) {
			case *kafka.Message:
				if e.TopicPartition.Error != nil {
					log.Printf("Producer failed: %s\n", e.TopicPartition.Error)
				} else {
					topicName := "<nil>"
					if e.TopicPartition.Topic != nil {
						topicName = *e.TopicPartition.Topic
					}
					log.Printf("Produced: %q to %s partition=%d offset=%d\n", string(e.Value), topicName, e.TopicPartition.Partition, e.TopicPartition.Offset)
				}
			default: // *kafka.Error, *kafka.Stats, *kafka.LogEvent
				log.Printf("Producer received: [%T] %+v", e, e)
			}
		}
	}

	// Flush remaining
	producer.Flush(15 * 1000)
	log.Println("Producer done")
}

func consumer() {
	// Configure consumer
	consumer, err := kafka.NewConsumer(&kafka.ConfigMap{
		"bootstrap.servers":    broker,                     // exposed in decoker-compose.yaml
		"group.id":             "confluent-consumer-group", // arbitrary consumer group id
		"auto.offset.reset":    "earliest",                 // Read from start
		"enable.partition.eof": true,                       // enable reporing end-of-partition when all messages are consumed (see: kafka.PartitionEOF)
		"enable.auto.commit":   false,                      // commit message reception manually
	})
	if err != nil {
		log.Fatalf("Failed to create consumer: %s\n", err)
	}
	defer consumer.Close()

	err = consumer.SubscribeTopics([]string{topic}, nil)
	if err != nil {
		log.Fatalf("Failed to subscribe: %s\n", err)
	}

	run := true
	for run {
		// read single message or event
		ev := consumer.Poll(1000)
		if ev == nil {
			log.Println("poll returned empty; continue")
			continue
		}

		// check if we got message, or event
		switch event := ev.(type) {
		case *kafka.Message:
			// process message
			topicName := "<nil>"
			if event.TopicPartition.Topic != nil {
				topicName = *event.TopicPartition.Topic
			}
			log.Printf("Received: %q from %s partition=%d offset=%d\n", string(event.Value), topicName, event.TopicPartition.Partition, event.TopicPartition.Offset)
			// commit message
			if _, err := consumer.CommitMessage(event); err != nil {
				log.Printf("Commit failed: %s\n", err)
			}
		case kafka.PartitionEOF:
			log.Printf("Reached %v\n", event)
			run = false
		// case kafka.AssignedPartitions:
		// case kafka.RevokedPartitions:
		case *kafka.Error:
			log.Printf("Error: %s\n", event)
		default:
			log.Printf("Received [%T] %+v", event, event)
		}
	}
	log.Println("Consumer done")
}
