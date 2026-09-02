package broker

import (
	"context"
	"encoding/json"
	"fmt"

	"chat-system/pkg/contracts"
	natsclient "chat-system/pkg/nats"

	"github.com/nats-io/nats.go"
)

// OutboundHandler is the callback function when a message arrives from NATS for this gateway node
type OutboundHandler func(event *contracts.OutboundBrokerEvent) error

// OutboundConsumer subscribes to messages destined for this specific gateway node subject (e.g. chat.gateway.gateway-node-01)
type OutboundConsumer interface {
	Start(ctx context.Context, handler OutboundHandler) error
	Close() error
}

type natsConsumer struct {
	conn    *nats.Conn
	subject string
	sub     *nats.Subscription
}

func NewOutboundConsumer(natsURL string, subject string) (OutboundConsumer, error) {
	nc, err := natsclient.Connect(natsURL, "ws-gateway-outbound-consumer")
	if err != nil {
		return nil, fmt.Errorf("failed to connect to NATS: %w", err)
	}
	return &natsConsumer{
		conn:    nc,
		subject: subject,
	}, nil
}

func (c *natsConsumer) Start(ctx context.Context, handler OutboundHandler) error {
	sub, err := c.conn.Subscribe(c.subject, func(msg *nats.Msg) {
		var event contracts.OutboundBrokerEvent
		if err := json.Unmarshal(msg.Data, &event); err != nil {
			fmt.Printf("failed to unmarshal message: %v\n", err)
			return
		}
		if err := handler(&event); err != nil {
			fmt.Printf("failed to handle message: %v\n", err)
		}
	})
	if err != nil {
		return fmt.Errorf("failed to subscribe to subject %s: %w", c.subject, err)
	}

	c.sub = sub

	go func() {
		<-ctx.Done()
		c.Close()
	}()

	return nil
}

func (c *natsConsumer) Close() error {
	if c.sub != nil {
		_ = c.sub.Drain()
	}

	if c.conn != nil && !c.conn.IsClosed() {
		c.conn.Close()
	}
	return nil
}
