package worker

import (
	"context"
	"fmt"
	"sync"
	"sync/atomic"
	"testing"
	"time"
)

type sampleTask struct {
	Key     string
	Payload string
}

func TestPartitionedPool_DeterministicRouting(t *testing.T) {
	pool, err := NewPartitionedPool(Config[sampleTask]{
		NumWorkers: 16,
		BufferSize: 100,
		KeyExtractor: func(item sampleTask) string {
			return item.Key
		},
		Handler: func(ctx context.Context, item sampleTask) error {
			return nil
		},
	})
	if err != nil {
		t.Fatalf("Failed to create pool: %v", err)
	}

	key := "room-alpha-123"
	expectedIdx := pool.getWorkerIndex(key)

	for i := 0; i < 50; i++ {
		idx := pool.getWorkerIndex(key)
		if idx != expectedIdx {
			t.Fatalf("Expected worker %d, got %d on iteration %d", expectedIdx, idx, i)
		}
	}
}

func TestPartitionedPool_PanicRecovery(t *testing.T) {
	var processed atomic.Int32

	pool, err := NewPartitionedPool(Config[sampleTask]{
		NumWorkers: 4,
		BufferSize: 100,
		KeyExtractor: func(item sampleTask) string {
			return item.Key
		},
		Handler: func(ctx context.Context, item sampleTask) error {
			if item.Payload == "panic-payload" {
				panic("simulated poison pill panic")
			}
			processed.Add(1)
			return nil
		},
	})
	if err != nil {
		t.Fatalf("Failed to create pool: %v", err)
	}

	pool.Start()

	ctx := context.Background()

	// Gửi task gây panic
	_ = pool.Submit(ctx, sampleTask{Key: "room-1", Payload: "panic-payload"})

	// Gửi tiếp các task bình thường cùng key
	for i := 0; i < 10; i++ {
		_ = pool.Submit(ctx, sampleTask{Key: "room-1", Payload: fmt.Sprintf("good-%d", i)})
	}

	time.Sleep(100 * time.Millisecond)

	drainCtx, cancel := context.WithTimeout(context.Background(), 2*time.Second)
	defer cancel()

	if err := pool.Stop(drainCtx); err != nil {
		t.Fatalf("Failed to stop pool: %v", err)
	}

	if processed.Load() != 10 {
		t.Fatalf("Expected 10 processed messages, got %d", processed.Load())
	}
}

func TestPartitionedPool_GracefulDrain(t *testing.T) {
	var processed atomic.Int32
	var mu sync.Mutex
	order := make([]string, 0)

	pool, err := NewPartitionedPool(Config[sampleTask]{
		NumWorkers: 8,
		BufferSize: 500,
		KeyExtractor: func(item sampleTask) string {
			return item.Key
		},
		Handler: func(ctx context.Context, item sampleTask) error {
			time.Sleep(2 * time.Millisecond)
			processed.Add(1)
			mu.Lock()
			order = append(order, item.Payload)
			mu.Unlock()
			return nil
		},
	})
	if err != nil {
		t.Fatalf("Failed to create pool: %v", err)
	}

	pool.Start()

	totalMsgs := 100
	ctx := context.Background()

	for i := 0; i < totalMsgs; i++ {
		err := pool.Submit(ctx, sampleTask{
			Key:     fmt.Sprintf("conv-%d", i%5),
			Payload: fmt.Sprintf("msg-%d", i),
		})
		if err != nil {
			t.Fatalf("Failed to submit message: %v", err)
		}
	}

	drainCtx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	if err := pool.Stop(drainCtx); err != nil {
		t.Fatalf("Stop timed out or failed: %v", err)
	}

	if int(processed.Load()) != totalMsgs {
		t.Fatalf("Expected %d messages drained, got %d", totalMsgs, processed.Load())
	}
}
