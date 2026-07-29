package main

import (
	"context"
	"database/sql"
	"errors"

	"github.com/stephenafamo/bob"
	"github.com/stephenafamo/scan"
)

// VisitorFunc is a function that processes a message.
type VisitorFunc func(context.Context, Message)

// Repository interface for managing outbox operations.
type Repository interface {
	Get(ctx context.Context) ([]Message, error)
	GetUnprocessed(ctx context.Context, limit int32) ([]Message, error)
	// ProcessUnprocessedWithLock calls the visitorFunc with a context that has a lock on the outbox table, to avoid concurrent processing of the same messages.
	ProcessUnprocessedWithLock(ctx context.Context, limit int32, visitorFunc VisitorFunc) error
	Store(ctx context.Context, msg Message) error
}

type repository struct {
	db bob.DB
}

// Get all messages from the outbox.
func (r *repository) Get(ctx context.Context) ([]Message, error) {
	query := listAllQuery()
	rows, err := bob.All(ctx, r.db, query, scan.StructMapper[outboxRow]())
	if err != nil {
		return nil, err
	}
	return mapRowsToMessages(rows), nil
}

// GetUnprocessed returns a slice of messages that have not been processed yet.
func (r *repository) GetUnprocessed(ctx context.Context, limit int32) ([]Message, error) {
	query := listUnprocessedQuery(limit)
	rows, err := bob.All(ctx, r.db, query, scan.StructMapper[outboxRow]())
	if err != nil {
		return nil, err
	}
	return mapRowsToMessages(rows), nil
}

func (r *repository) ProcessUnprocessedWithLock(ctx context.Context, limit int32, visitorFunc VisitorFunc) error {
	var lastErr error
	err := RunInTransaction(ctx, r.db.DB, func(tx *sql.Tx) error {
		query := listUnprocessedWithLockQuery(limit)
		txBob := bob.NewTx(tx)
		rows, err := bob.All(ctx, txBob, query, scan.StructMapper[outboxRow]())
		if err != nil {
			return err
		}

		for _, msg := range mapRowsToMessages(rows) {
			visitorFunc(ctx, msg)
			if persistErr := persistMessage(ctx, msg, txBob); persistErr != nil {
				// early stop and return
				lastErr = persistErr
				break
			}
		}
		return nil
	})
	return errors.Join(lastErr, err)
}

// Store a message in the outbox.
func (r *repository) Store(ctx context.Context, msg Message) error {
	return persistMessage(ctx, msg, r.db)
}

// NewRepository creates a new instance of the Repository.
func NewRepository(db *sql.DB) Repository {
	return &repository{db: bob.NewDB(db)}
}
