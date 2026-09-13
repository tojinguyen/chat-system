package dispatcher

import (
	"context"
	"fmt"
	"log"

	"chat-system/pkg/contracts"
	grpcclient "chat-system/pkg/grpcclient"
	natsclient "chat-system/pkg/nats"
	pb "chat-system/pkg/proto"
	"chat-worker/internal/config"

	"github.com/nats-io/nats.go"
)

type grpcEventDispatcher struct {
	ackPublisher  *natsclient.Publisher[contracts.OutboundBrokerEvent]
	clientManager *grpcclient.GatewayClientManager
}

// NewGRPCEventDispatcher khởi tạo Dispatcher sử dụng GatewayClientManager từ pkg/grpcclient
func NewGRPCEventDispatcher(cfg *config.DeliveryConfig, nc *nats.Conn) EventDispatcher {
	return &grpcEventDispatcher{
		ackPublisher:  natsclient.NewPublisher[contracts.OutboundBrokerEvent](nc, ""),
		clientManager: grpcclient.NewGatewayClientManager(cfg.GRPCPort, cfg.GRPCServiceSuffix),
	}
}

// SendAckToSender gửi sự kiện ACK về Node Gateway của người gửi qua Broker (NATS Topic)
func (d *grpcEventDispatcher) SendAckToSender(ctx context.Context, gatewayNode string, event contracts.OutboundBrokerEvent) error {
	if gatewayNode == "" {
		return fmt.Errorf("gatewayNode is empty, cannot route sender ack")
	}

	subject := contracts.GatewayNodeSubject(gatewayNode)
	if err := d.ackPublisher.PublishToSubject(ctx, subject, event); err != nil {
		return fmt.Errorf("failed to publish sender ack to subject '%s': %w", subject, err)
	}

	log.Printf("[Dispatcher:gRPC] Sender ACK published via Broker: msg_id=%s, client_msg_id=%s -> subject=%s",
		event.MessageID, event.ClientMsgID, subject)
	return nil
}

// DispatchToGateway gửi tin nhắn trực tiếp qua gRPC tới WS Gateway Pod mà người nhận kết nối
func (d *grpcEventDispatcher) DispatchToGateway(ctx context.Context, gatewayNode string, event contracts.OutboundBrokerEvent) error {
	if gatewayNode == "" {
		return fmt.Errorf("gatewayNode is empty, cannot route message to gateway via gRPC")
	}

	client, err := d.clientManager.GetClient(gatewayNode)
	if err != nil {
		return fmt.Errorf("failed to get gRPC client for node %s: %w", gatewayNode, err)
	}

	req := &pb.PushMessageRequest{
		MessageId:      event.MessageID,
		ClientMsgId:    event.ClientMsgID,
		ConversationId: event.ConversationID,
		SenderId:       event.SenderID,
		ReceiverId:     event.ReceiverID,
		Content:        event.Content,
		Type:           string(event.Type),
		Timestamp:      event.Timestamp,
		Payload:        event.Payload,
	}

	resp, err := client.PushMessageToUser(ctx, req)
	if err != nil {
		return fmt.Errorf("gRPC PushMessageToUser failed for receiver %s: %w", event.ReceiverID, err)
	}
	if !resp.Success {
		return fmt.Errorf("gRPC PushMessageToUser rejected: %s", resp.ErrorMessage)
	}

	log.Printf("[Dispatcher:gRPC] Message dispatched successfully: msg_id=%s, receiver=%s -> gateway=%s",
		event.MessageID, event.ReceiverID, gatewayNode)
	return nil
}

func (d *grpcEventDispatcher) Close() error {
	return d.clientManager.Close()
}
