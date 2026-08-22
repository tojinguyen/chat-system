package consumer

import (
	"context"
	"notification-service/internal/domain"
)

type NotificationHandler func(ctx context.Context, event *domain.NotificationEvent) error

type NotificationConsumer interface {
	Start(ctx context.Context, handler NotificationHandler) error
	Close() error
}

type natsNotificationConsumer struct {
	// TODO: nats.Conn / QueueSubscriber reference
}

func NewNotificationConsumer(natsURL string, subject string, queueGroup string) (NotificationConsumer, error) {
	return &natsNotificationConsumer{}, nil
}

func (c *natsNotificationConsumer) Start(ctx context.Context, handler NotificationHandler) error {
	// TODO: natsConn.QueueSubscribe(subject, queueGroup, func(msg *nats.Msg) { ... })
	return nil
}

func (c *natsNotificationConsumer) Close() error {
	return nil
}
