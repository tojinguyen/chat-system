package provider

import (
	"context"
	"notification-service/internal/domain"
)

type PushProvider interface {
	SendPush(ctx context.Context, payload *domain.PushPayload) error
}
