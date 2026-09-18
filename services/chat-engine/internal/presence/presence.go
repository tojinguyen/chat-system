package presence

import (
	"context"
	"fmt"
	"time"

	"chat-system/pkg/telemetry"

	"github.com/redis/go-redis/v9"
	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/attribute"
	"go.opentelemetry.io/otel/codes"
	"go.opentelemetry.io/otel/trace"
)

var tracer = otel.Tracer("chat-worker/presence")

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
	key := presenceKey(userID)
	ctx, span := tracer.Start(ctx, "Redis.GetUserRoutes",
		trace.WithSpanKind(trace.SpanKindClient),
		trace.WithAttributes(
			attribute.String("db.system", "redis"),
			attribute.String("db.operation", "hgetall"),
			attribute.String("db.redis.key", key),
			attribute.String("chat.user_id", userID),
		),
	)
	defer span.End()

	startTime := time.Now()
	res, err := r.client.HGetAll(ctx, key).Result()
	duration := time.Since(startTime).Seconds()

	status := "success"
	if err != nil {
		status = "error"
		span.RecordError(err)
		span.SetStatus(codes.Error, err.Error())
		telemetry.DatabaseLatency.WithLabelValues("redis", "hgetall", status).Observe(duration)
		return nil, err
	}

	span.SetAttributes(attribute.Int("chat.devices_count", len(res)))
	span.SetStatus(codes.Ok, "OK")
	telemetry.DatabaseLatency.WithLabelValues("redis", "hgetall", status).Observe(duration)

	return res, nil
}

// IsUserOnline returns true if user has at least one active device online in Redis.
func (r *redisPresenceReader) IsUserOnline(ctx context.Context, userID string) (bool, error) {
	key := presenceKey(userID)
	ctx, span := tracer.Start(ctx, "Redis.IsUserOnline",
		trace.WithSpanKind(trace.SpanKindClient),
		trace.WithAttributes(
			attribute.String("db.system", "redis"),
			attribute.String("db.operation", "hlen"),
			attribute.String("db.redis.key", key),
			attribute.String("chat.user_id", userID),
		),
	)
	defer span.End()

	startTime := time.Now()
	count, err := r.client.HLen(ctx, key).Result()
	duration := time.Since(startTime).Seconds()

	status := "success"
	if err != nil {
		status = "error"
		span.RecordError(err)
		span.SetStatus(codes.Error, err.Error())
		telemetry.DatabaseLatency.WithLabelValues("redis", "hlen", status).Observe(duration)
		return false, err
	}

	online := count > 0
	span.SetAttributes(attribute.Bool("chat.is_online", online))
	span.SetStatus(codes.Ok, "OK")
	telemetry.DatabaseLatency.WithLabelValues("redis", "hlen", status).Observe(duration)

	return online, nil
}
