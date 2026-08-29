package broker

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"time"
	"ws-gateway/internal/domain"

	"github.com/nats-io/nats.go"
)

// InboundProducer publishes incoming client messages to NATS Inbound Subject
type InboundProducer interface {
	PublishInbound(ctx context.Context, event *domain.InboundBrokerEvent) error
	Close() error
}

type natsProducer struct {
	conn    *nats.Conn
	subject string
}

func connectNATS(natsURL string, clientName string) (*nats.Conn, error) {
	opts := []nats.Option{
		nats.Name(clientName),
		nats.MaxReconnects(-1),
		nats.ReconnectWait(2 * time.Second),
		nats.DisconnectErrHandler(func(nc *nats.Conn, err error) {
			log.Printf("[NATS] Disconnected: %v", err)
		}),
		nats.ReconnectHandler(func(nc *nats.Conn) {
			log.Printf("[NATS] Reconnected to %s", nc.ConnectedUrl())
		}),
		nats.ClosedHandler(func(nc *nats.Conn) {
			log.Printf("[NATS] Connection closed permanently")
		}),
	}
	return nats.Connect(natsURL, opts...)
}

func NewInboundProducer(natsURL string, subject string) (InboundProducer, error) {
	nc, err := connectNATS(natsURL, "ws-gateway-inbound-producer")
	if err != nil {
		return nil, fmt.Errorf("failed to connect to NATS: %w", err)
	}
	return &natsProducer{
		conn:    nc,
		subject: subject,
	}, nil
}

func (p *natsProducer) PublishInbound(ctx context.Context, event *domain.InboundBrokerEvent) error {
	if err := ctx.Err(); err != nil {
		return fmt.Errorf("context canceled before publish: %w", err)
	}

	data, err := json.Marshal(event)
	if err != nil {
		return fmt.Errorf("failed to marshal event: %w", err)
	}

	if err := p.conn.Publish(p.subject, data); err != nil {
		return fmt.Errorf("failed to publish message to NATS: %w", err)
	}
	return nil
}

func (p *natsProducer) Close() error {
	if p.conn != nil && !p.conn.IsClosed() {
		p.conn.Close()
	}
	return nil
}
