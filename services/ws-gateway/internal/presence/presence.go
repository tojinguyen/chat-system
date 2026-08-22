package presence

import (
	"context"
	"time"
)

// Reporter reports connection presence to Redis
type Reporter interface {
	SetOnline(ctx context.Context, userID, deviceID, gatewayNode string, ttl time.Duration) error
	SetOffline(ctx context.Context, userID, deviceID string) error
	Heartbeat(ctx context.Context, userID, deviceID string, ttl time.Duration) error
}

type redisPresenceReporter struct {
	// TODO: redis client reference
}

func NewRedisReporter(redisAddr, password string, db int) Reporter {
	return &redisPresenceReporter{}
}

func (r *redisPresenceReporter) SetOnline(ctx context.Context, userID, deviceID, gatewayNode string, ttl time.Duration) error {
	// Key format: presence:{user_id}:{device_id} -> value: gatewayNode
	return nil
}

func (r *redisPresenceReporter) SetOffline(ctx context.Context, userID, deviceID string) error {
	return nil
}

func (r *redisPresenceReporter) Heartbeat(ctx context.Context, userID, deviceID string, ttl time.Duration) error {
	return nil
}
