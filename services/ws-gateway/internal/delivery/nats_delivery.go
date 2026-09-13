package delivery

import (
	"context"
	"fmt"
	"log"

	"chat-system/pkg/contracts"
	natsclient "chat-system/pkg/nats"
	"ws-gateway/internal/connection"
	"ws-gateway/internal/payload"

	"github.com/nats-io/nats.go"
)

type NATSListener struct {
	nodeID     string
	subscriber *natsclient.Subscriber[contracts.OutboundBrokerEvent]
	hub        *connection.Hub
}

func NewNATSListener(nc *nats.Conn, nodeID string, hub *connection.Hub) *NATSListener {
	subject := contracts.GatewayNodeSubject(nodeID)
	return &NATSListener{
		nodeID:     nodeID,
		subscriber: natsclient.NewSubscriber[contracts.OutboundBrokerEvent](nc, subject),
		hub:        hub,
	}
}

func (n *NATSListener) Start(ctx context.Context) error {
	err := n.subscriber.Start(ctx, func(ctx context.Context, event contracts.OutboundBrokerEvent) error {
		deliveryPayload := payload.MessageDeliveryPayload{
			MessageID:      event.MessageID,
			ClientMsgID:    event.ClientMsgID,
			ConversationID: event.ConversationID,
			SenderID:       event.SenderID,
			ReceiverID:     event.ReceiverID,
			Content:        event.Content,
			Type:           event.Type,
			Timestamp:      event.Timestamp,
		}

		wsMsg, err := deliveryPayload.NewWSMessageFromDelivery()
		if err != nil {
			return err
		}

		targetUser := event.ReceiverID
		if targetUser == "" {
			targetUser = event.SenderID
		}
		n.hub.SendToUser(targetUser, wsMsg)
		return nil
	})

	if err != nil {
		return fmt.Errorf("failed to start NATS listener: %w", err)
	}

	log.Printf("[Delivery] Subscribed to NATS Subject: %s", contracts.GatewayNodeSubject(n.nodeID))
	return nil
}

func (n *NATSListener) Stop(ctx context.Context) error {
	log.Println("[Delivery] NATS listener stopped")
	return nil
}
