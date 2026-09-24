package telemetry

import (
	"context"
	"fmt"
	"time"

	"go.opentelemetry.io/otel/attribute"
	"go.opentelemetry.io/otel/trace"
	"google.golang.org/grpc"
)

type successReporter interface {
	GetSuccess() bool
}

// UnaryClientInterceptor bọc ngoài tất cả các lệnh gRPC gọi từ Client (ví dụ chat-engine gọi ws-gateway)
// Tự động:
// 1. Gắn W3C TraceContext vào outgoing metadata
// 2. Tạo client span OpenTelemetry
// 3. Đo thời gian thực thi và ghi nhận vào DispatchDuration và MessageLatency
func UnaryClientInterceptor(targetNode string) grpc.UnaryClientInterceptor {
	return func(
		ctx context.Context,
		method string,
		req, reply any,
		cc *grpc.ClientConn,
		invoker grpc.UnaryInvoker,
		opts ...grpc.CallOption,
	) error {
		outCtx := InjectGRPCTraceContext(ctx)
		tracer := Tracer("grpc-client")
		spanCtx, span := tracer.Start(outCtx, fmt.Sprintf("grpc.call %s", method),
			trace.WithSpanKind(trace.SpanKindClient),
			trace.WithAttributes(
				attribute.String("rpc.method", method),
				attribute.String("gateway_node", targetNode),
			),
		)
		defer span.End()

		start := time.Now()
		err := invoker(spanCtx, method, req, reply, cc, opts...)
		duration := time.Since(start).Seconds()

		status := "success"
		if err != nil {
			status = "error"
		} else if sr, ok := reply.(successReporter); ok && !sr.GetSuccess() {
			status = "rejected"
		}

		DispatchDuration.WithLabelValues("grpc", targetNode, status).Observe(duration)
		MessageLatency.WithLabelValues("grpc", "outbound_dispatch", status).Observe(duration)

		return err
	}
}

// UnaryServerInterceptor bọc ngoài tất cả các request gRPC gửi đến Server (ví dụ ws-gateway)
// Tự động:
// 1. Trích xuất W3C TraceContext từ incoming metadata
// 2. Tạo server span OpenTelemetry
// 3. Đo thời gian thực thi và ghi nhận vào MessageLatency
func UnaryServerInterceptor(serviceName, stage string) grpc.UnaryServerInterceptor {
	return func(
		ctx context.Context,
		req any,
		info *grpc.UnaryServerInfo,
		handler grpc.UnaryHandler,
	) (any, error) {
		trCtx := ExtractGRPCTraceContext(ctx)
		tracer := Tracer(serviceName)
		spanCtx, span := tracer.Start(trCtx, fmt.Sprintf("grpc.server %s", info.FullMethod),
			trace.WithSpanKind(trace.SpanKindServer),
			trace.WithAttributes(
				attribute.String("rpc.method", info.FullMethod),
			),
		)
		defer span.End()

		start := time.Now()
		resp, err := handler(spanCtx, req)
		duration := time.Since(start).Seconds()

		status := "success"
		if err != nil {
			status = "error"
		}

		MessageLatency.WithLabelValues("grpc", stage, status).Observe(duration)
		return resp, err
	}
}
