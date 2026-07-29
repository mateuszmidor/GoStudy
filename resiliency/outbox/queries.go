package main

import (
	"time"

	"github.com/google/uuid"
	"github.com/stephenafamo/bob"
	"github.com/stephenafamo/bob/dialect/psql"
	"github.com/stephenafamo/bob/dialect/psql/im"
	"github.com/stephenafamo/bob/dialect/psql/sm"
)

const (
	tableOutbox = "outbox"

	idCol            = "id"
	eventNameCol     = "event_name"
	eventDataCol     = "event_data"
	occurredAtCol    = "occurred_at"
	processedAtCol   = "processed_at"
	failCountCol     = "fail_count"
	failedCol        = "failed"
	failureReasonCol = "failure_reason"
	traceIDCol       = "trace_id"
)

var databaseColumns = []any{
	idCol,
	eventNameCol,
	eventDataCol,
	occurredAtCol,
	processedAtCol,
	failCountCol,
	failedCol,
	failureReasonCol,
	traceIDCol,
}

// outboxRow represents the database row structure for outbox table
type outboxRow struct {
	ID            uuid.UUID  `db:"id"`
	EventName     string     `db:"event_name"`
	EventData     []byte     `db:"event_data"`
	OccurredAt    time.Time  `db:"occurred_at"`
	ProcessedAt   *time.Time `db:"processed_at"`
	FailCount     int32      `db:"fail_count"`
	Failed        bool       `db:"failed"`
	FailureReason *string    `db:"failure_reason"`
	TraceID       *string    `db:"trace_id"`
}

// listAllQuery returns all messages from the outbox ordered by occurred_at
func listAllQuery() bob.Query {
	return psql.Select(
		sm.Columns(databaseColumns...),
		sm.From(tableOutbox),
		sm.OrderBy(occurredAtCol),
	)
}

// listUnprocessedQuery returns unprocessed and non-failed messages
func listUnprocessedQuery(limit int32) bob.Query {
	return psql.Select(
		sm.Columns(databaseColumns...),
		sm.From(tableOutbox),
		sm.Where(psql.Quote(processedAtCol).IsNull()),
		sm.Where(psql.Quote(failedCol).NE(psql.Arg(true))),
		sm.OrderBy(occurredAtCol),
		sm.Limit(int(limit)),
	)
}

// listUnprocessedWithLockQuery returns unprocessed and non-failed messages with row-level lock
func listUnprocessedWithLockQuery(limit int32) bob.Query {
	return psql.Select(
		sm.Columns(databaseColumns...),
		sm.From(tableOutbox),
		sm.Where(psql.Quote(processedAtCol).IsNull()),
		sm.Where(psql.Quote(failedCol).NE(psql.Arg(true))),
		sm.OrderBy(occurredAtCol),
		sm.Limit(int(limit)),
		sm.ForUpdate().SkipLocked(),
	)
}

// upsertQuery creates an upsert query for a message
func upsertQuery(message Message) bob.Query {

	var (
		processedAt   *time.Time
		failureReason *string
		traceID       *string
	)

	if message.IsProcessed() {
		processedAt = message.ProcessedAt()
	}

	if message.FailureReason() != "" {
		reason := message.FailureReason()
		failureReason = &reason
	}

	if message.TraceID() != "" {
		trace := message.TraceID()
		traceID = &trace
	}

	return psql.Insert(
		im.Into(
			tableOutbox,
			idCol,
			eventNameCol,
			eventDataCol,
			occurredAtCol,
			processedAtCol,
			failCountCol,
			failedCol,
			failureReasonCol,
			traceIDCol,
		),
		im.Values(psql.Arg(
			message.ID(),
			message.EventName(),
			message.EventData(),
			message.OccurredAt(),
			processedAt,
			message.FailCount(),
			message.IsFailed(),
			failureReason,
			traceID,
		)),
		im.OnConflict(idCol).DoUpdate(
			im.SetExcluded(eventNameCol),
			im.SetExcluded(eventDataCol),
			im.SetExcluded(occurredAtCol),
			im.SetExcluded(processedAtCol),
			im.SetExcluded(failCountCol),
			im.SetExcluded(failedCol),
			im.SetExcluded(failureReasonCol),
			im.SetExcluded(traceIDCol),
		),
	)
}

// mapRowToMessage converts a database row to a Message
func mapRowToMessage(row outboxRow) Message {
	m := &message{
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

// mapRowsToMessages converts multiple database rows to Messages
func mapRowsToMessages(rows []outboxRow) []Message {
	messages := make([]Message, len(rows))
	for i, row := range rows {
		messages[i] = mapRowToMessage(row)
	}
	return messages
}
