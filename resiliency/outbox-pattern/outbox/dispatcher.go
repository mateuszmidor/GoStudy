// Package dispatcher provides functionality for dispatching outbox messages from a database table to a handler.
package outbox

import (
	"context"
	"log/slog"
	"time"
)

// DispatchFunc is a function that handles an outbox message after it's received from the outbox table
type DispatchFunc func(ctx context.Context, msg *Message) error

// Option is a function that modifies a Dispatcher based on options
type Option func(*Dispatcher)

// WithPollingRate sets the rate at which the dispatcher polls the outbox table for unprocessed messages
func WithPollingRate(pollEvery time.Duration) Option {
	return func(p *Dispatcher) {
		p.pollEvery = pollEvery
	}
}

// WithMaxAttempts sets the maximum number of processing attempts for a message
func WithMaxAttempts(maxAttempts int32) Option {
	return func(p *Dispatcher) {
		p.maxAttempts = maxAttempts
	}
}

// Dispatcher is a worker that polls the outbox table for unprocessed messages and dispatches them using the provided DispatchFunc function.
// It marks messages as processed or failed based on the result of the dispatch operation.
type Dispatcher struct {
	repository  *Repository
	pollEvery   time.Duration
	maxAttempts int32
}

// NewDispatcher creates a new Dispatcher instance.
func NewDispatcher(repository *Repository, options ...Option) *Dispatcher {
	d := &Dispatcher{
		repository: repository,
		pollEvery:  1 * time.Second,
	}

	for _, opt := range options {
		opt(d)
	}

	return d
}

// Run starts the dispatcher loop and polls the outbox table for unprocessed messages.
func (d *Dispatcher) Run(ctx context.Context, dispatchFunc DispatchFunc) error {
	dispatchWithRetry := func(ctx context.Context, msg *Message) {
		if err := dispatchFunc(ctx, msg); err != nil {
			slog.ErrorContext(ctx, "failed to dispatch message", "message_id", msg.ID(), "error", err.Error())

			if msg.FailCount()+1 >= d.maxAttempts {
				slog.ErrorContext(ctx, "message exceeded max retries, abandoning", "message_id", msg.ID(), "error", err.Error())
				msg.MarkAsFailed(err.Error())
			} else {
				msg.AddFailure(err.Error())
			}
		} else {
			msg.MarkAsProcessed()
		}
	}

	for {
		select {
		case <-time.After(d.pollEvery):
			if err := d.repository.ProcessUnprocessedWithLock(ctx, 100, dispatchWithRetry); err != nil {
				return err
			}

		case <-ctx.Done():
			return nil
		}
	}
}
