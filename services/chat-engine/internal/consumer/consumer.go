package consumer

import (
	"context"
	"chat-worker/internal/domain"
)

type MessageHandler func(ctx context.Context, event *domain.InboundMessageEvent) error

// InboundConsumer consumes messages from NATS Inbound Subject / JetStream Stream
type InboundConsumer interface {
	Start(ctx context.Context, handler MessageHandler) error
	Close() error
}

type natsJetStreamConsumer struct {
	// TODO: nats.Conn / JetStream Consumer reference
}

func NewInboundConsumer(natsURL string, subject string, queueGroup string) (InboundConsumer, error) {
	return &natsJetStreamConsumer{}, nil
}

func (c *natsJetStreamConsumer) Start(ctx context.Context, handler MessageHandler) error {
	// TODO: natsConn.QueueSubscribe(subject, queueGroup, ...) or JetStream Consume
	return nil
}

func (c *natsJetStreamConsumer) Close() error {
	return nil
}
