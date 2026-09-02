package domain

import (
	"encoding/json"

	"chat-system/pkg/contracts"
)

// WSEventType represents the type of WebSocket event
type WSEventType string

const (
	WSEventSendMessage  WSEventType = "SEND_MESSAGE"
	WSEventFailedToSend WSEventType = "FAILED_TO_SEND"
	WSEventHeartbeat    WSEventType = "HEARTBEAT"
)

func (t WSEventType) ToBrokerMessageType() (contracts.BrokerMessageType, bool) {
	switch t {
	case WSEventSendMessage:
		return contracts.BrokerEventMessageSubmitted, true
	default:
		return "", false
	}
}

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

type FailedToSendPayload struct {
	Error string `json:"error"`
}
