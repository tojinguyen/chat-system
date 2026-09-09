package delivery

import "context"

type DeliveryListener interface {
	Start(ctx context.Context) error
	Stop(ctx context.Context) error
}
