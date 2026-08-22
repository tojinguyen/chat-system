package usecase

import (
	"context"
	"time"

	"chat-worker/internal/dispatcher"
	"chat-worker/internal/domain"
	"chat-worker/internal/presence"
	"chat-worker/internal/repository"
)

type ChatWorkerUsecase struct {
	repo       repository.MessageRepository
	presence   presence.PresenceChecker
	dispatcher dispatcher.Dispatcher
}

func NewChatWorkerUsecase(
	repo repository.MessageRepository,
	presence presence.PresenceChecker,
	dispatcher dispatcher.Dispatcher,
) *ChatWorkerUsecase {
	return &ChatWorkerUsecase{
		repo:       repo,
		presence:   presence,
		dispatcher: dispatcher,
	}
}

// ProcessInboundMessage executes the core message flow:
// 1. Idempotency Check (Redis)
// 2. Persist Message (ScyllaDB)
// 3. Sender ACK delivery
// 4. Recipient Presence Check (Redis) -> Dispatch to WS Gateway OR Notification Queue
func (u *ChatWorkerUsecase) ProcessInboundMessage(ctx context.Context, event *domain.InboundMessageEvent) error {
	// Step 1: Idempotency Check
	isNew, err := u.presence.CheckIdempotency(ctx, event.ClientMsgID, 24*time.Hour)
	if err != nil || !isNew {
		return err // Skip duplicated processing
	}

	// Step 2: Persist Message with UUIDv7
	msg := &domain.Message{
		ConversationID: event.ConversationID,
		MessageID:      event.ClientMsgID, // or generated UUIDv7
		SenderID:       event.SenderID,
		Content:        event.Content,
		Status:         domain.StatusSent,
		CreatedAt:      time.Now().UTC(),
	}
	if err := u.repo.SaveMessage(ctx, msg); err != nil {
		return err
	}

	// Step 3 & 4: Route to Recipients (Direct / Group fanout)
	// TODO: Look up conversation members, check presence for each recipient
	// If online -> u.dispatcher.DispatchToGateway(ctx, node, outboundEvent)
	// If offline -> u.dispatcher.DispatchToNotification(ctx, notiEvent)

	return nil
}
