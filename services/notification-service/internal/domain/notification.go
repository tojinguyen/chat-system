package domain

import "time"

type DevicePlatform string

const (
	PlatformIOS     DevicePlatform = "IOS"
	PlatformAndroid DevicePlatform = "ANDROID"
	PlatformWeb     DevicePlatform = "WEB"
)

// NotificationEvent represents the message payload received from Kafka topic
type NotificationEvent struct {
	RecipientID    string    `json:"recipient_id"`
	SenderID       string    `json:"sender_id"`
	ConversationID string    `json:"conversation_id"`
	MessageID      string    `json:"message_id"`
	PreviewContent string    `json:"preview_content"`
	CreatedAt      time.Time `json:"created_at"`
}

// PushPayload represents normalized payload passed to APNs/FCM providers
type PushPayload struct {
	DeviceToken string
	Title       string
	Body        string
	Data        map[string]string
}
