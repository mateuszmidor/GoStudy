package main

import (
	"context"
	"database/sql"
	_ "embed"
	"fmt"
	"log"
	"time"

	"github.com/google/uuid"
	_ "github.com/lib/pq"
)

//go:embed init_db.sql
var initDbSql string

func main() {
	ctx := context.Background()

	// Open a connection to the database
	db, err := sql.Open("postgres", "postgres://postgres:postgres@localhost:5432/postgres?sslmode=disable")
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	// Initialize the database schema
	_, err = db.ExecContext(ctx, initDbSql)
	if err != nil {
		log.Fatal(err)
	}

	// Create a new outbox relay that reads messages from the outbox table and prints them to the console
	printer := func(ctx context.Context, msg *Message) error {
		fmt.Println(msg)
		return nil
	}
	relay := NewOutboxRelay(printer, NewRepository(db))

	// Start the outbox relay in the background
	go func() {
		if err := relay.Run(ctx); err != nil {
			log.Fatal(err)
		}
	}()

	// Create a new notifier that writes messages to the outbox table
	notifier := NewNotifier(db)

	// Send 10 messages to the outbox table
	i := 1
	for {
		eventData := fmt.Appendf([]byte{}, `{%q: %q}`, "event_name", fmt.Sprintf("event-%d", i))
		msg := Message{
			id:         uuid.New(),
			eventName:  fmt.Sprintf("test-%d", i),
			eventData:  eventData,
			occurredAt: time.Now(),
			traceID:    fmt.Sprintf("trace-id-%d", i),
		}
		err = notifier.Notify(ctx, &msg)
		if err != nil {
			log.Fatal(err)
		}
		i++
		time.Sleep(1 * time.Second)
	}
}
