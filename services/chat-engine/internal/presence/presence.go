package presence

import (
	"context"
	"fmt"

	"github.com/redis/go-redis/v9"
)

type PresenceReader interface {
	GetUserRoutes(ctx context.Context, userID string) (map[string]string, error)
	IsUserOnline(ctx context.Context, userID string) (bool, error)
}

type redisPresenceReader struct {
	client *redis.Client
}

func NewPresenceReader(client *redis.Client) PresenceReader {
	return &redisPresenceReader{
		client: client,
	}
}

func presenceKey(userID string) string {
	return fmt.Sprintf("presence:%s", userID)
}

// GetUserRoutes retrieves all active device -> gatewayNode mappings for a user.
// Returns an empty map if user is offline or has no active devices.
func (r *redisPresenceReader) GetUserRoutes(ctx context.Context, userID string) (map[string]string, error) {
	return r.client.HGetAll(ctx, presenceKey(userID)).Result()
}

// IsUserOnline returns true if user has at least one active device online in Redis.
func (r *redisPresenceReader) IsUserOnline(ctx context.Context, userID string) (bool, error) {
	count, err := r.client.HLen(ctx, presenceKey(userID)).Result()
	if err != nil {
		return false, err
	}
	return count > 0, nil
}
