package metrics

import (
	"fmt"
	"sort"
	"sync"
	"sync/atomic"
	"time"
)

// Tracker quản lý số liệu đo lường hiệu năng và độ trễ (Latency)
type Tracker struct {
	totalSent         atomic.Int64
	totalAck          atomic.Int64
	totalDelivered    atomic.Int64
	totalErrors       atomic.Int64
	totalReconnects   atomic.Int64
	totalLatencyNanos atomic.Int64

	// Sample buffer để tính Percentiles (p50, p90, p95, p99)
	samplesMu sync.Mutex
	samples   []float64 // Lưu latency bằng millisecond
}

func NewTracker() *Tracker {
	return &Tracker{
		samples: make([]float64, 0, 10000),
	}
}

func (t *Tracker) RecordSent() {
	t.totalSent.Add(1)
}

func (t *Tracker) RecordAck(d time.Duration) {
	t.totalAck.Add(1)
	nanos := d.Nanoseconds()
	t.totalLatencyNanos.Add(nanos)

	ms := float64(nanos) / 1e6
	t.samplesMu.Lock()
	if len(t.samples) < 50000 {
		t.samples = append(t.samples, ms)
	} else {
		// Ring-buffer giữ tối đa 50.000 samples gần nhất
		t.samples = t.samples[len(t.samples)-25000:]
		t.samples = append(t.samples, ms)
	}
	t.samplesMu.Unlock()
}

func (t *Tracker) RecordDelivered() {
	t.totalDelivered.Add(1)
}

func (t *Tracker) RecordError() {
	t.totalErrors.Add(1)
}

func (t *Tracker) RecordReconnect() {
	t.totalReconnects.Add(1)
}

func (t *Tracker) GetStats() (sent, ack, delivered, errs, reconnects int64, avgMs float64, p50, p90, p95, p99, maxMs float64) {
	sent = t.totalSent.Load()
	ack = t.totalAck.Load()
	delivered = t.totalDelivered.Load()
	errs = t.totalErrors.Load()
	reconnects = t.totalReconnects.Load()

	if ack > 0 {
		avgMs = float64(t.totalLatencyNanos.Load()) / float64(ack) / 1e6
	}

	t.samplesMu.Lock()
	n := len(t.samples)
	if n > 0 {
		sorted := make([]float64, n)
		copy(sorted, t.samples)
		t.samplesMu.Unlock()

		sort.Float64s(sorted)
		p50 = sorted[int(float64(n)*0.50)]
		p90 = sorted[int(float64(n)*0.90)]
		p95 = sorted[int(float64(n)*0.95)]
		p99 = sorted[int(float64(n)*0.99)]
		maxMs = sorted[n-1]
	} else {
		t.samplesMu.Unlock()
	}

	return
}

func (t *Tracker) PrintDashboard(activeBots, totalBots int, uptime time.Duration) {
	sent, ack, delivered, errs, reconnects, avgMs, p50, p90, p95, p99, maxMs := t.GetStats()

	secs := uptime.Seconds()
	var sendTps, ackTps float64
	if secs > 0 {
		sendTps = float64(sent) / secs
		ackTps = float64(ack) / secs
	}

	var ackRate float64
	if sent > 0 {
		ackRate = float64(ack) / float64(sent) * 100
	}

	fmt.Printf("\033[H\033[2J") // Clear console screen
	fmt.Println("========================= CHAT STRESS TEST DASHBOARD =========================")
	fmt.Printf("Active Bots:      %d / %d Online (WS Connected)\n", activeBots, totalBots)
	fmt.Printf("Elapsed Time:     %s\n", uptime.Round(time.Second))
	fmt.Println("------------------------------------------------------------------------------")
	fmt.Printf("Messages Sent:    %-8d | Send Rate:       %.1f msgs/sec\n", sent, sendTps)
	fmt.Printf("ACKs Received:    %-8d | ACK Rate:        %.1f%% (%.1f acks/sec)\n", ack, ackRate, ackTps)
	fmt.Printf("Msg Delivered:    %-8d | Errors:          %d\n", delivered, errs)
	fmt.Printf("Reconnects:       %-8d\n", reconnects)
	fmt.Println("------------------------------------------------------------------------------")
	fmt.Println("Round-Trip Latency (Sender ACK):")
	fmt.Printf("  - Average:      %6.2f ms\n", avgMs)
	fmt.Printf("  - p50 (Median): %6.2f ms\n", p50)
	fmt.Printf("  - p90:          %6.2f ms\n", p90)
	fmt.Printf("  - p95:          %6.2f ms\n", p95)
	fmt.Printf("  - p99:          %6.2f ms\n", p99)
	fmt.Printf("  - Max:          %6.2f ms\n", maxMs)
	fmt.Println("==============================================================================")
	fmt.Println("Nhấn Ctrl+C để dừng kiểm thử.")
}
