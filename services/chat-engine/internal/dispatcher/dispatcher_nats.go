package dispatcher

import (
	"context"
	"fmt"
	"log"

	"chat-system/pkg/contracts"
	natsclient "chat-system/pkg/nats"

	"github.com/nats-io/nats.go"
)

type natsEventDispatcher struct {
	outboundPublisher *natsclient.InstrumentedPublisher[contracts.OutboundBrokerEvent]
}

// NewNATSEventDispatcher khởi tạo Dispatcher sử dụng NATS Pub/Sub với Telemetry Decorator
func NewNATSEventDispatcher(nc *nats.Conn) EventDispatcher {
	rawPublisher := natsclient.NewPublisher[contracts.OutboundBrokerEvent](nc, "")
	return &natsEventDispatcher{
		outboundPublisher: natsclient.NewInstrumentedPublisher(rawPublisher, natsclient.PublisherConfig{
			ServiceName: "chat-engine",
			Stage:       "outbound_dispatch",
			EventType:   "outbound",
		}),
	}
}

// SendAckToSender gửi sự kiện ACK về đúng Topic của Node Gateway mà người gửi kết nối: "chat.gateway.{node_id}"
func (d *natsEventDispatcher) SendAckToSender(ctx context.Context, gatewayNode string, event contracts.OutboundBrokerEvent) error {
	if gatewayNode == "" {
		return fmt.Errorf("gatewayNode is empty, cannot route sender ack")
	}
	subject := contracts.GatewayNodeSubject(gatewayNode)
	if err := d.outboundPublisher.PublishToSubject(ctx, subject, event); err != nil {
		return fmt.Errorf("failed to publish sender ack to subject '%s': %w", subject, err)
	}
	log.Printf("[Dispatcher:NATS] Sender ACK published: msg_id=%s, client_msg_id=%s -> subject=%s",
		event.MessageID, event.ClientMsgID, subject)
	return nil
}

// DispatchToGateway chuyển tiếp tin nhắn tới đúng Topic của Node Gateway mà người nhận kết nối
func (d *natsEventDispatcher) DispatchToGateway(ctx context.Context, gatewayNode string, event contracts.OutboundBrokerEvent) error {
	if gatewayNode == "" {
		return fmt.Errorf("gatewayNode is empty, cannot route message to gateway")
	}

	subject := contracts.GatewayNodeSubject(gatewayNode)
	if err := d.outboundPublisher.PublishToNode(ctx, gatewayNode, subject, event); err != nil {
		return fmt.Errorf("failed to publish outbound message to subject '%s': %w", subject, err)
	}

	log.Printf("[Dispatcher:NATS] Message dispatched to receiver gateway: msg_id=%s, receiver=%s -> subject=%s",
		event.MessageID, event.ReceiverID, subject)
	return nil
}

func (d *natsEventDispatcher) Close() error {
	return nil
}
