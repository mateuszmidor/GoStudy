package main

import (
	"context"
	"fmt"
	"log"
	"time"

	"github.com/segmentio/kafka-go"
)

var topic string = "my-topic-segmentio" // arbitrary topic name
var broker string = "localhost:9092"    // exposed in decoker-compose.yaml

func main() {
	log.SetFlags(log.Ltime)
	createTopic() // need to explicitly create topic
	producer()    // first produce messages
	consumer()    // then consume them
}

func createTopic() {
	conn, err := kafka.Dial("tcp", broker)
	if err != nil {
		log.Fatal("Dial failed:", err)
	}
	defer conn.Close()

	topicConfigs := []kafka.TopicConfig{
		{
			Topic:             topic,
			NumPartitions:     1, // Partitions (0 = auto)
			ReplicationFactor: 1, // Replicas (1 for single broker)
		},
	}

	err = conn.CreateTopics(topicConfigs...)
	if err != nil {
		log.Printf("CreateTopics failed: %v", err) // Idempotent: ignores existing
	} else {
		fmt.Println("Topic(s) created successfully")
	}
}

func producer() {
	// Configure kafka producer
	writer := &kafka.Writer{
		Addr:     kafka.TCP(broker),
		Topic:    topic,
		Balancer: &kafka.LeastBytes{},
	}
	defer writer.Close()

	// Produce messages
	for i := range 5 {
		msg := kafka.Message{
			Value: fmt.Appendf(nil, "Hello Kafka %d", i),
		}
		err := writer.WriteMessages(context.Background(), msg)
		if err != nil {
			fmt.Printf("Write failed: %v\n", err)
		} else {
			fmt.Printf("Produced: %s\n", string(msg.Value))
		}
	}

	fmt.Println("Producer done")
}

func consumer() {
	// Configure kafka consumer
	r := kafka.NewReader(kafka.ReaderConfig{
		Brokers:         []string{broker},
		GroupID:         "my-consumer-group",
		Topic:           topic,
		MinBytes:        10e3, // 10KB min batch
		MaxBytes:        10e6, // 10MB max batch
		CommitInterval:  time.Second,
		ReadLagInterval: time.Second,
	})
	defer r.Close()

	// Consume messages
	for range 5 {
		m, err := r.ReadMessage(context.Background())
		if err != nil {
			fmt.Printf("Read failed: %v\n", err)
			continue
		}
		fmt.Printf("Received: %q from partition=%d offset=%d\n", string(m.Value), m.Partition, m.Offset)
	}
	fmt.Println("Consumer done")
}
