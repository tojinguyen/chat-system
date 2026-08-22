package provider

import (
	"context"
	"notification-service/internal/domain"
)

type APNsProvider struct {
	// TODO: Apple Push Notification HTTP/2 client
}

func NewAPNsProvider(keyFile, keyID, teamID, topic string, production bool) (PushProvider, error) {
	return &APNsProvider{}, nil
}

func (p *APNsProvider) SendPush(ctx context.Context, payload *domain.PushPayload) error {
	// TODO: Send push notification via APNs
	return nil
}
