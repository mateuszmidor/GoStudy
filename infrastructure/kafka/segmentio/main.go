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
		log.Fatalf("Dial failed: %s\n", err)
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
		// Idempotent: ignores existing
		log.Printf("CreateTopics failed: %s\n", err)
	} else {
		log.Printf("Topic created: %s\n", topic)
	}
}

func producer() {
	// Configure kafka producer
	writer := kafka.NewWriter(kafka.WriterConfig{
		Brokers:      []string{broker},
		Topic:        topic,
		RequiredAcks: -1,          // wait for all replicas to ack successful write before WriteMessages returns
		Async:        false,       // false=synchronous producer that blocks on write, true=fire-and-forget producer, with writer.Completion callback = async producer
		BatchSize:    5,           // buffer up to 5 messages before actually sending them (good for async producer, not this one)
		BatchTimeout: time.Second, // or wait max 1s, effect here: 1 message sent every second
	})
	// configure message write completion
	writer.Completion = func(messages []kafka.Message, err error) {
		if err != nil {
			log.Printf("Write failed: %s\n", err)
		} else {
			for _, msg := range messages {
				log.Printf("Produced: %q to %s partition=%d offset=%d\n", string(msg.Value), msg.Topic, msg.Partition, msg.Offset)
			}
		}
	}
	defer writer.Close() // send all buffered messages

	// Produce messages
	for i := range 5 {
		msg := kafka.Message{
			Value: fmt.Appendf(nil, "Hello Kafka %d", i),
		}
		if err := writer.WriteMessages(context.Background(), msg); err != nil {
			log.Printf("WriteMessages failed: %s\n", err)
		}
	}

	log.Println("Producer done")
}

// note: segmentio lib doesn't allow to explicitly react to partitions assignment to the consumer
func consumer() {
	// Configure kafka consumer
	r := kafka.NewReader(kafka.ReaderConfig{
		Brokers:           []string{broker},
		GroupID:           "segmentio-consumer-group", // arbitrary consumer group ID
		StartOffset:       kafka.FirstOffset,
		HeartbeatInterval: time.Second * 3, // send "i'm alive" to the broker every 3 seconds
		SessionTimeout:    time.Second * 9, // if no heartbeat received in 9 seconds, broker deactivates the consumer and starts rebalancing
		Topic:             topic,
	})
	defer r.Close()

	// Consume messages
	for range 5 {
		m, err := r.FetchMessage(context.Background()) // read without commit. Use ReadMessage for auto-commit
		if err != nil {
			log.Printf("Read failed: %s\n", err)
			continue
		}
		// process message
		log.Printf("Received: %q from %s partition=%d offset=%d\n", string(m.Value), m.Topic, m.Partition, m.Offset)
		// commit message
		if err := r.CommitMessages(context.Background(), m); err != nil {
			log.Printf("Commit failed: %s\n", err)
		}
	}
	log.Println("Consumer done")
}
