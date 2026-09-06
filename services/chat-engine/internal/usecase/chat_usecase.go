package usecase

import (
	"chat-system/pkg/contracts"
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
}

func NewChatUsecase(messageRepo repository.MessageRepository) ChatUsecase {
	return &chatUsecase{
		messageRepo: messageRepo,
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
		ID:             event.ClientMsgID, // Dùng ClientMsgID (UUIDv7) làm Message ID
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
	return nil
}
