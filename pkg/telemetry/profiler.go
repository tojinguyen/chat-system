package telemetry

import (
	"fmt"
	"log"
	"runtime"

	"github.com/grafana/pyroscope-go"
)

// ProfilerConfig configures Pyroscope continuous profiling
type ProfilerConfig struct {
	ServerAddress string // e.g. "http://pyroscope.observability:4040"
	ApplicationName string // e.g. "chat-system.ws-gateway"
	NodeID          string
	Environment     string
	Disabled        bool
}

// InitProfiler starts the Grafana Pyroscope continuous profiling agent
func InitProfiler(cfg ProfilerConfig) (func() error, error) {
	if cfg.Disabled || cfg.ServerAddress == "" {
		log.Printf("[Profiler] Continuous profiling disabled or server address empty.")
		return func() error { return nil }, nil
	}

	// Configure runtime profile rates for contention & blocking
	runtime.SetMutexProfileFraction(5)
	runtime.SetBlockProfileRate(10000)

	profiler, err := pyroscope.Start(pyroscope.Config{
		ApplicationName: cfg.ApplicationName,
		ServerAddress:   cfg.ServerAddress,
		Logger:          nil, // standard logger
		Tags: map[string]string{
			"env":     cfg.Environment,
			"node_id": cfg.NodeID,
		},
		ProfileTypes: []pyroscope.ProfileType{
			pyroscope.ProfileCPU,
			pyroscope.ProfileAllocObjects,
			pyroscope.ProfileAllocSpace,
			pyroscope.ProfileInuseObjects,
			pyroscope.ProfileInuseSpace,
			pyroscope.ProfileGoroutines,
			pyroscope.ProfileMutexCount,
			pyroscope.ProfileMutexDuration,
			pyroscope.ProfileBlockCount,
			pyroscope.ProfileBlockDuration,
		},
	})
	if err != nil {
		return nil, fmt.Errorf("failed to start pyroscope profiler: %w", err)
	}

	log.Printf("[Profiler] Continuous profiling active -> Pyroscope: %s (App: %s)",
		cfg.ServerAddress, cfg.ApplicationName)

	return func() error {
		log.Printf("[Profiler] Stopping Pyroscope profiler for %s...", cfg.ApplicationName)
		return profiler.Stop()
	}, nil
}
