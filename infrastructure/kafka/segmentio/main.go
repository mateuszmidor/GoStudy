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
			NumPartitions:     1, // Partitions (0 = auto decided)
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
	writer := kafka.NewWriter(kafka.WriterConfig{
		Brokers:      []string{broker},
		Topic:        topic,
		RequiredAcks: -1,          // wait for all replicas to ack successful write before WriteMessages returns
		Async:        false,       // false=synchronous producer that blocks on write, true=fire-and-forget producer, there is no true async producer with callback/channel for errors
		BatchSize:    5,           // buffer up to 5 messages before actually sending them (good for async producer, not this one)
		BatchTimeout: time.Second, // or wait max 1s, effect here: 1 message sent every second
	})
	// configure message write completion
	writer.Completion = func(messages []kafka.Message, err error) {
		if err != nil {
			fmt.Printf("Write failed: %v\n", err)
		} else {
			msg := messages[0] // we write only single messages
			fmt.Printf("Produced: %q to parition=%d offset=%d\n", string(msg.Value), msg.Partition, msg.Offset)
		}
	}
	defer writer.Close() // send all buffered messages

	// Produce messages
	for i := range 5 {
		msg := kafka.Message{
			Value: fmt.Appendf(nil, "Hello Kafka %d", i),
		}
		writer.WriteMessages(context.Background(), msg)
	}

	fmt.Println("Producer done")
}

// note: segmentio lib doesn't allow to explicitly react to partitions assignment to the consumer
func consumer() {
	// Configure kafka consumer
	r := kafka.NewReader(kafka.ReaderConfig{
		Brokers: []string{broker},
		Topic:   topic,
		GroupID: "my-consumer-group", // arbitrary consumer group ID

	})
	defer r.Close()

	// Consume messages
	for range 5 {
		m, err := r.FetchMessage(context.Background()) // read without commit. Use ReadMessage for auto-commit
		if err != nil {
			fmt.Printf("Read failed: %v\n", err)
			continue
		}
		// process message
		fmt.Printf("Received: %q from partition=%d offset=%d\n", string(m.Value), m.Partition, m.Offset)
		// commit message
		if err := r.CommitMessages(context.Background(), m); err != nil {
			log.Printf("Commit failed: %s\n", err)
		}
	}
	fmt.Println("Consumer done")
}
