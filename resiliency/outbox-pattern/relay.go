// Package relay provides functionality for relaying outbox messages from a database table to a message publisher.
package main

import (
	"context"
	"log/slog"
	"time"
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

// OutboxRelay is a worker that polls the outbox table for unprocessed messages and publishes them using the provided RelayMessage function. It marks messages as processed or failed based on the result of the publish operation.
type OutboxRelay struct {
	repository  Repository
	pollEvery   time.Duration
	publishFunc RelayMessage
}

// NewOutboxRelay creates a new OutboxRelay instance.
func NewOutboxRelay(publisher RelayMessage, repository Repository, options ...Option) *OutboxRelay {
	p := &OutboxRelay{
		repository:  repository,
		pollEvery:   1 * time.Second,
		publishFunc: publisher,
	}

	for _, opt := range options {
		opt(p)
	}

	return p
}

// Run starts the publisher loop and polls the outbox table for unprocessed messages.
func (p *OutboxRelay) Run(ctx context.Context) error {
	for {
		select {
		case <-time.After(p.pollEvery):
			err := p.repository.ProcessUnprocessedWithLock(ctx, 100, func(ctx context.Context, msg *Message) {
				if err := p.publishFunc(ctx, msg); err != nil {
					slog.ErrorContext(ctx, "failed to relay message", "message_id", msg.ID(), "error", err.Error())
					msg.MarkAsFailed(err.Error())
					return
				}

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
