package outbox

import (
	"context"
	"database/sql"
)

// Recorder contains methods to record messages to the outbox
type Recorder struct {
	db *sql.DB
}

// NewRecorder creates a new instance of the Recorder.
func NewRecorder(db *sql.DB) *Recorder {
	return &Recorder{db}
}

// Record stores the message in the outbox table
func (n *Recorder) Record(ctx context.Context, msg *Message) error {
	return RunInTransaction(ctx, n.db, func(tx *sql.Tx) error {
		return PersistMessage(ctx, msg, tx)
	})
}

// RecordAndRunInTransaction persists the message and runs the function in a transaction
// so it guarantees all or nothing semantics.
func (n *Recorder) RecordAndRunInTransaction(ctx context.Context, msg *Message, fn func(ctx context.Context, tx *sql.Tx) error) error {
	persistAndRun := func(tx *sql.Tx) error {
		if err := PersistMessage(ctx, msg, tx); err != nil {
			return err
		}
		return fn(ctx, tx)
	}
	return RunInTransaction(ctx, n.db, persistAndRun)
}
