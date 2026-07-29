package main

import (
	"errors"
	"fmt"
	"time"

	cloudevents "github.com/cloudevents/sdk-go/v2"
	"github.com/google/uuid"
)

type Message struct {
	id            uuid.UUID
	eventName     string
	eventData     []byte
	occurredAt    time.Time
	processedAt   *time.Time
	failCount     int32
	failed        bool
	failureReason string
	traceID       string
}

// NewMessage creates a new outbox message.
func NewMessage(id uuid.UUID, eventName string, eventData []byte, occurredAt time.Time, traceID string) *Message {
	return &Message{
		id:            id,
		eventName:     eventName,
		eventData:     eventData,
		occurredAt:    occurredAt,
		processedAt:   nil,
		failCount:     0,
		failed:        false,
		failureReason: "",
		traceID:       traceID,
	}
}

func (m *Message) String() string {
	return fmt.Sprintf("Message{ID: %s, EventName: %s, EventData: %s, OccurredAt: %s, FailCount: %d}",
		m.id, m.eventName, m.eventData, m.occurredAt.Format(time.DateTime), m.failCount)
}

// ID returns the unique identifier of the message.
func (m *Message) ID() uuid.UUID {
	return m.id
}

// EventName returns the name of the event that triggered the message.
func (m *Message) EventName() string {
	return m.eventName
}

// EventData returns the event data that triggered the message.
func (m *Message) EventData() []byte {
	return m.eventData
}

// OccurredAt returns the time when the message was created.
func (m *Message) OccurredAt() time.Time {
	return m.occurredAt
}

// ProcessedAt returns the time when the message was processed.
func (m *Message) ProcessedAt() *time.Time {
	return m.processedAt
}

// IsProcessed returns true if the message has been processed.
func (m *Message) IsProcessed() bool {
	return m.processedAt != nil
}

// FailCount returns the number of times the message has failed.
func (m *Message) FailCount() int32 {
	return m.failCount
}

// IsFailed returns true if the message has failed.
func (m *Message) IsFailed() bool {
	return m.failed
}

// FailureReason returns the reason for the failure.
func (m *Message) FailureReason() string {
	return m.failureReason
}

// TraceID returns the trace ID of the message.
func (m *Message) TraceID() string {
	return m.traceID
}

// AddFailure increments the failure count and sets the failure reason.
func (m *Message) AddFailure(reason string) *Message {
	m.failCount++
	m.failureReason = reason
	return m
}

// MarkAsFailed marks the message as failed and sets the failure reason.
func (m *Message) MarkAsFailed(reason string) *Message {
	m.AddFailure(reason)
	m.failed = true
	return m
}

// MarkAsProcessed marks the message as processed.
func (m *Message) MarkAsProcessed() *Message {
	now := time.Now()
	m.processedAt = &now
	return m
}

// ToCloudEvent converts the outbox message to a CloudEvent format, including the trace ID as an extension.
func (m *Message) ToCloudEvent() (cloudevents.Event, error) {
	event := cloudevents.NewEvent()
	event.SetID(m.ID().String())
	event.SetType(m.EventName())
	event.SetTime(m.OccurredAt())
	event.SetExtension("traceID", m.TraceID())

	if err := event.SetData(cloudevents.ApplicationJSON, m.EventData()); err != nil {
		return event, errors.Join(errors.New("failed to set CloudEvent data"), err)
	}

	return event, nil

}
