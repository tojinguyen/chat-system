package presence

import (
	"context"
	"fmt"
	"time"
	"ws-gateway/internal/connection"

	"github.com/redis/go-redis/v9"
)

type redisPresenceService struct {
	client *redis.Client
}

func NewPresenceService(redisAddr, password string, db int) connection.PresenceService {
	return &redisPresenceService{
		client: redis.NewClient(&redis.Options{
			Addr:     redisAddr,
			Password: password,
			DB:       db,
		}),
	}
}

func (r *redisPresenceService) SetOnline(ctx context.Context, userID, deviceID, gatewayNode string, ttl time.Duration) error {
	// Key format: presence:{user_id}:{device_id} -> value: gatewayNode
	return r.client.Set(ctx, presenceKey(userID, deviceID), gatewayNode, ttl).Err()
}

func (r *redisPresenceService) SetOffline(ctx context.Context, userID, deviceID string) error {
	return r.client.Del(ctx, presenceKey(userID, deviceID)).Err()
}

func (r *redisPresenceService) Heartbeat(ctx context.Context, userID, deviceID string, ttl time.Duration) error {
	return r.client.Expire(ctx, presenceKey(userID, deviceID), ttl).Err()
}

func presenceKey(userID, deviceID string) string {
	return fmt.Sprintf("presence:%s:%s", userID, deviceID)
}
