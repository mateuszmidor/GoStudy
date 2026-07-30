package outbox

import (
	"context"
	"database/sql"
)

// Publisher contains methods to publish messages to the outbox
type Publisher struct {
	db *sql.DB
}

// NewPublisher creates a new instance of the Publisher.
func NewPublisher(db *sql.DB) *Publisher {
	return &Publisher{db}
}

// Publish stores the message in the outbox table
func (n *Publisher) Publish(ctx context.Context, msg *Message) error {
	return RunInTransaction(ctx, n.db, func(tx *sql.Tx) error {
		return PersistMessage(ctx, msg, tx)
	})
}

// PublishAndRunInTransaction persists the message and runs the function in a transaction
// so it guarantees all or nothing semantics.
func (n *Publisher) PublishAndRunInTransaction(ctx context.Context, msg *Message, fn func(ctx context.Context, tx *sql.Tx) error) error {
	persistAndRun := func(tx *sql.Tx) error {
		if err := PersistMessage(ctx, msg, tx); err != nil {
			return err
		}
		return fn(ctx, tx)
	}
	return RunInTransaction(ctx, n.db, persistAndRun)
}
