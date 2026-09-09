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
	pipe := r.client.Pipeline()
	key := presenceKey(userID)

	pipe.HSet(ctx, key, deviceID, gatewayNode)
	pipe.Expire(ctx, key, ttl)

	_, err := pipe.Exec(ctx)
	return err
}

func (r *redisPresenceService) SetOffline(ctx context.Context, userID, deviceID string) error {
	return r.client.HDel(ctx, presenceKey(userID), deviceID).Err()
}

func (r *redisPresenceService) Heartbeat(ctx context.Context, userID, deviceID, gatewayNode string, ttl time.Duration) error {
	return r.SetOnline(ctx, userID, deviceID, gatewayNode, ttl)
}

func (r *redisPresenceService) GetUserRoutes(ctx context.Context, userID string) (map[string]string, error) {
	return r.client.HGetAll(ctx, presenceKey(userID)).Result()
}

func (r *redisPresenceService) IsUserOnline(ctx context.Context, userID string) (bool, error) {
	count, err := r.client.HLen(ctx, presenceKey(userID)).Result()
	if err != nil {
		return false, err
	}
	return count > 0, nil
}

func presenceKey(userID string) string {
	return fmt.Sprintf("presence:%s", userID)
}

