package nats

import (
	"context"
	"encoding/json"
	"fmt"

	"github.com/nats-io/nats.go"
)

// Publisher provides a generic, type-safe interface for publishing structured messages to NATS
type Publisher[T any] struct {
	conn           *nats.Conn
	defaultSubject string
}

// NewPublisher creates a new generic Publisher for the specified default subject
func NewPublisher[T any](conn *nats.Conn, defaultSubject string) *Publisher[T] {
	return &Publisher[T]{
		conn:           conn,
		defaultSubject: defaultSubject,
	}
}

// Publish serializes data to JSON and publishes it to the default subject
func (p *Publisher[T]) Publish(ctx context.Context, data T) error {
	return p.PublishToSubject(ctx, p.defaultSubject, data)
}

// PublishToSubject serializes data to JSON and publishes it to a specific custom subject
func (p *Publisher[T]) PublishToSubject(ctx context.Context, subject string, data T) error {
	if err := ctx.Err(); err != nil {
		return fmt.Errorf("context canceled before publishing: %w", err)
	}

	payload, err := json.Marshal(data)
	if err != nil {
		return fmt.Errorf("failed to marshal message for subject %s: %w", subject, err)
	}

	if err := p.conn.Publish(subject, payload); err != nil {
		return fmt.Errorf("failed to publish message to subject %s: %w", subject, err)
	}

	return nil
}
