package repository

import (
	"context"
	"fmt"
	"time"

	"github.com/redis/go-redis/v9"
)

// IdempotencyRepository quản lý khóa chống trùng lặp trên Redis
type IdempotencyRepository interface {
	// AcquireLock cố gắng chiếm khóa idempotency cho clientMsgID.
	// Trả về (true, nil) nếu đây là tin nhắn mới chưa từng xuất hiện.
	// Trả về (false, nil) nếu tin nhắn này đã tồn tại trong Redis.
	AcquireLock(ctx context.Context, clientMsgID string) (bool, error)

	// ReleaseLock xóa khóa trên Redis khi quá trình xử lý gặp lỗi (như lỗi DB),
	// cho phép client có thể retry lại sau đó.
	ReleaseLock(ctx context.Context, clientMsgID string) error
}

type redisIdempotencyRepository struct {
	client *redis.Client
	ttl    time.Duration
}

// NewRedisIdempotencyRepository khởi tạo repository với Redis client và thời gian sống TTL
func NewRedisIdempotencyRepository(client *redis.Client, ttl time.Duration) IdempotencyRepository {
	if ttl <= 0 {
		ttl = 24 * time.Hour // Mặc định giữ 24h
	}
	return &redisIdempotencyRepository{
		client: client,
		ttl:    ttl,
	}
}

func (r *redisIdempotencyRepository) AcquireLock(ctx context.Context, clientMsgID string) (bool, error) {
	if clientMsgID == "" {
		return false, fmt.Errorf("clientMsgID cannot be empty")
	}

	key := fmt.Sprintf("idempotency:msg:%s", clientMsgID)

	// Lệnh SET key value NX EX ttl là thao tác nguyên tử (Atomic) trong Redis
	acquired, err := r.client.SetNX(ctx, key, "PROCESSED", r.ttl).Result()
	if err != nil {
		return false, fmt.Errorf("redis setnx failed for key %s: %w", key, err)
	}

	return acquired, nil
}

func (r *redisIdempotencyRepository) ReleaseLock(ctx context.Context, clientMsgID string) error {
	if clientMsgID == "" {
		return nil
	}

	key := fmt.Sprintf("idempotency:msg:%s", clientMsgID)
	if err := r.client.Del(ctx, key).Err(); err != nil {
		return fmt.Errorf("redis del failed for key %s: %w", key, err)
	}

	return nil
}
