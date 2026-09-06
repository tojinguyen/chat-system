package consumer

import (
	"chat-worker/internal/usecase"
	"context"
	"log"

	"chat-system/pkg/contracts"
	natsclient "chat-system/pkg/nats"

	"github.com/nats-io/nats.go"
)

// Consumer wraps NATS Queue Subscriber to consume inbound events from broker
type Consumer struct {
	subscriber  *natsclient.Subscriber[contracts.InboundBrokerEvent]
	chatUsecase usecase.ChatUsecase
}

// NewConsumer creates a new Inbound Consumer with queue group support for worker load balancing
func NewConsumer(nc *nats.Conn, subject, queueGroup string, chatUsecase usecase.ChatUsecase) *Consumer {
	return &Consumer{
		subscriber:  natsclient.NewQueueSubscriber[contracts.InboundBrokerEvent](nc, subject, queueGroup),
		chatUsecase: chatUsecase,
	}
}

// Start begins consuming messages asynchronously
func (c *Consumer) Start(ctx context.Context) error {
	return c.subscriber.Start(ctx, c.handleMessage)
}

// handleMessage processes inbound broker events
func (c *Consumer) handleMessage(ctx context.Context, event contracts.InboundBrokerEvent) error {
	if err := c.chatUsecase.ProcessInboundMessage(ctx, event); err != nil {
		log.Printf("[Consumer] Error processing event client_msg_id=%s: %v", event.ClientMsgID, err)
		return err
	}
	return nil
}

// Stop gracefully closes and drains the NATS subscription
func (c *Consumer) Stop() error {
	return c.subscriber.Close()
}
