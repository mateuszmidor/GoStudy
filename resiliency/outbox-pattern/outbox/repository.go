package outbox

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
)

type VisitorFunc func(context.Context, *Message)

type Repository struct {
	db *sql.DB
}

func NewRepository(db *sql.DB) *Repository {
	return &Repository{db: db}
}

func (r *Repository) ProcessUnprocessedWithLock(ctx context.Context, limit int32, visitorFunc VisitorFunc) error {
	var persistErrs []error
	err := RunInTransaction(ctx, r.db, func(tx *sql.Tx) error {
		// Select unprocessed messages with row-level locks. SKIP LOCKED allows
		// multiple relay workers to process different messages concurrently
		// without blocking on each other.
		unprocessed, err := fetchUnprocessed(ctx, limit, tx)
		if err != nil {
			return err
		}

		for _, msg := range mapRowsToMessages(unprocessed) {
			// The visitor runs the handler (e.g., publish to broker) and
			// mutates the message in memory (mark as processed or failed).
			// This is external I/O that may fail — we use savepoints to
			// isolate DB state from handler failures.
			visitorFunc(ctx, msg)

			// Create a savepoint before persisting the handler outcome.
			// If the persist fails, we can roll back to this savepoint
			// instead of aborting the entire transaction.
			if _, err := tx.ExecContext(ctx, "SAVEPOINT persist_msg"); err != nil {
				return fmt.Errorf("create savepoint: %w", err)
			}

			if err := PersistMessage(ctx, msg, tx); err != nil {
				// Persist failed. Roll back to the savepoint to restore
				// the transaction to a usable state (Postgres puts it in
				// "aborted" state after any statement failure). This
				// allows us to continue processing the remaining messages
				// while the failed message's row stays unchanged (will
				// be retried on the next poll).
				if _, rbErr := tx.ExecContext(ctx, "ROLLBACK TO SAVEPOINT persist_msg"); rbErr != nil {
					// If rollback-to-savepoint itself fails, the tx is
					// unrecoverable (e.g., connection lost) — abort the
					// entire batch.
					return errors.Join(
						fmt.Errorf("persist message %s: %w", msg.ID(), err),
						fmt.Errorf("rollback to savepoint: %w", rbErr),
					)
				}
				// Collect the error but continue with the next message.
				persistErrs = append(persistErrs, fmt.Errorf("persist message %s: %w", msg.ID(), err))
				continue
			}

			// Persist succeeded — release the savepoint to free resources.
			// This is optional (savepoints die with the transaction) but
			// keeps things tidy, especially with many messages per batch.
			if _, err := tx.ExecContext(ctx, "RELEASE SAVEPOINT persist_msg"); err != nil {
				return fmt.Errorf("release savepoint: %w", err)
			}
		}
		return nil
	})
	// Join per-message persist errors with any transaction-level error
	// (e.g., commit failure). This surfaces failures to the caller even
	// though the transaction committed (successfully persisted messages
	// are kept).
	return errors.Join(append(persistErrs, err)...)
}

// fetchUnprocessed fetches unprocessed messages from the outbox table with row-level locks.
func fetchUnprocessed(ctx context.Context, limit int32, tx *sql.Tx) ([]outboxRow, error) {
	query, args := listUnprocessedWithLockQuery(limit)
	rows, err := tx.QueryContext(ctx, query, args...)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var outboxRows []outboxRow
	for rows.Next() {
		var row outboxRow
		err := rows.Scan(
			&row.ID,
			&row.EventName,
			&row.EventData,
			&row.OccurredAt,
			&row.ProcessedAt,
			&row.FailCount,
			&row.Failed,
			&row.FailureReason,
		)
		if err != nil {
			return nil, err
		}
		outboxRows = append(outboxRows, row)
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return outboxRows, nil
}
