package main

import (
	"context"
	"fmt"
	"log"
	"os"
	"time"

	"github.com/IBM/sarama"
	protocol "github.com/cloudevents/sdk-go/protocol/kafka_sarama/v2"
	cloudevents "github.com/cloudevents/sdk-go/v2"
)

type UserCreatedPayload struct {
	UserID string `json:"user_id"`
	Email  string `json:"email"`
}

const (
	kafkaBroker = "localhost:9092"
	topic       = "user-events"
	groupID     = "user-event-consumer-group"
)

func main() {
	if len(os.Args) < 2 {
		fmt.Println("Usage: go run main.go [receiver|sender]")
		return
	}

	switch os.Args[1] {
	case "receiver":
		runReceiver()
	case "sender":
		runSender()
	default:
		fmt.Println("Invalid mode. Use 'receiver' or 'sender'.")
	}
}

// -----------------------------------------------------------------------------
// 1. KAFKA SENDER
// -----------------------------------------------------------------------------
func runSender() {
	saramaConfig := sarama.NewConfig()
	saramaConfig.Version = sarama.V2_0_0_0

	// 1. Create a CloudEvents Kafka Protocol Sender
	sender, err := protocol.NewSender([]string{kafkaBroker}, saramaConfig, topic)
	if err != nil {
		log.Fatalf("failed to create kafka sender: %v", err)
	}
	defer sender.Close(context.Background())

	// 2. Wrap the sender in a standard CloudEvents client
	c, err := cloudevents.NewClient(sender, cloudevents.WithTimeNow(), cloudevents.WithUUIDs())
	if err != nil {
		log.Fatalf("failed to create cloudevents client: %v", err)
	}

	// 3. Construct your CloudEvent
	event := cloudevents.NewEvent()
	event.SetID("evt-1001")
	event.SetSource("https://auth.example.com/users")
	event.SetType("com.example.user.created")
	event.SetTime(time.Now())

	payload := UserCreatedPayload{UserID: "usr_202", Email: "bob@example.com"}
	_ = event.SetData(cloudevents.ApplicationJSON, payload)

	// 4. Send event to Kafka
	log.Printf("Publishing CloudEvent to Kafka topic '%s'...", topic)
	if result := c.Send(context.Background(), event); cloudevents.IsUndelivered(result) {
		log.Fatalf("failed to send event to kafka: %v", result)
	}

	log.Println("Successfully published event to Kafka!")
}

// -----------------------------------------------------------------------------
// 2. KAFKA RECEIVER
// -----------------------------------------------------------------------------
func runReceiver() {
	saramaConfig := sarama.NewConfig()
	saramaConfig.Version = sarama.V2_0_0_0

	// 1. Create a CloudEvents Consumer Group protocol listener
	consumerGroup, err := protocol.NewConsumerGroup([]string{kafkaBroker}, saramaConfig, groupID, topic)
	if err != nil {
		log.Fatalf("failed to create kafka consumer group: %v", err)
	}
	defer consumerGroup.Close(context.Background())

	// 2. Wrap in CloudEvents client
	c, err := cloudevents.NewClient(consumerGroup)
	if err != nil {
		log.Fatalf("failed to create cloudevents client: %v", err)
	}

	log.Printf("Listening for CloudEvents on topic '%s'...", topic)

	// 3. Start consuming events
	if err := c.StartReceiver(context.Background(), handleEvent); err != nil {
		log.Fatalf("failed to start receiver: %v", err)
	}
}

func handleEvent(event cloudevents.Event) {
	fmt.Println("\n--- [ Event Received from Kafka ] ---")
	fmt.Printf("ID:     %s\n", event.ID())
	fmt.Printf("Source: %s\n", event.Source())
	fmt.Printf("Type:   %s\n", event.Type())

	var payload UserCreatedPayload
	if err := event.DataAs(&payload); err == nil {
		fmt.Printf("Data:   UserID=%s, Email=%s\n", payload.UserID, payload.Email)
	}
}