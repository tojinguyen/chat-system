package consumer

import (
	"context"
	"encoding/json"
	"fmt"
	"log"

	"chat-system/pkg/contracts"
	natsclient "chat-system/pkg/nats"
	"chat-system/pkg/telemetry"
	"chat-system/pkg/worker"
	"chat-worker/internal/usecase"

	"github.com/nats-io/nats.go"
	"go.opentelemetry.io/otel/attribute"
	"go.opentelemetry.io/otel/trace"
)

// Consumer wraps NATS Queue Subscriber and dispatches to a generic Key-based Partitioned Worker Pool
type Consumer struct {
	subscriber  *natsclient.Subscriber[contracts.InboundBrokerEvent]
	chatUsecase usecase.ChatUsecase
	pool        *worker.PartitionedPool[contracts.InboundBrokerEvent]
}

// Config chứa cấu hình cần thiết để khởi tạo Consumer
type Config struct {
	NC          *nats.Conn
	Subject     string
	QueueGroup  string
	NumWorkers  int
	BufferSize  int
	ChatUsecase usecase.ChatUsecase
}

// NewConsumer creates a new Inbound Consumer with generic partitioned worker pool support
func NewConsumer(cfg Config) (*Consumer, error) {
	c := &Consumer{
		subscriber:  natsclient.NewQueueSubscriber[contracts.InboundBrokerEvent](cfg.NC, cfg.Subject, cfg.QueueGroup),
		chatUsecase: cfg.ChatUsecase,
	}

	pool, err := worker.NewPartitionedPool(worker.Config[contracts.InboundBrokerEvent]{
		NumWorkers:   cfg.NumWorkers,
		BufferSize:   cfg.BufferSize,
		KeyExtractor: extractConversationKey,
		Handler:      c.handleWorkerProcess,
	})
	if err != nil {
		return nil, fmt.Errorf("failed to create worker pool: %w", err)
	}

	c.pool = pool
	return c, nil
}

// Start begins consuming messages and starting worker goroutines
func (c *Consumer) Start(ctx context.Context) error {
	// 1. Khởi động worker pool
	c.pool.Start()

	// 2. Wrap handler với telemetry stage "worker_dispatch"
	instrumentedHandler := telemetry.InstrumentHandler(telemetry.HandlerConfig{
		Service:   "chat-engine",
		Mode:      "engine",
		Stage:     "worker_dispatch",
		EventType: "inbound",
	}, c.dispatchMessage)

	return c.subscriber.Start(ctx, instrumentedHandler)
}

// dispatchMessage trích xuất span context và định tuyến message vào worker channel
func (c *Consumer) dispatchMessage(ctx context.Context, event contracts.InboundBrokerEvent) error {
	if err := c.pool.Submit(ctx, event); err != nil {
		log.Printf("[Consumer] Failed to submit message to worker pool msg_id=%s: %v", event.ClientMsgID, err)
		return err
	}
	return nil
}

// handleWorkerProcess được thực thi bên trong Worker Goroutine được phân bổ
func (c *Consumer) handleWorkerProcess(ctx context.Context, event contracts.InboundBrokerEvent) error {
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

// Stop gracefully closes NATS subscription and drains the worker pool
func (c *Consumer) Stop(drainCtx context.Context) error {
	log.Println("[Consumer] Stopping NATS subscription...")
	if err := c.subscriber.Close(); err != nil {
		log.Printf("[Consumer] Warning: error closing NATS subscriber: %v", err)
	}

	log.Println("[Consumer] Draining worker pool...")
	return c.pool.Stop(drainCtx)
}

// extractConversationKey trích xuất nhanh conversation_id để băm định tuyến
func extractConversationKey(event contracts.InboundBrokerEvent) string {
	type quickPayload struct {
		ConversationID string `json:"conversation_id"`
	}

	var pLoad quickPayload
	if len(event.Payload) > 0 {
		if err := json.Unmarshal(event.Payload, &pLoad); err == nil && pLoad.ConversationID != "" {
			return pLoad.ConversationID
		}
	}

	if event.ClientMsgID != "" {
		return event.ClientMsgID
	}
	return event.SenderID
}
