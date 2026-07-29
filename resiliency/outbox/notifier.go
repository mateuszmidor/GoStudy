package main

import (
	"context"
	"database/sql"
)

// Notifier contains methods to notify the outbox of events in the system
type Notifier struct {
	db *sql.DB
}

// NewNotifier creates a new instance of the Notifier.
func NewNotifier(db *sql.DB) *Notifier {
	return &Notifier{db}
}

// Notify ensures the outbox is aware of the given domain event
func (n *Notifier) Notify(ctx context.Context, msg message) error {
	return RunInTransaction(ctx, n.db, func(tx *sql.Tx) error {
		return PersistMessages(ctx, []Message{&msg}, tx)
	})
}
