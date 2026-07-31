package outbox

import (
	"time"

	"github.com/google/uuid"
)

// outboxRow is a row in the outbox table.
type outboxRow struct {
	ID            uuid.UUID
	MessageName   string
	MessageData   []byte
	OccurredAt    time.Time
	ProcessedAt   *time.Time
	FailCount     int32
	Failed        bool
	FailureReason *string
}

// listUnprocessedWithLockQuery returns a query to list unprocessed messages. With row-level locks for multi-threaded processing.
func listUnprocessedWithLockQuery(limit int32) (string, []any) {
	return `SELECT id, 	message_name, message_data, occurred_at, processed_at, fail_count, failed, failure_reason
	          FROM outbox
	         WHERE processed_at IS NULL
	           AND failed != true
	           AND occurred_at + (power(2, fail_count - 1) * interval '1 second') < now()
	         ORDER BY occurred_at
	         LIMIT $1
	           FOR UPDATE SKIP LOCKED`, []any{limit}
}

func upsertQuery(msg *Message) (string, []any) {
	var (
		processedAt   *time.Time
		failureReason *string
	)

	if msg.IsProcessed() {
		processedAt = msg.ProcessedAt()
	}

	if msg.FailureReason() != "" {
		reason := msg.FailureReason()
		failureReason = &reason
	}

	sql := `INSERT INTO outbox (id, 	message_name, 	message_data, occurred_at, processed_at, fail_count, failed, failure_reason)
	         VALUES ($1, $2, $3, $4, $5, $6, $7, $8)
	         ON CONFLICT (id) DO UPDATE SET
	               message_name   = EXCLUDED.message_name,
	               message_data   = EXCLUDED.message_data,
	               occurred_at    = EXCLUDED.occurred_at,
	               processed_at   = EXCLUDED.processed_at,
	               fail_count     = EXCLUDED.fail_count,
	               failed         = EXCLUDED.failed,
	               failure_reason = EXCLUDED.failure_reason`
	args := []any{
		msg.ID(),
		msg.MessageName(),
		msg.MessageData(),
		msg.OccurredAt(),
		processedAt,
		msg.FailCount(),
		msg.IsFailed(),
		failureReason,
	}
	return sql, args
}

func mapRowToMessage(row outboxRow) *Message {
	m := &Message{
		id:          row.ID,
		messageName: row.MessageName,
		messageData: row.MessageData,
		occurredAt:  row.OccurredAt,
		failCount:   row.FailCount,
		failed:      row.Failed,
	}

	if row.ProcessedAt != nil {
		m.processedAt = row.ProcessedAt
	}

	if row.FailureReason != nil {
		m.failureReason = *row.FailureReason
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
