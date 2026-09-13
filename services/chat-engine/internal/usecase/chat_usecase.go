package usecase

import (
	"chat-system/pkg/contracts"
	"chat-worker/internal/dispatcher"
	"chat-worker/internal/domain"
	"chat-worker/internal/presence"
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
	messageRepo     repository.MessageRepository
	idempotencyRepo repository.IdempotencyRepository
	presenceReader  presence.PresenceReader
	dispatcher      dispatcher.EventDispatcher
	dbTimeout       time.Duration
}

func NewChatUsecase(
	messageRepo repository.MessageRepository,
	idempotencyRepo repository.IdempotencyRepository,
	presenceReader presence.PresenceReader,
	dispatcher dispatcher.EventDispatcher,
	dbTimeout time.Duration,
) ChatUsecase {
	if dbTimeout <= 0 {
		dbTimeout = 5 * time.Second
	}
	return &chatUsecase{
		messageRepo:     messageRepo,
		idempotencyRepo: idempotencyRepo,
		presenceReader:  presenceReader,
		dispatcher:      dispatcher,
		dbTimeout:       dbTimeout,
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

	now := time.Now().UTC()

	// 2. IDEMPOTENCY CHECK (Atomic SetNX trên Redis)
	isNew, err := u.idempotencyRepo.AcquireLock(ctx, event.ClientMsgID)
	if err != nil {
		// Nếu Redis lỗi (Redis down), log cảnh báo nhưng cho phép luồng tiếp tục (Graceful Degradation)
		log.Printf("[Usecase] WARN: Redis idempotency check failed: %v", err)
	} else if !isNew {
		// Tin nhắn TRÙNG LẶP: Bỏ qua ghi DB, gửi lại Sender ACK cho Client
		log.Printf("[Usecase] Duplicate message detected: client_msg_id=%s. Resending Sender ACK only.", event.ClientMsgID)

		ackEvent := contracts.OutboundBrokerEvent{
			MessageID:      event.ClientMsgID,
			ClientMsgID:    event.ClientMsgID,
			ConversationID: payload.ConversationID,
			SenderID:       event.SenderID,
			Type:           contracts.BrokerEventMessageSubmitted,
			Timestamp:      now.UnixMilli(),
		}
		_ = u.dispatcher.SendAckToSender(ctx, event.GatewayNode, ackEvent)
		return nil
	}

	// 3. Chuyển đổi sang Domain Entity
	msg := &domain.Message{
		ConversationID: payload.ConversationID,
		ID:             event.ClientMsgID,
		SenderID:       event.SenderID,
		Content:        payload.Content,
		CreatedAt:      now,
		UpdatedAt:      now,
	}

	// 4. Giới hạn timeout khi ghi vào Cassandra (Defensive Concurrency - Configurable)
	dbCtx, cancel := context.WithTimeout(ctx, u.dbTimeout)
	defer cancel()
	if err := u.messageRepo.SaveMessage(dbCtx, msg); err != nil {
		// Rollback/Release lock trên Redis nếu ghi DB thất bại để cho phép Client retry
		if relErr := u.idempotencyRepo.ReleaseLock(ctx, event.ClientMsgID); relErr != nil {
			log.Printf("[Usecase] ERROR: Failed to release idempotency lock for client_msg_id=%s: %v", event.ClientMsgID, relErr)
		}
		return fmt.Errorf("persistence failed: %w", err)
	}
	log.Printf("[Usecase] Message persisted successfully: msg_id=%s, conversation_id=%s, sender=%s",
		msg.ID, msg.ConversationID, msg.SenderID)

	// 5. Gửi Sender ACK về cho người gửi (Client A)
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

	// 6. OUTBOUND DELIVERY: Tra cứu Presence và chuyển tiếp tin nhắn tới người nhận
	receiverID := payload.ReceiverID
	if receiverID != "" && receiverID != event.SenderID {
		routes, err := u.presenceReader.GetUserRoutes(ctx, receiverID)
		if err != nil {
			log.Printf("[Usecase] ERROR: Failed to query presence for receiver %s: %v", receiverID, err)
		}

		if len(routes) > 0 {
			// Gom nhóm theo unique gatewayNode để tránh gửi trùng lặp nếu 1 node có nhiều device
			dispatchedNodes := make(map[string]bool)
			for _, gatewayNode := range routes {
				if dispatchedNodes[gatewayNode] {
					continue
				}
				dispatchedNodes[gatewayNode] = true

				outboundMsg := contracts.OutboundBrokerEvent{
					MessageID:      msg.ID,
					ClientMsgID:    event.ClientMsgID,
					ConversationID: msg.ConversationID,
					SenderID:       msg.SenderID,
					ReceiverID:     receiverID,
					Content:        msg.Content,
					Type:           contracts.BrokerEventMessageSubmitted,
					Timestamp:      now.UnixMilli(),
				}

				if err := u.dispatcher.DispatchToGateway(ctx, gatewayNode, outboundMsg); err != nil {
					log.Printf("[Usecase] ERROR: Failed to dispatch message to receiver gateway %s: %v", gatewayNode, err)
				}
			}
		} else {
			log.Printf("[Usecase] Receiver %s is offline. Skipping realtime push.", receiverID)
		}
	}

	return nil
}
