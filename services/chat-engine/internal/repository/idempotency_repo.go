package repository

import (
	"context"
	"fmt"
	"time"

	"chat-system/pkg/telemetry"

	"github.com/redis/go-redis/v9"
	"go.opentelemetry.io/otel/attribute"
	"go.opentelemetry.io/otel/codes"
	"go.opentelemetry.io/otel/trace"
)

type IdempotencyRepository interface {
	AcquireLock(ctx context.Context, clientMsgID string) (bool, error)
	ReleaseLock(ctx context.Context, clientMsgID string) error
}

type redisIdempotencyRepository struct {
	client *redis.Client
	ttl    time.Duration
}

func NewRedisIdempotencyRepository(client *redis.Client, ttl time.Duration) IdempotencyRepository {
	if ttl <= 0 {
		ttl = 24 * time.Hour
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

	ctx, span := tracer.Start(ctx, "Redis.AcquireLock",
		trace.WithSpanKind(trace.SpanKindClient),
		trace.WithAttributes(
			attribute.String("db.system", "redis"),
			attribute.String("db.operation", "setnx"),
			attribute.String("db.redis.key", key),
			attribute.String("chat.client_msg_id", clientMsgID),
		),
	)
	defer span.End()

	startTime := time.Now()
	// Lệnh SET key value NX EX ttl là thao tác nguyên tử (Atomic) trong Redis
	acquired, err := r.client.SetNX(ctx, key, "PROCESSED", r.ttl).Result()
	duration := time.Since(startTime).Seconds()

	status := "success"
	if err != nil {
		status = "error"
		span.RecordError(err)
		span.SetStatus(codes.Error, err.Error())
		telemetry.DatabaseLatency.WithLabelValues("redis", "setnx", status).Observe(duration)
		return false, fmt.Errorf("redis setnx failed for key %s: %w", key, err)
	}

	span.SetAttributes(attribute.Bool("chat.lock_acquired", acquired))
	span.SetStatus(codes.Ok, "OK")
	telemetry.DatabaseLatency.WithLabelValues("redis", "setnx", status).Observe(duration)

	return acquired, nil
}

func (r *redisIdempotencyRepository) ReleaseLock(ctx context.Context, clientMsgID string) error {
	if clientMsgID == "" {
		return nil
	}

	key := fmt.Sprintf("idempotency:msg:%s", clientMsgID)

	ctx, span := tracer.Start(ctx, "Redis.ReleaseLock",
		trace.WithSpanKind(trace.SpanKindClient),
		trace.WithAttributes(
			attribute.String("db.system", "redis"),
			attribute.String("db.operation", "del"),
			attribute.String("db.redis.key", key),
			attribute.String("chat.client_msg_id", clientMsgID),
		),
	)
	defer span.End()

	startTime := time.Now()
	err := r.client.Del(ctx, key).Err()
	duration := time.Since(startTime).Seconds()

	status := "success"
	if err != nil {
		status = "error"
		span.RecordError(err)
		span.SetStatus(codes.Error, err.Error())
		telemetry.DatabaseLatency.WithLabelValues("redis", "del", status).Observe(duration)
		return fmt.Errorf("redis del failed for key %s: %w", key, err)
	}

	span.SetStatus(codes.Ok, "OK")
	telemetry.DatabaseLatency.WithLabelValues("redis", "del", status).Observe(duration)

	return nil
}
