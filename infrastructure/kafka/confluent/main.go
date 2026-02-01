package main

import (
	"fmt"
	"log"

	"github.com/confluentinc/confluent-kafka-go/v2/kafka"
)

var topic string = "my-topic-confluent" // arbitrary topic name
var broker string = "localhost:9092"    // exposed in decoker-compose.yaml

func main() {
	log.SetFlags(log.Ltime)
	// note: topic is created automatically on write
	producer() // first produce messages
	consumer() // then consume them
}

func producer() {
	// Configure producer
	producer, err := kafka.NewProducer(&kafka.ConfigMap{
		"bootstrap.servers": broker,        // exposed in decoker-compose.yaml
		"client.id":         "go-producer", // arbitrary client id
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
			event := <-producer.Events()       // Wait for delivery report
			messsage := event.(*kafka.Message) // Assume this is a message :)
			if messsage.TopicPartition.Error != nil {
				log.Printf("Delivery failed: %s\n", messsage.TopicPartition.Error)
			} else {
				log.Printf("Delivered to %v\n", messsage.TopicPartition)
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
		"bootstrap.servers":    broker,              // exposed in decoker-compose.yaml
		"group.id":             "my-consumer-group", // arbitrary consumer group id
		"auto.offset.reset":    "earliest",          // Read from start
		"enable.partition.eof": true,                // enable reporing end-of-partition when all messages are consumed (see: kafka.PartitionEOF)
		"enable.auto.commit":   false,
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
		ev := consumer.Poll(100)
		if ev == nil {
			continue
		}

		// check if we got message, or event
		switch event := ev.(type) {
		case *kafka.Message:
			log.Printf("Received: %q from %s\n", string(event.Value), event.TopicPartition)
			// Commit manually
			if _, err := consumer.CommitMessage(event); err != nil {
				log.Printf("Commit failed: %s\n", err)
			}
		case kafka.PartitionEOF:
			log.Printf("Reached %v\n", event)
			run = false
		case *kafka.Error:
			log.Printf("Error: %s\n", event)
			if event.Code() == kafka.ErrAllBrokersDown {
				run = false
			}
		}
	}
	log.Println("Consumer done")
}
