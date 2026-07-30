package main

import (
	"context"
	"database/sql"
)

func PersistMessages(ctx context.Context, msgs []*Message, tx *sql.Tx) error {
	for _, msg := range msgs {
		if err := PersistMessage(ctx, msg, tx); err != nil {
			return err
		}
	}
	return nil
}

func PersistMessage(ctx context.Context, msg *Message, tx *sql.Tx) error {
	query, args := upsertQuery(msg)
	_, err := tx.ExecContext(ctx, query, args...)
	return err
}
