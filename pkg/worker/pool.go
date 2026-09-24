package worker

import (
	"context"
	"errors"
	"fmt"
	"log"
	"strconv"
	"sync"
	"sync/atomic"

	"chat-system/pkg/telemetry"
)

const (
	DefaultNumWorkers = 32
	DefaultBufferSize = 1024
)

var (
	ErrPoolClosed   = errors.New("worker pool is closed")
	ErrPoolDraining = errors.New("worker pool timed out while draining")
)

// HandlerFunc định nghĩa hàm xử lý sự kiện trong Worker theo kiểu generic
type HandlerFunc[T any] func(ctx context.Context, item T) error

// KeyExtractorFunc định nghĩa hàm trích xuất routing key (ví dụ: conversation_id) từ item
type KeyExtractorFunc[T any] func(item T) string

// Job đại diện cho một tác vụ kèm Trace Context được chuyển tiếp sang Worker channel
type Job[T any] struct {
	Ctx  context.Context
	Item T
}

// Worker đại diện cho một goroutine xử lý độc lập trên channel riêng
type Worker[T any] struct {
	id      int
	idStr   string
	jobCh   chan Job[T]
	handler HandlerFunc[T]
}

// PartitionedPool quản lý N workers với cơ chế băm partition theo key
type PartitionedPool[T any] struct {
	numWorkers   int
	bufferSize   int
	workers      []*Worker[T]
	handler      HandlerFunc[T]
	keyExtractor KeyExtractorFunc[T]
	wg           sync.WaitGroup
	isClosed     atomic.Bool
}

// Config cấu hình generic cho PartitionedPool
type Config[T any] struct {
	NumWorkers   int
	BufferSize   int
	KeyExtractor KeyExtractorFunc[T]
	Handler      HandlerFunc[T]
}

// NewPartitionedPool khởi tạo một PartitionedPool generic
func NewPartitionedPool[T any](cfg Config[T]) (*PartitionedPool[T], error) {
	if cfg.NumWorkers <= 0 {
		cfg.NumWorkers = DefaultNumWorkers
	}
	if cfg.BufferSize <= 0 {
		cfg.BufferSize = DefaultBufferSize
	}
	if cfg.Handler == nil {
		return nil, errors.New("worker pool handler cannot be nil")
	}
	if cfg.KeyExtractor == nil {
		return nil, errors.New("worker pool key extractor cannot be nil")
	}

	pool := &PartitionedPool[T]{
		numWorkers:   cfg.NumWorkers,
		bufferSize:   cfg.BufferSize,
		workers:      make([]*Worker[T], cfg.NumWorkers),
		handler:      cfg.Handler,
		keyExtractor: cfg.KeyExtractor,
	}

	for i := 0; i < cfg.NumWorkers; i++ {
		idStr := strconv.Itoa(i)
		pool.workers[i] = &Worker[T]{
			id:      i,
			idStr:   idStr,
			jobCh:   make(chan Job[T], cfg.BufferSize),
			handler: cfg.Handler,
		}

		// Initialize channel metrics
		telemetry.WorkerChannelCapacity.WithLabelValues(idStr).Set(float64(cfg.BufferSize))
		telemetry.WorkerChannelDepth.WithLabelValues(idStr).Set(0)
		telemetry.WorkerChannelSaturation.WithLabelValues(idStr).Set(0)
	}

	return pool, nil
}

// Start khởi chạy toàn bộ worker goroutines
func (p *PartitionedPool[T]) Start() {
	for i := 0; i < p.numWorkers; i++ {
		p.wg.Add(1)
		go p.runWorker(p.workers[i])
	}
	log.Printf("[WorkerPool] Started %d partitioned workers (buffer_size=%d per worker)", p.numWorkers, p.bufferSize)
}

// runWorker thực thi vòng lặp nhận job của từng worker kèm cơ chế Panic Recovery per-job
func (p *PartitionedPool[T]) runWorker(w *Worker[T]) {
	defer p.wg.Done()

	for job := range w.jobCh {
		// Update metrics after taking a job from queue
		depth := len(w.jobCh)
		telemetry.WorkerChannelDepth.WithLabelValues(w.idStr).Set(float64(depth))
		telemetry.WorkerChannelSaturation.WithLabelValues(w.idStr).Set(float64(depth) / float64(p.bufferSize))

		p.processJobWithRecovery(w, job)
	}
}

// processJobWithRecovery bọc xử lý job trong defer recover để cô lập sự cố (Panic Isolation)
func (p *PartitionedPool[T]) processJobWithRecovery(w *Worker[T], job Job[T]) {
	defer func() {
		if r := recover(); r != nil {
			telemetry.WorkerJobsTotal.WithLabelValues(w.idStr, "panic").Inc()
			log.Printf("[Worker %d] CRITICAL: Panic recovered while processing job: %v", w.id, r)
		}
	}()

	if err := w.handler(job.Ctx, job.Item); err != nil {
		telemetry.WorkerJobsTotal.WithLabelValues(w.idStr, "error").Inc()
		log.Printf("[Worker %d] Error processing job: %v", w.id, err)
		return
	}

	telemetry.WorkerJobsTotal.WithLabelValues(w.idStr, "success").Inc()
}

// Submit định tuyến item vào channel của worker tương ứng theo hash(keyExtractor(item))
func (p *PartitionedPool[T]) Submit(ctx context.Context, item T) error {
	if p.isClosed.Load() {
		return ErrPoolClosed
	}

	key := p.keyExtractor(item)
	workerIdx := p.getWorkerIndex(key)
	targetWorker := p.workers[workerIdx]

	job := Job[T]{
		Ctx:  ctx,
		Item: item,
	}

	select {
	case targetWorker.jobCh <- job:
		depth := len(targetWorker.jobCh)
		telemetry.WorkerChannelDepth.WithLabelValues(targetWorker.idStr).Set(float64(depth))
		telemetry.WorkerChannelSaturation.WithLabelValues(targetWorker.idStr).Set(float64(depth) / float64(p.bufferSize))
		return nil
	case <-ctx.Done():
		return ctx.Err()
	}
}

// Stop đóng toàn bộ worker channels và chờ xả cạn backlog (Graceful Drain) kèm timeout context
func (p *PartitionedPool[T]) Stop(ctx context.Context) error {
	if !p.isClosed.CompareAndSwap(false, true) {
		return nil // Đã đóng trước đó
	}

	log.Printf("[WorkerPool] Initiating graceful drain for %d workers...", p.numWorkers)

	// Đóng từng channel để báo hiệu cho workers không còn job mới
	for _, w := range p.workers {
		close(w.jobCh)
	}

	done := make(chan struct{})
	go func() {
		p.wg.Wait()
		close(done)
	}()

	select {
	case <-done:
		log.Println("[WorkerPool] All workers drained and stopped cleanly.")
		return nil
	case <-ctx.Done():
		log.Printf("[WorkerPool] WARNING: Drain timeout reached: %v", ctx.Err())
		return fmt.Errorf("%w: %v", ErrPoolDraining, ctx.Err())
	}
}

// NumWorkers trả về số lượng workers
func (p *PartitionedPool[T]) NumWorkers() int {
	return p.numWorkers
}

// getWorkerIndex tính chỉ số worker bằng thuật toán FNV-1a 32-bit
func (p *PartitionedPool[T]) getWorkerIndex(key string) int {
	hash := fnv32a(key)
	return int(hash % uint32(p.numWorkers))
}

// fnv32a triển khai thuật toán băm FNV-1a 32-bit (Non-cryptographic, zero-alloc)
func fnv32a(s string) uint32 {
	var h uint32 = 2166136261
	for i := 0; i < len(s); i++ {
		h ^= uint32(s[i])
		h *= 16777619
	}
	return h
}
