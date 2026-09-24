package telemetry_test

import (
	"context"
	"errors"
	"testing"
	"time"

	"chat-system/pkg/telemetry"
)

func TestSetup_Disabled(t *testing.T) {
	cleanup, err := telemetry.Setup(context.Background(), telemetry.SetupConfig{
		ServiceName:      "test-service",
		ServiceVersion:   "1.0.0",
		NodeID:           "test-node-1",
		MetricsPort:      0, // Không start port
		DisableTracing:   true,
		DisableProfiling: true,
	})
	if err != nil {
		t.Fatalf("expected nil error when telemetry disabled, got: %v", err)
	}
	if cleanup == nil {
		t.Fatal("expected cleanup function, got nil")
	}

	ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)
	defer cancel()
	if err := cleanup(ctx); err != nil {
		t.Fatalf("expected clean shutdown, got: %v", err)
	}
}

func TestInstrumentHandler(t *testing.T) {
	cfg := telemetry.HandlerConfig{
		Service:   "test-service",
		Mode:      "nats",
		Stage:     "test_stage",
		EventType: "test_event",
	}

	// Case 1: Handler success
	successHandler := telemetry.InstrumentHandler(cfg, func(ctx context.Context, msg string) error {
		if msg != "hello" {
			t.Fatalf("unexpected message: %s", msg)
		}
		return nil
	})

	if err := successHandler(context.Background(), "hello"); err != nil {
		t.Fatalf("expected nil error, got: %v", err)
	}

	// Case 2: Handler error
	errExpected := errors.New("boom")
	errorHandler := telemetry.InstrumentHandler(cfg, func(ctx context.Context, msg string) error {
		return errExpected
	})

	if err := errorHandler(context.Background(), "fail"); !errors.Is(err, errExpected) {
		t.Fatalf("expected error %v, got: %v", errExpected, err)
	}
}
