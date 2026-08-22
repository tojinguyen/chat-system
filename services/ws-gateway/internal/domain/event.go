package domain

import "time"

// EventType represents the type of WebSocket event
type EventType string

const (
	EventSendMessage   EventType = "SEND_MESSAGE"
	EventReceiveMessage EventType = "RECEIVE_MESSAGE"
	EventAckSent       EventType = "ACK_SENT"
	EventAckDelivered  EventType = "ACK_DELIVERED"
	EventAckRead       EventType = "ACK_READ"
	EventHeartbeat     EventType = "HEARTBEAT"
)

// WSMessage represents the payload exchanged with client over WebSocket
type WSMessage struct {
	Type           EventType   `json:"type"`
	ClientMsgID    string      `json:"client_msg_id,omitempty"`
	MessageID      string      `json:"message_id,omitempty"`
	ConversationID string      `json:"conversation_id,omitempty"`
	SenderID       string      `json:"sender_id,omitempty"`
	Content        string      `json:"content,omitempty"`
	Timestamp      int64       `json:"timestamp,omitempty"`
	Payload        interface{} `json:"payload,omitempty"`
}

// InboundBrokerEvent represents the message forwarded from Gateway to Broker
type InboundBrokerEvent struct {
	ClientMsgID    string    `json:"client_msg_id"`
	ConversationID string    `json:"conversation_id"`
	SenderID       string    `json:"sender_id"`
	DeviceID       string    `json:"device_id"`
	GatewayNode    string    `json:"gateway_node"`
	Content        string    `json:"content"`
	SentAt         time.Time `json:"sent_at"`
}

// OutboundBrokerEvent represents the message received from Broker to be pushed to client
type OutboundBrokerEvent struct {
	MessageID      string    `json:"message_id"`
	ConversationID string    `json:"conversation_id"`
	SenderID       string    `json:"sender_id"`
	RecipientID    string    `json:"recipient_id"`
	DeviceID       string    `json:"device_id"`
	Content        string    `json:"content"`
	Timestamp      time.Time `json:"timestamp"`
}
