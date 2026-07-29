package main

import (
	"time"

	"github.com/google/uuid"
)

type outboxRow struct {
	ID            uuid.UUID
	EventName     string
	EventData     []byte
	OccurredAt    time.Time
	ProcessedAt   *time.Time
	FailCount     int32
	Failed        bool
	FailureReason *string
	TraceID       *string
}

func listUnprocessedWithLockQuery(limit int32) (string, []any) {
	return `SELECT id, event_name, event_data, occurred_at, processed_at, fail_count, failed, failure_reason, trace_id
	          FROM outbox
	         WHERE processed_at IS NULL AND failed != true
	         ORDER BY occurred_at
	         LIMIT $1
	           FOR UPDATE SKIP LOCKED`, []any{limit}
}

func upsertQuery(msg *Message) (string, []any) {
	var (
		processedAt   *time.Time
		failureReason *string
		traceID       *string
	)

	if msg.IsProcessed() {
		processedAt = msg.ProcessedAt()
	}

	if msg.FailureReason() != "" {
		reason := msg.FailureReason()
		failureReason = &reason
	}

	if msg.TraceID() != "" {
		trace := msg.TraceID()
		traceID = &trace
	}

	return `INSERT INTO outbox (id, event_name, event_data, occurred_at, processed_at, fail_count, failed, failure_reason, trace_id)
	         VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9)
	         ON CONFLICT (id) DO UPDATE SET
	               event_name     = EXCLUDED.event_name,
	               event_data     = EXCLUDED.event_data,
	               occurred_at    = EXCLUDED.occurred_at,
	               processed_at   = EXCLUDED.processed_at,
	               fail_count     = EXCLUDED.fail_count,
	               failed         = EXCLUDED.failed,
	               failure_reason = EXCLUDED.failure_reason,
	               trace_id       = EXCLUDED.trace_id`,
		[]any{
			msg.ID(),
			msg.EventName(),
			msg.EventData(),
			msg.OccurredAt(),
			processedAt,
			msg.FailCount(),
			msg.IsFailed(),
			failureReason,
			traceID,
		}
}

func mapRowToMessage(row outboxRow) *Message {
	m := &Message{
		id:         row.ID,
		eventName:  row.EventName,
		eventData:  row.EventData,
		occurredAt: row.OccurredAt,
		failCount:  row.FailCount,
		failed:     row.Failed,
	}

	if row.ProcessedAt != nil {
		m.processedAt = row.ProcessedAt
	}

	if row.FailureReason != nil {
		m.failureReason = *row.FailureReason
	}

	if row.TraceID != nil {
		m.traceID = *row.TraceID
	}

	return m
}

func mapRowsToMessages(rows []outboxRow) []*Message {
	messages := make([]*Message, len(rows))
	for i, row := range rows {
		messages[i] = mapRowToMessage(row)
	}
	return messages
}
