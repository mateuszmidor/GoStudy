package main

import (
	"errors"
	"fmt"
	"time"

	cloudevents "github.com/cloudevents/sdk-go/v2"
	"github.com/google/uuid"
)

// Message is an outbox message intended to be stored alongside the storing the aggregate in the database
type Message interface {
	ID() uuid.UUID
	EventName() string
	EventData() []byte
	OccurredAt() time.Time
	ProcessedAt() *time.Time
	IsProcessed() bool
	FailCount() int32
	IsFailed() bool
	FailureReason() string
	TraceID() string
	ToCloudEvent() (cloudevents.Event, error)

	AddFailure(reason string) Message
	MarkAsFailed(reason string) Message
	MarkAsProcessed() Message

	String() string
}

type message struct {
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
func NewMessage(id uuid.UUID, eventName string, eventData []byte, occurredAt time.Time, traceID string) Message {
	return &message{
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

func (m *message) String() string {
	return fmt.Sprintf("Message{ID: %s, EventName: %s, EventData: %s, OccurredAt: %s, ProcessedAt: %s, FailCount: %d, Failed: %t, FailureReason: %s, TraceID: %s}",
		m.id, m.eventName, m.eventData, m.occurredAt, m.processedAt, m.failCount, m.failed, m.failureReason, m.traceID)
}

// ID returns the unique identifier of the message.
func (m *message) ID() uuid.UUID {
	return m.id
}

// EventName returns the name of the event that triggered the message.
func (m *message) EventName() string {
	return m.eventName
}

// EventData returns the event data that triggered the message.
func (m *message) EventData() []byte {
	return m.eventData
}

// OccurredAt returns the time when the message was created.
func (m *message) OccurredAt() time.Time {
	return m.occurredAt
}

// ProcessedAt returns the time when the message was processed.
func (m *message) ProcessedAt() *time.Time {
	return m.processedAt
}

// IsProcessed returns true if the message has been processed.
func (m *message) IsProcessed() bool {
	return m.processedAt != nil
}

// FailCount returns the number of times the message has failed.
func (m *message) FailCount() int32 {
	return m.failCount
}

// IsFailed returns true if the message has failed.
func (m *message) IsFailed() bool {
	return m.failed
}

// FailureReason returns the reason for the failure.
func (m *message) FailureReason() string {
	return m.failureReason
}

// TraceID returns the trace ID of the message.
func (m *message) TraceID() string {
	return m.traceID
}

// AddFailure increments the failure count and sets the failure reason.
func (m *message) AddFailure(reason string) Message {
	m.failCount++
	m.failureReason = reason
	return m
}

// MarkAsFailed marks the message as failed and sets the failure reason.
func (m *message) MarkAsFailed(reason string) Message {
	m.AddFailure(reason)
	m.failed = true
	return m
}

// MarkAsProcessed marks the message as processed.
func (m *message) MarkAsProcessed() Message {
	now := time.Now()
	m.processedAt = &now
	return m
}

// ToCloudEvent converts the outbox message to a CloudEvent format, including the trace ID as an extension.
func (m *message) ToCloudEvent() (cloudevents.Event, error) {
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
