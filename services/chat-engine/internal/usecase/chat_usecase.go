package usecase

import (
	"chat-system/pkg/contracts"
	"chat-worker/internal/dispatcher"
	"chat-worker/internal/domain"
	"chat-worker/internal/repository"
	"context"
	"encoding/json"
	"fmt"
	"log"
	"time"
)

type ChatUsecase interface {
	ProcessInboundMessage(ctx context.Context, event contracts.InboundBrokerEvent) error
}
type chatUsecase struct {
	messageRepo repository.MessageRepository
	dispatcher  dispatcher.EventDispatcher
}

func NewChatUsecase(messageRepo repository.MessageRepository, dispatcher dispatcher.EventDispatcher) ChatUsecase {
	return &chatUsecase{
		messageRepo: messageRepo,
		dispatcher:  dispatcher,
	}
}

// ProcessInboundMessage xử lý sự kiện tin nhắn gửi đến từ WebSocket Gateway
func (u *chatUsecase) ProcessInboundMessage(ctx context.Context, event contracts.InboundBrokerEvent) error {
	// 1. Unmarshal payload tin nhắn từ InboundBrokerEvent
	var payload contracts.SendMessagePayload
	if err := json.Unmarshal(event.Payload, &payload); err != nil {
		return fmt.Errorf("failed to unmarshal message payload: %w", err)
	}
	if payload.ConversationID == "" || payload.Content == "" {
		return fmt.Errorf("invalid payload: conversation_id or content is empty")
	}
	// 2. Chuyển đổi sang Domain Entity
	now := time.Now().UTC()
	msg := &domain.Message{
		ConversationID: payload.ConversationID,
		ID:             event.ClientMsgID,
		SenderID:       event.SenderID,
		Content:        payload.Content,
		CreatedAt:      now,
		UpdatedAt:      now,
	}
	// 3. Giới hạn timeout khi ghi vào Cassandra (Defensive Concurrency)
	dbCtx, cancel := context.WithTimeout(ctx, 3*time.Second)
	defer cancel()
	if err := u.messageRepo.SaveMessage(dbCtx, msg); err != nil {
		return fmt.Errorf("persistence failed: %w", err)
	}
	log.Printf("[Usecase] Message persisted successfully: msg_id=%s, conversation_id=%s, sender=%s",
		msg.ID, msg.ConversationID, msg.SenderID)

	// 4. Gửi Sender ACK về cho người gửi (Client A)
	ackEvent := contracts.OutboundBrokerEvent{
		MessageID:      msg.ID,
		ClientMsgID:    event.ClientMsgID,
		ConversationID: msg.ConversationID,
		SenderID:       msg.SenderID,
		Type:           contracts.BrokerEventMessageSubmitted,
		Timestamp:      now.UnixMilli(),
	}
	if err := u.dispatcher.SendAckToSender(ctx, event.GatewayNode, ackEvent); err != nil {
		// Log cảnh báo nhưng không làm fail cả luồng vì tin nhắn đã được lưu DB an toàn
		log.Printf("[Usecase] WARN: Failed to send Sender ACK: %v", err)
	}
	return nil
}
