package repository

import (
	"context"
	"fmt"
	"net"
	"strconv"
	"strings"
	"time"

	"chat-worker/internal/config"
	"chat-worker/internal/domain"

	"github.com/gocql/gocql"
)

// MessageRepository defines the data access contract for messages
type MessageRepository interface {
	SaveMessage(ctx context.Context, msg *domain.Message) error
	GetMessagesByConversation(ctx context.Context, conversationID string, limit int) ([]*domain.Message, error)
}

// CassandraMessageRepository implements MessageRepository using Apache Cassandra
type CassandraMessageRepository struct {
	session *gocql.Session
	table   string
}

// NewCassandraSession creates and initializes a Cassandra session based on configuration
func NewCassandraSession(cfg *config.DatabaseConfig) (*gocql.Session, error) {
	if len(cfg.Hosts) == 0 {
		return nil, fmt.Errorf("cassandra hosts list is empty")
	}

	var hostList []string
	port := 9042

	for _, h := range cfg.Hosts {
		if strings.Contains(h, ":") {
			host, portStr, err := net.SplitHostPort(h)
			if err == nil {
				hostList = append(hostList, host)
				if parsedPort, err := strconv.Atoi(portStr); err == nil && parsedPort > 0 {
					port = parsedPort
				}
				continue
			}
		}
		hostList = append(hostList, h)
	}

	cluster := gocql.NewCluster(hostList...)
	cluster.Port = port
	cluster.Keyspace = cfg.Keyspace
	cluster.Consistency = gocql.LocalQuorum
	cluster.Timeout = 5 * time.Second
	cluster.ConnectTimeout = 10 * time.Second
	cluster.RetryPolicy = &gocql.SimpleRetryPolicy{NumRetries: 3}

	session, err := cluster.CreateSession()
	if err != nil {
		return nil, fmt.Errorf("failed to connect to cassandra cluster: %w", err)
	}

	return session, nil
}

// NewCassandraMessageRepository creates a new instance of CassandraMessageRepository
func NewCassandraMessageRepository(session *gocql.Session, table string) *CassandraMessageRepository {
	if table == "" {
		table = "messages"
	}
	return &CassandraMessageRepository{
		session: session,
		table:   table,
	}
}

// SaveMessage inserts a new message into Cassandra
func (r *CassandraMessageRepository) SaveMessage(ctx context.Context, msg *domain.Message) error {
	query := fmt.Sprintf(`
		INSERT INTO %s (conversation_id, id, sender_id, content, created_at, updated_at)
		VALUES (?, ?, ?, ?, ?, ?)
	`, r.table)

	return r.session.Query(query,
		msg.ConversationID,
		msg.ID,
		msg.SenderID,
		msg.Content,
		msg.CreatedAt,
		msg.UpdatedAt,
	).WithContext(ctx).Exec()
}

// GetMessagesByConversation retrieves recent messages for a conversation
func (r *CassandraMessageRepository) GetMessagesByConversation(ctx context.Context, conversationID string, limit int) ([]*domain.Message, error) {
	if limit <= 0 {
		limit = 50
	}

	query := fmt.Sprintf(`
		SELECT conversation_id, id, sender_id, content, created_at, updated_at
		FROM %s
		WHERE conversation_id = ?
		LIMIT ?
	`, r.table)

	iter := r.session.Query(query, conversationID, limit).WithContext(ctx).Iter()
	defer iter.Close()

	var messages []*domain.Message
	var convID, id, senderID, content string
	var createdAt, updatedAt time.Time

	for iter.Scan(&convID, &id, &senderID, &content, &createdAt, &updatedAt) {
		messages = append(messages, &domain.Message{
			ConversationID: convID,
			ID:             id,
			SenderID:       senderID,
			Content:        content,
			CreatedAt:      createdAt,
			UpdatedAt:      updatedAt,
		})
	}

	if err := iter.Close(); err != nil {
		return nil, fmt.Errorf("failed to query messages: %w", err)
	}

	return messages, nil
}
