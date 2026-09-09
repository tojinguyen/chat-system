package domain

import "time"

// Message represents the core message entity persisted in Cassandra
type Message struct {
	ConversationID string    `json:"conversation_id"`
	ID             string    `json:"id"`
	SenderID       string    `json:"sender_id"`
	Content        string    `json:"content"`
	CreatedAt      time.Time `json:"created_at"`
	UpdatedAt      time.Time `json:"updated_at"`
}
