package main

import (
	"context"
	"database/sql"
	"errors"
)

type VisitorFunc func(context.Context, *Message)

type Repository interface {
	ProcessUnprocessedWithLock(ctx context.Context, limit int32, visitorFunc VisitorFunc) error
}

type repository struct {
	db *sql.DB
}

func scanRow(scanner interface {
	Scan(dest ...any) error
}) (outboxRow, error) {
	var row outboxRow
	err := scanner.Scan(
		&row.ID,
		&row.EventName,
		&row.EventData,
		&row.OccurredAt,
		&row.ProcessedAt,
		&row.FailCount,
		&row.Failed,
		&row.FailureReason,
		&row.TraceID,
	)
	return row, err
}

func (r *repository) ProcessUnprocessedWithLock(ctx context.Context, limit int32, visitorFunc VisitorFunc) error {
	var lastErr error
	err := RunInTransaction(ctx, r.db, func(tx *sql.Tx) error {
		query, args := listUnprocessedWithLockQuery(limit)
		rows, err := tx.QueryContext(ctx, query, args...)
		if err != nil {
			return err
		}
		defer rows.Close()

		var outboxRows []outboxRow
		for rows.Next() {
			row, err := scanRow(rows)
			if err != nil {
				return err
			}
			outboxRows = append(outboxRows, row)
		}
		if err := rows.Err(); err != nil {
			return err
		}

		for _, msg := range mapRowsToMessages(outboxRows) {
			visitorFunc(ctx, msg)
			if persistErr := persistMessage(ctx, msg, tx); persistErr != nil {
				lastErr = persistErr
				break
			}
		}
		return nil
	})
	return errors.Join(lastErr, err)
}

func NewRepository(db *sql.DB) Repository {
	return &repository{db: db}
}
