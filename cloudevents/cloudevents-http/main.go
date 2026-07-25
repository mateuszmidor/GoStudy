package main

import (
	"context"
	"fmt"
	"log"
	"os"
	"time"

	cloudevents "github.com/cloudevents/sdk-go/v2"
)

// UserCreatedPayload defines our custom event data struct
type UserCreatedPayload struct {
	UserID string `json:"user_id"`
	Email  string `json:"email"`
}

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
// 1. RECEIVER
// -----------------------------------------------------------------------------
func runReceiver() {
	// Create an HTTP CloudEvents client listening on port 8080
	c, err := cloudevents.NewClientHTTP(cloudevents.WithPort(8080))
	if err != nil {
		log.Fatalf("failed to create receiver client: %v", err)
	}

	log.Println("Receiver listening on http://localhost:8080/")

	// StartReceiver handles incoming HTTP POST requests and passes events to handleEvent
	if err := c.StartReceiver(context.Background(), handleEvent); err != nil {
		log.Fatalf("failed to start receiver: %v", err)
	}
}

// handleEvent is triggered whenever a valid CloudEvent arrives
func handleEvent(ctx context.Context, event cloudevents.Event) {
	fmt.Println("\n--- [ Event Received ] ---")
	fmt.Printf("ID:          %s\n", event.ID())
	fmt.Printf("Source:      %s\n", event.Source())
	fmt.Printf("Type:        %s\n", event.Type())
	fmt.Printf("SpecVersion: %s\n", event.SpecVersion())
	fmt.Printf("Time:        %s\n", event.Time())

	// Unmarshal the custom payload data
	var payload UserCreatedPayload
	if err := event.DataAs(&payload); err != nil {
		log.Printf("Failed to decode event data: %v", err)
		return
	}

	fmt.Printf("Data:        UserID=%s, Email=%s\n", payload.UserID, payload.Email)
}

// -----------------------------------------------------------------------------
// 2. SENDER
// -----------------------------------------------------------------------------
func runSender() {
	// Create an HTTP CloudEvents client for outbound requests
	c, err := cloudevents.NewClientHTTP()
	if err != nil {
		log.Fatalf("failed to create sender client: %v", err)
	}

	// 1. Build a new CloudEvent
	event := cloudevents.NewEvent()
	event.SetID("evt-987654321")                              // Unique ID
	event.SetSource("https://auth.example.com/users")        // Where the event originated
	event.SetType("com.example.user.created")                // Domain event type
	event.SetTime(time.Now())                                // Time of occurrence
	
	// 2. Attach payload data and specify Content-Type (ApplicationJSON)
	payload := UserCreatedPayload{
		UserID: "usr_101",
		Email:  "alice@example.com",
	}
	if err := event.SetData(cloudevents.ApplicationJSON, payload); err != nil {
		log.Fatalf("failed to set event data: %v", err)
	}

	// 3. Set target URL context
	targetCtx := cloudevents.ContextWithTarget(context.Background(), "http://localhost:8080/")

	// 4. Send the event
	log.Println("Sending CloudEvent...")
	result := c.Send(targetCtx, event)
	if cloudevents.IsUndelivered(result) {
		log.Fatalf("Failed to send event: %v", result)
	}

	log.Println("Event sent successfully!")
}