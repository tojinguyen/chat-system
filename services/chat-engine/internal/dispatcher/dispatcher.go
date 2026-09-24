package dispatcher

import (
	"context"
	"fmt"
	"log"

	"chat-system/pkg/contracts"
	"chat-worker/internal/config"

	"github.com/nats-io/nats.go"
)

// EventDispatcher định nghĩa giao diện chung cho việc định tuyến outbound event (gRPC hoặc NATS Broker)
type EventDispatcher interface {
	// SendAckToSender gửi sự kiện ACK về đúng Node Gateway mà người gửi đang kết nối
	SendAckToSender(ctx context.Context, gatewayNode string, event contracts.OutboundBrokerEvent) error

	// DispatchToGateway chuyển tiếp tin nhắn tới đúng Node Gateway mà người nhận đang kết nối
	DispatchToGateway(ctx context.Context, gatewayNode string, event contracts.OutboundBrokerEvent) error

	// Close giải phóng tài nguyên kết nối (gRPC conns, etc.)
	Close() error
}

// NewEventDispatcher khởi tạo dispatcher tương ứng với DeliveryMode đã cấu hình
func NewEventDispatcher(cfg *config.DeliveryConfig, nc *nats.Conn) (EventDispatcher, error) {
	switch cfg.Mode {
	case "grpc":
		log.Printf("[Dispatcher] Initializing gRPC Event Dispatcher (default port: %d, suffix: '%s')",
			cfg.GRPCPort, cfg.GRPCServiceSuffix)
		return NewGRPCEventDispatcher(cfg, nc), nil
	case "broker":
		log.Println("[Dispatcher] Initializing NATS Broker Event Dispatcher")
		return NewNATSEventDispatcher(nc), nil
	default:
		return nil, fmt.Errorf("unsupported delivery mode '%s', must be 'grpc' or 'broker'", cfg.Mode)
	}
}
