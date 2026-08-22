package repository

import (
	"context"
	"chat-worker/internal/domain"
)

// MessageRepository manages persistence to ScyllaDB / Cassandra
type MessageRepository interface {
	SaveMessage(ctx context.Context, msg *domain.Message) error
	UpdateStatus(ctx context.Context, conversationID, messageID string, status domain.MessageStatus) error
}

type scyllaMessageRepo struct {
	// TODO: scylla / gocql session
}

func NewScyllaMessageRepository(hosts []string, keyspace string) MessageRepository {
	return &scyllaMessageRepo{}
}

func (r *scyllaMessageRepo) SaveMessage(ctx context.Context, msg *domain.Message) error {
	// TODO: INSERT INTO messages (conversation_id, message_id, sender_id, content, status, created_at) VALUES (...)
	return nil
}

func (r *scyllaMessageRepo) UpdateStatus(ctx context.Context, conversationID, messageID string, status domain.MessageStatus) error {
	// TODO: UPDATE messages SET status = ? WHERE conversation_id = ? AND message_id = ?
	return nil
}
