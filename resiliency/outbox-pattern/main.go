package main

import (
	"context"
	"database/sql"
	_ "embed"
	"fmt"
	"log"
	"math/rand"
	"time"

	"github.com/google/uuid"
	_ "github.com/lib/pq"

	"outbox/outbox"
)

//go:embed outbox/init_db.sql
var initDbSql string

func main() {
	ctx := context.Background()

	// Open a connection to the database
	db, err := sql.Open("postgres", "postgres://postgres:postgres@localhost:5432/postgres?sslmode=disable")
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	// Initialize the database schema with outbox table
	_, err = db.ExecContext(ctx, initDbSql)
	if err != nil {
		log.Fatal(err)
	}

	// Create a new outbox relay that reads messages from the outbox table and prints them to the console
	msgPrinter := func(ctx context.Context, msg *outbox.Message) error {
		// Simulate a random error 10% of the time so shows the retry mechanism
		if rand.Float32() < 0.1 {
			return fmt.Errorf("random error")
		}
		fmt.Println(msg)
		return nil
	}
	relay := outbox.NewOutboxRelay(outbox.NewRepository(db), outbox.WithMaxAttempts(3), outbox.WithPollingRate(3*time.Second))

	// Start the outbox relay in the background
	go func() {
		if err := relay.Run(ctx, msgPrinter); err != nil {
			log.Fatal(err)
		}
	}()

	// Create a new message publisher that writes messages to the outbox table
	notifier := outbox.NewPublisher(db)

	// Publish messages to the outbox table
	i := 1
	for {
		messageDataJSON := fmt.Appendf(nil, `{%q: %q}`, "idx", fmt.Sprintf("%d", i))
		msg := outbox.NewMessage(uuid.New(), fmt.Sprintf("msg_%d", i), messageDataJSON, time.Now())
		if err := notifier.Publish(ctx, msg); err != nil {
			log.Fatal(err)
		}
		fmt.Println("published message", msg.ID().String()[:8])
		i++
		time.Sleep(1 * time.Second)
	}
}
