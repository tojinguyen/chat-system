package domain

import (
	"encoding/json"
	"time"
)

// WSEventType represents the type of WebSocket event
type WSEventType string

const (
	WSEventSendMessage    WSEventType = "SEND_MESSAGE"
	WSEventReceiveMessage WSEventType = "RECEIVE_MESSAGE"
	WSEventAckSent        WSEventType = "ACK_SENT"
	WSEventAckDelivered   WSEventType = "ACK_DELIVERED"
	WSEventAckRead        WSEventType = "ACK_READ"
	WSEventHeartbeat      WSEventType = "HEARTBEAT"
)

// Broker message type
type BrokerMessageType string

const (
	BrokerEventMessageSubmitted BrokerMessageType = "MESSAGE_SUBMITTED"
)

// WSMessage represents the payload exchanged with client over WebSocket
type WSMessage struct {
	Type        WSEventType     `json:"type"`
	ClientMsgID string          `json:"client_msg_id,omitempty"`
	Timestamp   int64           `json:"timestamp,omitempty"`
	Payload     json.RawMessage `json:"payload,omitempty"`
}

type SendMessagePayload struct {
	ConversationID string `json:"conversation_id"`
	Content        string `json:"content"`
}

// InboundBrokerEvent represents the message forwarded from Gateway to Broker
type InboundBrokerEvent struct {
	Type        BrokerMessageType `json:"type"`
	ClientMsgID string            `json:"client_msg_id"`
	SenderID    string            `json:"sender_id"`
	DeviceID    string            `json:"device_id"`
	GatewayNode string            `json:"gateway_node"`
	Payload     string            `json:"payload"`
	SentAt      time.Time         `json:"sent_at"`
}

// OutboundBrokerEvent represents the message received from Broker to be pushed to client
type OutboundBrokerEvent struct {
}
