package broker

import (
	"context"
	"ws-gateway/internal/domain"
)

// OutboundHandler is the callback function when a message arrives from NATS for this gateway node
type OutboundHandler func(event *domain.OutboundBrokerEvent) error

// OutboundConsumer subscribes to messages destined for this specific gateway node subject (e.g. chat.gateway.gateway-node-01)
type OutboundConsumer interface {
	Start(ctx context.Context, handler OutboundHandler) error
	Close() error
}

type natsConsumer struct {
	// TODO: nats.Conn / Subscription reference
}

func NewOutboundConsumer(natsURL string, subject string) (OutboundConsumer, error) {
	return &natsConsumer{}, nil
}

func (c *natsConsumer) Start(ctx context.Context, handler OutboundHandler) error {
	// TODO: natsConn.Subscribe(subject, func(msg *nats.Msg) { ... })
	return nil
}

func (c *natsConsumer) Close() error {
	return nil
}
