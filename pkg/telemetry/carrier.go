package telemetry

import (
	"context"

	"github.com/nats-io/nats.go"
	"go.opentelemetry.io/otel"
	"google.golang.org/grpc/metadata"
)

// NATSHeaderCarrier adapts nats.Header to propagation.TextMapCarrier
type NATSHeaderCarrier nats.Header

func (c NATSHeaderCarrier) Get(key string) string {
	return nats.Header(c).Get(key)
}

func (c NATSHeaderCarrier) Set(key string, value string) {
	nats.Header(c).Set(key, value)
}

func (c NATSHeaderCarrier) Keys() []string {
	keys := make([]string, 0, len(c))
	for k := range c {
		keys = append(keys, k)
	}
	return keys
}

// InjectNATSTraceContext injects the OpenTelemetry TraceContext into nats.Msg Header
func InjectNATSTraceContext(ctx context.Context, msg *nats.Msg) {
	if msg.Header == nil {
		msg.Header = make(nats.Header)
	}
	otel.GetTextMapPropagator().Inject(ctx, NATSHeaderCarrier(msg.Header))
}

// ExtractNATSTraceContext extracts OpenTelemetry TraceContext from nats.Msg Header into a new context
func ExtractNATSTraceContext(ctx context.Context, msg *nats.Msg) context.Context {
	if msg.Header == nil {
		return ctx
	}
	return otel.GetTextMapPropagator().Extract(ctx, NATSHeaderCarrier(msg.Header))
}

// GRPCMetadataCarrier adapts metadata.MD to propagation.TextMapCarrier
type GRPCMetadataCarrier metadata.MD

func (c GRPCMetadataCarrier) Get(key string) string {
	vals := metadata.MD(c).Get(key)
	if len(vals) > 0 {
		return vals[0]
	}
	return ""
}

func (c GRPCMetadataCarrier) Set(key string, value string) {
	metadata.MD(c).Set(key, value)
}

func (c GRPCMetadataCarrier) Keys() []string {
	keys := make([]string, 0, len(c))
	for k := range c {
		keys = append(keys, k)
	}
	return keys
}

// InjectGRPCTraceContext injects the trace context into outgoing gRPC context metadata
func InjectGRPCTraceContext(ctx context.Context) context.Context {
	md, ok := metadata.FromOutgoingContext(ctx)
	if !ok {
		md = metadata.New(nil)
	} else {
		md = md.Copy()
	}
	otel.GetTextMapPropagator().Inject(ctx, GRPCMetadataCarrier(md))
	return metadata.NewOutgoingContext(ctx, md)
}

// ExtractGRPCTraceContext extracts trace context from incoming gRPC context metadata
func ExtractGRPCTraceContext(ctx context.Context) context.Context {
	md, ok := metadata.FromIncomingContext(ctx)
	if !ok {
		return ctx
	}
	return otel.GetTextMapPropagator().Extract(ctx, GRPCMetadataCarrier(md))
}
