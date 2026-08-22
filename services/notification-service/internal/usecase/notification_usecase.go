package usecase

import (
	"context"
	"notification-service/internal/domain"
	"notification-service/internal/provider"
)

type NotificationUsecase struct {
	fcmProvider  provider.PushProvider
	apnsProvider provider.PushProvider
}

func NewNotificationUsecase(fcm, apns provider.PushProvider) *NotificationUsecase {
	return &NotificationUsecase{
		fcmProvider:  fcm,
		apnsProvider: apns,
	}
}

func (u *NotificationUsecase) HandleNotification(ctx context.Context, event *domain.NotificationEvent) error {
	return nil
}
