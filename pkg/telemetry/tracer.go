package telemetry

import (
	"context"
	"fmt"
	"log"
	"time"

	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/exporters/otlp/otlptrace/otlptracegrpc"
	"go.opentelemetry.io/otel/propagation"
	"go.opentelemetry.io/otel/sdk/resource"
	sdktrace "go.opentelemetry.io/otel/sdk/trace"
	semconv "go.opentelemetry.io/otel/semconv/v1.26.0"
	"go.opentelemetry.io/otel/trace"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

// Config defines the configuration for telemetry setup
type Config struct {
	ServiceName     string
	ServiceVersion  string
	Environment     string
	NodeID          string
	CollectorTarget string // e.g. "localhost:4317" or "otel-collector.observability:4317"
	Disabled        bool
}

// InitTracer initializes OpenTelemetry TracerProvider and registers global propagator
func InitTracer(ctx context.Context, cfg Config) (func(context.Context) error, error) {
	if cfg.Disabled || cfg.CollectorTarget == "" {
		log.Printf("[Telemetry] Tracing disabled or collector endpoint empty. Running in NoOp mode.")
		return func(ctx context.Context) error { return nil }, nil
	}

	// 1. Setup Resource Attributes
	res, err := resource.New(ctx,
		resource.WithAttributes(
			semconv.ServiceNameKey.String(cfg.ServiceName),
			semconv.ServiceVersionKey.String(cfg.ServiceVersion),
			semconv.DeploymentEnvironmentKey.String(cfg.Environment),
			semconv.ServiceInstanceIDKey.String(cfg.NodeID),
		),
		resource.WithHost(),
		resource.WithProcess(),
	)
	if err != nil {
		return nil, fmt.Errorf("failed to create resource: %w", err)
	}

	// 2. Setup OTLP gRPC Exporter
	exporterCtx, cancel := context.WithTimeout(ctx, 5*time.Second)
	defer cancel()

	exporter, err := otlptracegrpc.New(exporterCtx,
		otlptracegrpc.WithInsecure(),
		otlptracegrpc.WithEndpoint(cfg.CollectorTarget),
		otlptracegrpc.WithDialOption(grpc.WithTransportCredentials(insecure.NewCredentials())),
	)
	if err != nil {
		return nil, fmt.Errorf("failed to initialize OTLP gRPC trace exporter: %w", err)
	}

	// 3. Batch Span Processor with high-throughput tuning to prevent OOM
	bsp := sdktrace.NewBatchSpanProcessor(exporter,
		sdktrace.WithMaxQueueSize(16384),
		sdktrace.WithMaxExportBatchSize(2048),
		sdktrace.WithBatchTimeout(1*time.Second),
		sdktrace.WithExportTimeout(5*time.Second),
	)

	tp := sdktrace.NewTracerProvider(
		sdktrace.WithSampler(sdktrace.AlwaysSample()),
		sdktrace.WithResource(res),
		sdktrace.WithSpanProcessor(bsp),
	)

	otel.SetTracerProvider(tp)

	// 4. Setup Global W3C TraceContext Propagator
	otel.SetTextMapPropagator(propagation.NewCompositeTextMapPropagator(
		propagation.TraceContext{},
		propagation.Baggage{},
	))

	log.Printf("[Telemetry] OpenTelemetry Tracing initialized -> Collector: %s (Service: %s)",
		cfg.CollectorTarget, cfg.ServiceName)

	// Return graceful shutdown function
	return func(shutdownCtx context.Context) error {
		log.Printf("[Telemetry] Flushing and shutting down TracerProvider for %s...", cfg.ServiceName)
		return tp.Shutdown(shutdownCtx)
	}, nil
}

// Tracer returns a named tracer instance
func Tracer(name string) trace.Tracer {
	return otel.Tracer(name)
}
