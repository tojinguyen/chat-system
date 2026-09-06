package dispatcher

import (
	"chat-system/pkg/contracts"
	"context"
	"fmt"
	"log"

	natsclient "chat-system/pkg/nats"

	"github.com/nats-io/nats.go"
)

type EventDispatcher interface {
	SendAckToSender(ctx context.Context, gatewayNode string, event contracts.OutboundBrokerEvent) error
}

type natsEventDispatcher struct {
	publisher *natsclient.Publisher[contracts.OutboundBrokerEvent]
}

// NewEventDispatcher khởi tạo Dispatcher sử dụng NATS Publisher
func NewEventDispatcher(nc *nats.Conn) EventDispatcher {
	return &natsEventDispatcher{
		publisher: natsclient.NewPublisher[contracts.OutboundBrokerEvent](nc, ""),
	}
}

// SendAckToSender gửi sự kiện ACK về đúng Node Gateway mà người gửi đang kết nối
func (d *natsEventDispatcher) SendAckToSender(ctx context.Context, gatewayNode string, event contracts.OutboundBrokerEvent) error {
	if gatewayNode == "" {
		return fmt.Errorf("gatewayNode is empty, cannot route sender ack")
	}
	// Topic định tuyến riêng cho từng node gateway: "chat.gateway.{node_id}"
	subject := contracts.GatewayNodeSubject(gatewayNode)
	if err := d.publisher.PublishToSubject(ctx, subject, event); err != nil {
		return fmt.Errorf("failed to publish sender ack to subject '%s': %w", subject, err)
	}
	log.Printf("[Dispatcher] Sender ACK published: msg_id=%s, client_msg_id=%s -> subject=%s",
		event.MessageID, event.ClientMsgID, subject)
	return nil
}
