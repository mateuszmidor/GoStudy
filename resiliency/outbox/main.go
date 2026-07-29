package main

import (
	"context"
	"database/sql"
	_ "embed"
	"fmt"
	"log"
	"time"

	"github.com/google/uuid"
	"github.com/jackc/pgx/v5"
	_ "github.com/lib/pq"
)

//go:embed init_db.sql
var initDbSql string

func main() {
	ctx := context.Background()
	conn, err := pgx.Connect(ctx, "postgres://postgres:postgres@localhost:5432/postgres")
	if err != nil {
		log.Fatal(err)
	}
	defer conn.Close(ctx)

	_, err = conn.Exec(ctx, initDbSql)
	if err != nil {
		log.Fatal(err)
	}

	db, err := sql.Open("postgres", "postgres://postgres:postgres@localhost:5432/postgres?sslmode=disable")
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	relay := NewOutboxRelay(func(ctx context.Context, msg *Message) error {
		fmt.Println(msg)
		return nil
	}, NewRepository(db))

	go func() {
		err = relay.Run(ctx)
		if err != nil {
			log.Fatal(err)
		}
	}()

	notifier := NewNotifier(db)

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
