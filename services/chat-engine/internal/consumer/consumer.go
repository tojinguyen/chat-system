package consumer

import (
	"context"
	"log"

	"chat-system/pkg/contracts"
	natsclient "chat-system/pkg/nats"
	"chat-system/pkg/telemetry"
	"chat-worker/internal/usecase"

	"github.com/nats-io/nats.go"
	"go.opentelemetry.io/otel/attribute"
	"go.opentelemetry.io/otel/trace"
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
	instrumentedHandler := telemetry.InstrumentHandler(telemetry.HandlerConfig{
		Service:   "chat-engine",
		Mode:      "engine",
		Stage:     "worker_process",
		EventType: "inbound",
	}, c.handleMessage)
	return c.subscriber.Start(ctx, instrumentedHandler)
}

// handleMessage processes inbound broker events
func (c *Consumer) handleMessage(ctx context.Context, event contracts.InboundBrokerEvent) error {
	tracer := telemetry.Tracer("chat-engine")
	spanCtx, span := tracer.Start(ctx, "chat_engine.process",
		trace.WithAttributes(
			attribute.String("client_msg_id", event.ClientMsgID),
			attribute.String("sender_id", event.SenderID),
			attribute.String("gateway_node", event.GatewayNode),
		),
	)
	defer span.End()

	if err := c.chatUsecase.ProcessInboundMessage(spanCtx, event); err != nil {
		log.Printf("[Consumer] Error processing event client_msg_id=%s: %v", event.ClientMsgID, err)
		return err
	}
	return nil
}

// Stop gracefully closes and drains the NATS subscription
func (c *Consumer) Stop() error {
	return c.subscriber.Close()
}
