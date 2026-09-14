package telemetry

import (
	"context"
	"log"
)

// SetupConfig đóng gói toàn bộ cấu hình cho 3 trụ cột Telemetry
type SetupConfig struct {
	ServiceName      string
	ServiceVersion   string
	Environment      string // Mặc định là "production" nếu để trống
	NodeID           string
	CollectorTarget  string
	MetricsPort      int
	ProfilerServer   string
	DisableTracing   bool
	DisableProfiling bool
}

// Setup là hàm Facade khởi tạo đồng bộ cả 3 hệ thống:
// 1. HTTP Metrics Server (:metrics_port) cho Prometheus/VictoriaMetrics
// 2. OpenTelemetry TracerProvider đẩy sang OTel Collector (:collector_target)
// 3. Grafana Pyroscope Profiler Agent đẩy sang Pyroscope Server (:server_address)
//
// Trả về duy nhất 1 hàm cleanup để thực hiện Graceful Shutdown theo đúng thứ tự an toàn:
// Profiler -> Metrics Server -> TracerProvider (Flush toàn bộ trace spans còn tồn đọng).
func Setup(ctx context.Context, cfg SetupConfig) (func(context.Context) error, error) {
	if cfg.Environment == "" {
		cfg.Environment = "production"
	}
	if cfg.ServiceVersion == "" {
		cfg.ServiceVersion = "1.0.0"
	}

	// 1. Khởi động Metrics Server
	metricsServer := StartMetricsServer(cfg.MetricsPort)

	// 2. Khởi tạo OpenTelemetry Tracing
	shutdownTracer, errTracer := InitTracer(ctx, Config{
		ServiceName:     cfg.ServiceName,
		ServiceVersion:  cfg.ServiceVersion,
		Environment:     cfg.Environment,
		NodeID:          cfg.NodeID,
		CollectorTarget: cfg.CollectorTarget,
		Disabled:        cfg.DisableTracing,
	})
	if errTracer != nil {
		log.Printf("[Telemetry] Warning: failed to initialize tracer: %v", errTracer)
	}

	// 3. Khởi tạo Continuous Profiler (Grafana Pyroscope)
	stopProfiler, errProfiler := InitProfiler(ProfilerConfig{
		ServerAddress:   cfg.ProfilerServer,
		ApplicationName: "chat-system." + cfg.ServiceName,
		NodeID:          cfg.NodeID,
		Environment:     cfg.Environment,
		Disabled:        cfg.DisableProfiling,
	})
	if errProfiler != nil {
		log.Printf("[Telemetry] Warning: failed to start profiler: %v", errProfiler)
	}

	cleanup := func(shutdownCtx context.Context) error {
		log.Printf("[Telemetry] Shutting down telemetry subsystems for %s...", cfg.ServiceName)
		if stopProfiler != nil {
			_ = stopProfiler()
		}
		if metricsServer != nil {
			_ = metricsServer.Shutdown(shutdownCtx)
		}
		if shutdownTracer != nil {
			_ = shutdownTracer(shutdownCtx)
		}
		return nil
	}

	if errTracer != nil {
		return cleanup, errTracer
	}
	if errProfiler != nil {
		return cleanup, errProfiler
	}

	return cleanup, nil
}
