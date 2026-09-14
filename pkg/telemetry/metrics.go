package telemetry

import (
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promhttp"
)

var (
	// High-resolution latency buckets: 100µs up to 10s
	latencyBuckets = []float64{
		0.0001, 0.00025, 0.0005, 0.001, 0.0025, 0.005, 0.01, 0.025,
		0.05, 0.1, 0.25, 0.5, 1.0, 2.5, 5.0, 10.0,
	}

	// MessageLatency tracks latency across different stages of message delivery
	MessageLatency = prometheus.NewHistogramVec(
		prometheus.HistogramOpts{
			Namespace: "chat",
			Subsystem: "system",
			Name:      "message_latency_seconds",
			Help:      "Latency of chat message processing broken down by mode and stage.",
			Buckets:   latencyBuckets,
		},
		[]string{"mode", "stage", "status"},
	)

	// DispatchDuration tracks the exact execution duration of gRPC vs NATS outbound dispatch
	DispatchDuration = prometheus.NewHistogramVec(
		prometheus.HistogramOpts{
			Namespace: "chat",
			Subsystem: "engine",
			Name:      "dispatch_duration_seconds",
			Help:      "Duration of outbound dispatch to gateway (gRPC vs NATS).",
			Buckets:   latencyBuckets,
		},
		[]string{"mode", "node_id", "status"},
	)

	// ActiveConnections tracks the number of active WebSocket connections on a gateway
	ActiveConnections = prometheus.NewGaugeVec(
		prometheus.GaugeOpts{
			Namespace: "chat",
			Subsystem: "gateway",
			Name:      "active_connections",
			Help:      "Current active WebSocket client connections on this node.",
		},
		[]string{"node_id"},
	)

	// MessagesProcessed tracks message volume
	MessagesProcessed = prometheus.NewCounterVec(
		prometheus.CounterOpts{
			Namespace: "chat",
			Subsystem: "system",
			Name:      "messages_total",
			Help:      "Total number of messages processed.",
		},
		[]string{"service", "event_type", "status"},
	)
)

func init() {
	prometheus.MustRegister(
		MessageLatency,
		DispatchDuration,
		ActiveConnections,
		MessagesProcessed,
	)
}

// StartMetricsServer starts an HTTP endpoint on the given port to serve Prometheus/VictoriaMetrics
func StartMetricsServer(port int) *http.Server {
	if port <= 0 {
		return nil
	}

	mux := http.NewServeMux()
	mux.Handle("/metrics", promhttp.Handler())

	srv := &http.Server{
		Addr:         fmt.Sprintf(":%d", port),
		Handler:      mux,
		ReadTimeout:  5 * time.Second,
		WriteTimeout: 5 * time.Second,
	}

	go func() {
		log.Printf("[Metrics] Prometheus/VictoriaMetrics metrics listening on :%d/metrics", port)
		if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Printf("[Metrics] HTTP server error: %v", err)
		}
	}()

	return srv
}
