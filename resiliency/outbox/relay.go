// Package relay provides functionality for relaying outbox messages from a database table to a message publisher.
package main

import (
	"context"
	"log/slog"
	"time"

	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/attribute"
	"go.opentelemetry.io/otel/trace"
)

// RelayMessage is a function that relays an outbox message after its received from the outbox table
type RelayMessage func(ctx context.Context, msg *Message) error

// Option is a function that modifies a OutboxRelay based on options
type Option func(*OutboxRelay)

// WithPollingRate sets the rate at which the publisher polls the outbox table for unprocessed messages
func WithPollingRate(pollEvery time.Duration) Option {
	return func(p *OutboxRelay) {
		p.pollEvery = pollEvery
	}
}

// WithTracingServiceName set an alternative name for the tracing service used by the OutboxRelay. This is useful when you want to distinguish between different relays in your tracing system.
func WithTracingServiceName(serviceName string) Option {
	return func(p *OutboxRelay) {
		p.serviceName = serviceName
	}
}

// OutboxRelay is a worker that polls the outbox table for unprocessed messages and publishes them using the provided RelayMessage function. It marks messages as processed or failed based on the result of the publish operation.
type OutboxRelay struct {
	repository  Repository
	pollEvery   time.Duration
	publish     RelayMessage
	serviceName string
}

// NewOutboxRelay creates a new OutboxRelay instance.
func NewOutboxRelay(publisher RelayMessage, repository Repository, options ...Option) *OutboxRelay {
	p := &OutboxRelay{
		repository:  repository,
		pollEvery:   1 * time.Second,
		publish:     publisher,
		serviceName: "outbox.relay",
	}

	for _, opt := range options {
		opt(p)
	}

	return p
}

// Run starts the publisher loop and polls the outbox table for unprocessed messages.
func (p *OutboxRelay) Run(ctx context.Context) error {
	tracer := otel.Tracer(p.serviceName)
	for {
		select {
		case <-time.After(p.pollEvery):
			err := p.repository.ProcessUnprocessedWithLock(ctx, 100, func(ctx context.Context, msg *Message) {
				ctx, span := tracer.Start(resumeContextWithTraceID(ctx, msg.TraceID()),
					"outbox.relay.publish",
					trace.WithAttributes(
						attribute.String("message.id", msg.ID().String()),
						attribute.String("message.event_name", msg.EventName()),
					))
				defer span.End()

				if err := p.publish(ctx, msg); err != nil {
					slog.ErrorContext(ctx, "failed to relay message", "message_id", msg.ID(), "error", err.Error())
					span.RecordError(err)

					msg.MarkAsFailed(err.Error())
					return
				}

				span.AddEvent("published message")
				msg.MarkAsProcessed()
			})

			if err != nil {
				return err
			}

		case <-ctx.Done():
			return nil
		}
	}
}

func resumeContextWithTraceID(ctx context.Context, traceID string) context.Context {
	msgCtx := ctx
	if traceID != "" {
		traceID, err := trace.TraceIDFromHex(traceID)
		if err == nil {
			spanContext := trace.NewSpanContext(trace.SpanContextConfig{
				TraceID:    traceID,
				SpanID:     trace.SpanID{},
				TraceFlags: trace.FlagsSampled,
			})
			msgCtx = trace.ContextWithSpanContext(ctx, spanContext)
		}
	}
	return msgCtx
}
