package provider

import (
	"context"
	"notification-service/internal/domain"
)

type FCMProvider struct {
	// TODO: Firebase messaging client
}

func NewFCMProvider(credentialsFile string) (PushProvider, error) {
	return &FCMProvider{}, nil
}

func (p *FCMProvider) SendPush(ctx context.Context, payload *domain.PushPayload) error {
	// TODO: Send push notification via Firebase Cloud Messaging
	return nil
}
