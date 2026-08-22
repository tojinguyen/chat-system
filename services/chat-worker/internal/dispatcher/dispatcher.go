package dispatcher

import (
	"context"
	"chat-worker/internal/domain"
)

// Dispatcher routes messages either to specific WebSocket Gateway Subject or to Notification Subject
type Dispatcher interface {
	DispatchToGateway(ctx context.Context, gatewayNode string, event *domain.OutboundMessageEvent) error
	DispatchToNotification(ctx context.Context, event *domain.NotificationEvent) error
}

type natsDispatcher struct {
	// TODO: nats.Conn publisher reference
}

func NewNATSDispatcher(natsURL string) Dispatcher {
	return &natsDispatcher{}
}

func (d *natsDispatcher) DispatchToGateway(ctx context.Context, gatewayNode string, event *domain.OutboundMessageEvent) error {
	// Publish to NATS Subject: chat.gateway.{gatewayNode}
	return nil
}

func (d *natsDispatcher) DispatchToNotification(ctx context.Context, event *domain.NotificationEvent) error {
	// Publish to NATS Subject: chat.notifications
	return nil
}
