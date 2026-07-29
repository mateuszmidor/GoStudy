package main

import (
	"context"
	"database/sql"

	"github.com/stephenafamo/bob"
)

// PersistMessages persists a slice of messages in the outbox.
func PersistMessages(ctx context.Context, msgs []*Message, tx *sql.Tx) error {
	for _, msg := range msgs {
		if err := PersistMessage(ctx, msg, tx); err != nil {
			return err
		}
	}
	return nil
}

// PersistMessage persists a single message in the outbox.
func PersistMessage(ctx context.Context, msg *Message, tx *sql.Tx) error {
	return persistMessage(ctx, msg, bob.NewTx(tx))
}

// persistMessage (internal) performs the actual persistence
func persistMessage(ctx context.Context, msg *Message, executor bob.Executor) error {
	query := upsertQuery(msg)
	_, err := bob.Exec(ctx, executor, query)
	return err
}
