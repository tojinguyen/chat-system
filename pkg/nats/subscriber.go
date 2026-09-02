package nats

import (
	"context"
	"encoding/json"
	"fmt"
	"log"

	"github.com/nats-io/nats.go"
)

// Handler represents a type-safe message handler function
type Handler[T any] func(ctx context.Context, data T) error

// Subscriber provides a generic, type-safe interface for consuming structured messages from NATS
type Subscriber[T any] struct {
	conn       *nats.Conn
	subject    string
	queueGroup string
	sub        *nats.Subscription
}

// NewSubscriber creates a new generic Subscriber for a subject (direct fan-out or unicast)
func NewSubscriber[T any](conn *nats.Conn, subject string) *Subscriber[T] {
	return &Subscriber[T]{
		conn:    conn,
		subject: subject,
	}
}

// NewQueueSubscriber creates a new generic Subscriber with a queue group for worker load balancing
func NewQueueSubscriber[T any](conn *nats.Conn, subject string, queueGroup string) *Subscriber[T] {
	return &Subscriber[T]{
		conn:       conn,
		subject:    subject,
		queueGroup: queueGroup,
	}
}

// Start begins consuming messages asynchronously and invokes the handler for each received message
func (s *Subscriber[T]) Start(ctx context.Context, handler Handler[T]) error {
	msgHandler := func(msg *nats.Msg) {
		var data T
		if err := json.Unmarshal(msg.Data, &data); err != nil {
			log.Printf("[NATS Subscriber] Failed to unmarshal message from %s: %v", s.subject, err)
			return
		}

		if err := handler(ctx, data); err != nil {
			log.Printf("[NATS Subscriber] Error handling message from %s: %v", s.subject, err)
		}
	}

	var sub *nats.Subscription
	var err error

	if s.queueGroup != "" {
		sub, err = s.conn.QueueSubscribe(s.subject, s.queueGroup, msgHandler)
	} else {
		sub, err = s.conn.Subscribe(s.subject, msgHandler)
	}

	if err != nil {
		return fmt.Errorf("failed to subscribe to subject %s: %w", s.subject, err)
	}

	s.sub = sub

	// Auto cleanup when context is canceled
	go func() {
		<-ctx.Done()
		_ = s.Close()
	}()

	return nil
}

// Close gracefully drains and removes the subscription
func (s *Subscriber[T]) Close() error {
	if s.sub != nil && s.sub.IsValid() {
		return s.sub.Drain()
	}
	return nil
}
