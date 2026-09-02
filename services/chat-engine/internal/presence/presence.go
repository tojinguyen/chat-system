package presence

import (
	"context"
	"time"
)

type UserDevicePresence struct {
	DeviceID    string
	GatewayNode string
	IsOnline    bool
}

// PresenceChecker queries Redis for user online/offline status and gateway routing
type PresenceChecker interface {
	GetUserPresence(ctx context.Context, userID string) ([]UserDevicePresence, error)
	CheckIdempotency(ctx context.Context, clientMsgID string, ttl time.Duration) (bool, error)
}

type redisPresenceChecker struct {
	// TODO: redis client reference
}

func NewRedisPresenceChecker(addr, password string, db int) PresenceChecker {
	return &redisPresenceChecker{}
}

func (r *redisPresenceChecker) GetUserPresence(ctx context.Context, userID string) ([]UserDevicePresence, error) {
	// Keys matching: presence:{user_id}:* -> returns list of active devices and their gateway nodes
	return nil, nil
}

func (r *redisPresenceChecker) CheckIdempotency(ctx context.Context, clientMsgID string, ttl time.Duration) (bool, error) {
	// SETNX idempotency:{client_msg_id} 1 EX {ttl}
	return true, nil
}
