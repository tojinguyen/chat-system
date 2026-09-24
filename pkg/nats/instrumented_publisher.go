package nats

import (
	"context"
	"time"

	"chat-system/pkg/telemetry"

	"go.opentelemetry.io/otel/attribute"
	"go.opentelemetry.io/otel/trace"
)

// PublisherConfig chứa cấu hình ngữ cảnh đo lường cho Publisher
type PublisherConfig struct {
	ServiceName string // ví dụ: "ws-gateway", "chat-engine"
	Stage       string // ví dụ: "inbound", "outbound_dispatch"
	EventType   string // ví dụ: "inbound", "outbound"
}

// InstrumentedPublisher là Decorator bọc quanh Publisher[T], tự động thu thập Trace & Metrics
type InstrumentedPublisher[T any] struct {
	inner *Publisher[T]
	cfg   PublisherConfig
}

// NewInstrumentedPublisher khởi tạo một InstrumentedPublisher bọc ngoài Publisher[T]
func NewInstrumentedPublisher[T any](inner *Publisher[T], cfg PublisherConfig) *InstrumentedPublisher[T] {
	if cfg.EventType == "" {
		cfg.EventType = cfg.Stage
	}
	return &InstrumentedPublisher[T]{
		inner: inner,
		cfg:   cfg,
	}
}

// Publish gửi dữ liệu vào default subject, tự động đo latency và message count
func (p *InstrumentedPublisher[T]) Publish(ctx context.Context, data T) error {
	start := time.Now()
	err := p.inner.Publish(ctx, data)
	duration := time.Since(start).Seconds()

	status := "success"
	if err != nil {
		status = "error"
	}

	telemetry.MessageLatency.WithLabelValues("nats", p.cfg.Stage, status).Observe(duration)
	telemetry.MessagesProcessed.WithLabelValues(p.cfg.ServiceName, p.cfg.EventType, status).Inc()
	return err
}

// PublishToSubject gửi dữ liệu vào một subject cụ thể, tự động đo latency và message count
func (p *InstrumentedPublisher[T]) PublishToSubject(ctx context.Context, subject string, data T) error {
	start := time.Now()
	err := p.inner.PublishToSubject(ctx, subject, data)
	duration := time.Since(start).Seconds()

	status := "success"
	if err != nil {
		status = "error"
	}

	telemetry.MessageLatency.WithLabelValues("nats", p.cfg.Stage, status).Observe(duration)
	telemetry.MessagesProcessed.WithLabelValues(p.cfg.ServiceName, p.cfg.EventType, status).Inc()
	return err
}

// PublishToNode gửi dữ liệu tới subject của một Gateway Node cụ thể, tự động tạo Span và đo cả DispatchDuration lẫn MessageLatency
func (p *InstrumentedPublisher[T]) PublishToNode(ctx context.Context, nodeID, subject string, data T) error {
	tracer := telemetry.Tracer(p.cfg.ServiceName)
	spanCtx, span := tracer.Start(ctx, "nats.DispatchToGateway",
		trace.WithAttributes(
			attribute.String("gateway_node", nodeID),
			attribute.String("subject", subject),
		),
	)
	defer span.End()

	start := time.Now()
	err := p.inner.PublishToSubject(spanCtx, subject, data)
	duration := time.Since(start).Seconds()

	status := "success"
	if err != nil {
		status = "error"
	}

	telemetry.DispatchDuration.WithLabelValues("nats", nodeID, status).Observe(duration)
	telemetry.MessageLatency.WithLabelValues("nats", p.cfg.Stage, status).Observe(duration)
	telemetry.MessagesProcessed.WithLabelValues(p.cfg.ServiceName, p.cfg.EventType, status).Inc()

	return err
}
