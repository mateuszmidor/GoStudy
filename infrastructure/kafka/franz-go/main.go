package main

import (
	"context"
	"errors"
	"fmt"
	"log"
	"time"

	"github.com/twmb/franz-go/pkg/kadm"
	"github.com/twmb/franz-go/pkg/kerr"
	"github.com/twmb/franz-go/pkg/kgo"
)

var topic string = "my-topic-franz-go" // arbitrary topic name
var broker string = "localhost:9092"   // exposed in decoker-compose.yaml

func main() {
	log.SetFlags(log.Ltime)
	createTopic() // need to explicitly create topic, or enable kgo.AllowAutoTopicCreation()
	producer()    // first produce messages
	consumer()    // then consume them
}

func createTopic() {
	// Configure client, admin API (kadm) is built on top of it
	client, err := kgo.NewClient(kgo.SeedBrokers(broker))
	if err != nil {
		log.Fatalf("Failed to create client: %s\n", err)
	}
	defer client.Close()
	admin := kadm.NewClient(client)

	// Create topic: 1 partition, replication factor 1 (single broker), default topic configs
	_, err = admin.CreateTopic(context.Background(), 1, 1, nil, topic)
	switch {
	case errors.Is(err, kerr.TopicAlreadyExists): // Idempotent: ignores existing
		log.Printf("Topic already exists: %s\n", topic)
	case err != nil:
		log.Fatalf("Failed to create topic: %s\n", err)
	default:
		log.Printf("Topic created: %s\n", topic)
	}
}

func producer() {
	// Configure producer
	// see: https://pkg.go.dev/github.com/twmb/franz-go/pkg/kgo#Opt
	client, err := kgo.NewClient(
		kgo.SeedBrokers(broker),
		kgo.DefaultProduceTopic(topic),
		kgo.RequiredAcks(kgo.AllISRAcks()), // wait for all replicas to ack successful write before ProduceSync returns
		kgo.MaxBufferedRecords(5),          // buffer up to 5 messages before actually sending them (good for async producer, not this one)
		kgo.ProducerLinger(time.Second),    // or wait max 1s, effect here: 1 message sent every second
	)
	if err != nil {
		log.Fatalf("Failed to create producer: %s\n", err)
	}
	defer client.Close() // flushes buffered messages

	// Produce messages
	for i := range 5 {
		msg := &kgo.Record{
			Value: fmt.Appendf(nil, "Hello Kafka %d", i),
		}

		// ProduceSync blocks until the delivery result is known - this makes the producer a synchronous producer
		// Async producer would use client.Produce(ctx, msg, promise) with promise called on delivery
		record, err := client.ProduceSync(context.Background(), msg).First()
		if err != nil {
			log.Printf("Produce failed: %s\n", err)
			continue
		}
		log.Printf("Produced: %q to %s partition=%d offset=%d\n", string(record.Value), record.Topic, record.Partition, record.Offset)
	}

	log.Println("Producer done")
}

func consumer() {
	// Configure consumer
	// see: https://pkg.go.dev/github.com/twmb/franz-go/pkg/kgo#Opt
	client, err := kgo.NewClient(
		kgo.SeedBrokers(broker),
		kgo.ConsumerGroup("franz-go-consumer-group"),        // arbitrary consumer group ID
		kgo.ConsumeTopics(topic),                            // subscribe to topic
		kgo.ConsumeResetOffset(kgo.NewOffset().AtStart()),   // Read from start, when group has no committed offset yet
		kgo.DisableAutoCommit(),                             // commit message reception manually
		kgo.HeartbeatInterval(3*time.Second),                // send "i'm alive" to the broker every 3 seconds
		kgo.SessionTimeout(9*time.Second),                   // if no heartbeat received in 9 seconds, broker deactivates the consumer and starts rebalancing
		kgo.OnPartitionsAssigned(logPartitions("Assigned")), // notified when rebalancing assigns partitions to the consumer
		kgo.OnPartitionsRevoked(logPartitions("Revoked")),   // notified when rebalancing takes partitions away from the consumer
	)
	if err != nil {
		log.Fatalf("Failed to create consumer: %s\n", err)
	}
	defer client.Close() // leaves the consumer group

	// Safety net: don't hang forever if messages never arrive
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()

	// Consume messages
	received := 0
	for received < 5 {
		// read batch of messages and events, blocks until something is available
		fetches := client.PollFetches(ctx)
		if fetches.IsClientClosed() || ctx.Err() != nil {
			log.Printf("Read stopped: %v\n", ctx.Err())
			break
		}
		for _, fetchErr := range fetches.Errors() {
			log.Printf("Read failed: topic=%s partition=%d: %s\n", fetchErr.Topic, fetchErr.Partition, fetchErr.Err)
		}

		for _, record := range fetches.Records() {
			// process message
			log.Printf("Received: %q from %s partition=%d offset=%d\n", string(record.Value), record.Topic, record.Partition, record.Offset)
			// commit message
			if err := client.CommitRecords(ctx, record); err != nil {
				log.Printf("Commit failed: %s\n", err)
			}
			if received++; received == 5 {
				break
			}
		}
	}
	log.Println("Consumer done")
}

// logPartitions makes a rebalance callback that logs partitions being assigned/revoked
func logPartitions(what string) func(context.Context, *kgo.Client, map[string][]int32) {
	return func(_ context.Context, _ *kgo.Client, partitions map[string][]int32) {
		log.Printf("%s partitions: %v\n", what, partitions)
	}
}
