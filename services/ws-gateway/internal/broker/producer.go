package broker

import (
	"context"
	"ws-gateway/internal/domain"
)

// InboundProducer publishes incoming client messages to NATS Inbound Subject
type InboundProducer interface {
	PublishInbound(ctx context.Context, event *domain.InboundBrokerEvent) error
	Close() error
}

type natsProducer struct {
	// TODO: nats.Conn or jetstream.JetStream reference
}

func NewInboundProducer(natsURL string, subject string) (InboundProducer, error) {
	return &natsProducer{}, nil
}

func (p *natsProducer) PublishInbound(ctx context.Context, event *domain.InboundBrokerEvent) error {
	// TODO: natsConn.Publish(subject, data) or js.Publish(ctx, subject, data)
	return nil
}

func (p *natsProducer) Close() error {
	return nil
}
