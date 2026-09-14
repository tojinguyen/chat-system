package telemetry

import (
	"context"
	"time"
)

// HandlerConfig chứa cấu hình đo lường cho các message handler
type HandlerConfig struct {
	Service   string // ví dụ: "chat-engine", để trống nếu không đếm MessagesProcessed
	Mode      string // ví dụ: "engine", "nats", "grpc"
	Stage     string // ví dụ: "worker_process", "gateway_delivery"
	EventType string // ví dụ: "inbound", mặc định lấy theo Stage nếu để trống
}

// InstrumentHandler bọc ngoài hàm xử lý message (ví dụ Subscriber callback hoặc Consumer handler)
// để tự động đo đạc MessageLatency và MessagesProcessed mà không làm ô nhiễm business logic
func InstrumentHandler[T any](cfg HandlerConfig, handler func(ctx context.Context, data T) error) func(ctx context.Context, data T) error {
	if cfg.EventType == "" {
		cfg.EventType = cfg.Stage
	}
	return func(ctx context.Context, data T) error {
		start := time.Now()
		err := handler(ctx, data)
		duration := time.Since(start).Seconds()

		status := "success"
		if err != nil {
			status = "error"
		}

		MessageLatency.WithLabelValues(cfg.Mode, cfg.Stage, status).Observe(duration)
		if cfg.Service != "" {
			MessagesProcessed.WithLabelValues(cfg.Service, cfg.EventType, status).Inc()
		}
		return err
	}
}
