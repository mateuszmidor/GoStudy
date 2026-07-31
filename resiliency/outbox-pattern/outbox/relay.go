// Package relay provides functionality for relaying outbox messages from a database table to a message publisher.
package outbox

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

// WithMaxAttempts sets the maximum number of processing attempts for a message
func WithMaxAttempts(maxAttempts int32) Option {
	return func(p *OutboxRelay) {
		p.maxAttempts = maxAttempts
	}
}

// OutboxRelay is a worker that polls the outbox table for unprocessed messages and publishes them using the provided RelayMessage function. It marks messages as processed or failed based on the result of the publish operation.
type OutboxRelay struct {
	repository  *Repository
	pollEvery   time.Duration
	maxAttempts int32
}

// NewOutboxRelay creates a new OutboxRelay instance.
func NewOutboxRelay(repository *Repository, options ...Option) *OutboxRelay {
	p := &OutboxRelay{
		repository: repository,
		pollEvery:  1 * time.Second,
	}

	for _, opt := range options {
		opt(p)
	}

	return p
}

// Run starts the publisher loop and polls the outbox table for unprocessed messages.
func (p *OutboxRelay) Run(ctx context.Context, publishFunc RelayMessage) error {
	// publish single message, use with retry mechanism
	publishWithRetry := func(ctx context.Context, msg *Message) {
		// try to publish...
		if err := publishFunc(ctx, msg); err != nil {
			// 1. failure, retry if any attempts left
			slog.ErrorContext(ctx, "failed to relay message", "message_id", msg.ID(), "error", err.Error())

			if msg.FailCount()+1 >= p.maxAttempts {
				slog.ErrorContext(ctx, "message exceeded max retries, abandoning", "message_id", msg.ID(), "error", err.Error())
				msg.MarkAsFailed(err.Error())
			} else {
				msg.AddFailure(err.Error())
			}
		} else {
			// 2. success
			msg.MarkAsProcessed()
		}
	}

	// relay loop
	for {
		select {
		case <-time.After(p.pollEvery):
			if err := p.repository.ProcessUnprocessedWithLock(ctx, 100, publishWithRetry); err != nil {
				return err
			}

		case <-ctx.Done():
			return nil
		}
	}
}
