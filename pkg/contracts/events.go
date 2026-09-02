package contracts

import (
	"encoding/json"
	"time"
)

// BrokerMessageType represents standardized event types published across the message broker
type BrokerMessageType string

const (
	BrokerEventMessageSubmitted BrokerMessageType = "MESSAGE_SUBMITTED"
	BrokerEventMessageDelivered BrokerMessageType = "MESSAGE_DELIVERED"
	BrokerEventNotification     BrokerMessageType = "NOTIFICATION_DISPATCH"
)

// InboundBrokerEvent represents the message forwarded from WebSocket Gateway to Chat Engine via NATS
type InboundBrokerEvent struct {
	Type        BrokerMessageType `json:"type"`
	ClientMsgID string            `json:"client_msg_id"`
	SenderID    string            `json:"sender_id"`
	DeviceID    string            `json:"device_id"`
	GatewayNode string            `json:"gateway_node"`
	Payload     json.RawMessage   `json:"payload"`
	SentAt      time.Time         `json:"sent_at"`
}

// SendMessagePayload is the unmarshaled payload of a MESSAGE_SUBMITTED event
type SendMessagePayload struct {
	ConversationID string `json:"conversation_id"`
	Content        string `json:"content"`
}

// OutboundBrokerEvent represents the message routed from Chat Engine to a specific Gateway Node
type OutboundBrokerEvent struct {
	MessageID      string          `json:"message_id"`
	ClientMsgID    string          `json:"client_msg_id,omitempty"`
	ConversationID string          `json:"conversation_id"`
	SenderID       string          `json:"sender_id"`
	ReceiverID     string          `json:"receiver_id,omitempty"`
	Content        string          `json:"content"`
	Type           string          `json:"type"`
	Timestamp      int64           `json:"timestamp"`
	Payload        json.RawMessage `json:"payload,omitempty"`
}

// NotificationBrokerEvent represents payload routed to Notification Service when receiver is offline
type NotificationBrokerEvent struct {
	RecipientID    string `json:"recipient_id"`
	SenderID       string `json:"sender_id"`
	ConversationID string `json:"conversation_id"`
	ContentSnippet string `json:"content_snippet"`
	Timestamp      int64  `json:"timestamp"`
}
