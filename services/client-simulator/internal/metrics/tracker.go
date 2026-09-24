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

	// Sample buffer để tính Percentiles cho Sender ACK (E2E Latency)
	samplesMu sync.Mutex
	samples   []float64 // Lưu latency bằng millisecond

	// Thống kê độ trễ mạng thuần túy (Network Delay / Ping-Pong RTT)
	netPingCount          atomic.Int64
	totalNetLatencyNanos  atomic.Int64
	netSamplesMu          sync.Mutex
	netSamples            []float64 // Lưu latency bằng millisecond
}

func NewTracker() *Tracker {
	return &Tracker{
		samples:    make([]float64, 0, 10000),
		netSamples: make([]float64, 0, 10000),
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

func (t *Tracker) RecordNetworkDelay(d time.Duration) {
	t.netPingCount.Add(1)
	nanos := d.Nanoseconds()
	t.totalNetLatencyNanos.Add(nanos)

	ms := float64(nanos) / 1e6
	t.netSamplesMu.Lock()
	if len(t.netSamples) < 50000 {
		t.netSamples = append(t.netSamples, ms)
	} else {
		t.netSamples = t.netSamples[len(t.netSamples)-25000:]
		t.netSamples = append(t.netSamples, ms)
	}
	t.netSamplesMu.Unlock()
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

func (t *Tracker) GetNetworkStats() (count int64, avgMs float64, p50, p90, p95, p99, maxMs float64) {
	count = t.netPingCount.Load()
	if count > 0 {
		avgMs = float64(t.totalNetLatencyNanos.Load()) / float64(count) / 1e6
	}

	t.netSamplesMu.Lock()
	n := len(t.netSamples)
	if n > 0 {
		sorted := make([]float64, n)
		copy(sorted, t.netSamples)
		t.netSamplesMu.Unlock()

		sort.Float64s(sorted)
		p50 = sorted[int(float64(n)*0.50)]
		p90 = sorted[int(float64(n)*0.90)]
		p95 = sorted[int(float64(n)*0.95)]
		p99 = sorted[int(float64(n)*0.99)]
		maxMs = sorted[n-1]
	} else {
		t.netSamplesMu.Unlock()
	}

	return
}

func (t *Tracker) PrintDashboard(activeBots, totalBots int, uptime time.Duration) {
	sent, ack, delivered, errs, reconnects, ackAvgMs, ackP50, ackP90, ackP95, ackP99, ackMaxMs := t.GetStats()
	pingCount, netAvgMs, netP50, netP90, netP95, netP99, netMaxMs := t.GetNetworkStats()

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
	fmt.Printf("Reconnects:       %-8d | Ping Checks:     %d\n", reconnects, pingCount)
	fmt.Println("------------------------------------------------------------------------------")
	fmt.Println("📡 Network Delay / RTT (Client <-> Gateway Ping):")
	if pingCount > 0 {
		fmt.Printf("  - Average:      %6.2f ms | p50 (Median): %6.2f ms\n", netAvgMs, netP50)
		fmt.Printf("  - p90:          %6.2f ms | p95:          %6.2f ms\n", netP90, netP95)
		fmt.Printf("  - p99:          %6.2f ms | Max:          %6.2f ms\n", netP99, netMaxMs)
	} else {
		fmt.Println("  (Đang thu thập mẫu ping heartbeat...)")
	}
	fmt.Println("------------------------------------------------------------------------------")
	fmt.Println("⚡ End-to-End Latency (Sender ACK):")
	if ack > 0 {
		fmt.Printf("  - Average:      %6.2f ms | p50 (Median): %6.2f ms\n", ackAvgMs, ackP50)
		fmt.Printf("  - p90:          %6.2f ms | p95:          %6.2f ms\n", ackP90, ackP95)
		fmt.Printf("  - p99:          %6.2f ms | Max:          %6.2f ms\n", ackP99, ackMaxMs)
		if pingCount > 0 && ackP50 >= netP50 {
			estProcessing := ackP50 - netP50
			fmt.Printf("  -> Ước tính thời gian xử lý Backend (Median E2E - Network): ~%.2f ms\n", estProcessing)
		}
	} else {
		fmt.Println("  (Chưa có message ACK)")
	}
	fmt.Println("==============================================================================")
	fmt.Println("Nhấn Ctrl+C để dừng kiểm thử.")
}
