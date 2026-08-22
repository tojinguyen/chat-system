package domain

import "time"

type MessageStatus string

const (
	StatusSent      MessageStatus = "SENT"
	StatusDelivered MessageStatus = "DELIVERED"
	StatusRead      MessageStatus = "READ"
)

type ConversationType string

const (
	ConversationDirect ConversationType = "DIRECT"
	ConversationGroup  ConversationType = "GROUP"
	ConversationClan   ConversationType = "CLAN"
	ConversationRegion ConversationType = "REGION"
	ConversationGlobal ConversationType = "GLOBAL"
)

// Message represents the persisted message entity (ScyllaDB / Cassandra)
// Partition key: ConversationID, Clustering key: MessageID (UUIDv7)
type Message struct {
	ConversationID string        `json:"conversation_id"`
	MessageID      string        `json:"message_id"`
	SenderID       string        `json:"sender_id"`
	Content        string        `json:"content"`
	Status         MessageStatus `json:"status"`
	CreatedAt      time.Time     `json:"created_at"`
}

// InboundMessageEvent is the event consumed from chat.inbound
type InboundMessageEvent struct {
	ClientMsgID    string    `json:"client_msg_id"`
	ConversationID string    `json:"conversation_id"`
	SenderID       string    `json:"sender_id"`
	DeviceID       string    `json:"device_id"`
	GatewayNode    string    `json:"gateway_node"`
	Content        string    `json:"content"`
	SentAt         time.Time `json:"sent_at"`
}

// OutboundMessageEvent is published to specific gateway node topic
type OutboundMessageEvent struct {
	MessageID      string    `json:"message_id"`
	ConversationID string    `json:"conversation_id"`
	SenderID       string    `json:"sender_id"`
	RecipientID    string    `json:"recipient_id"`
	DeviceID       string    `json:"device_id"`
	Content        string    `json:"content"`
	Timestamp      time.Time `json:"timestamp"`
}

// NotificationEvent is published to notification topic if recipient is offline
type NotificationEvent struct {
	RecipientID    string    `json:"recipient_id"`
	SenderID       string    `json:"sender_id"`
	ConversationID string    `json:"conversation_id"`
	MessageID      string    `json:"message_id"`
	PreviewContent string    `json:"preview_content"`
	CreatedAt      time.Time `json:"created_at"`
}
