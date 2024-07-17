package ratelimit

import (
	"testing"
	"time"
)

func TestLeakingBucket(t *testing.T) {
	lb := NewLeakingBucket(100*time.Millisecond, 3)

	workers := make([]Worker, 3)
	for i := 0; i < 3; i++ {
		worker := Worker{ID: i + 1, in: lb.RequestsChannel, done: make(chan struct{})}
		workers[i] = worker
		worker.Start()
	}

	lb.Start()
	defer lb.Stop()

	request := Request{UserId: "A"}
	if err := lb.Add(request); err != nil {
		t.Errorf("unexpected error for %s", request)
	}

	request = Request{UserId: "B"}
	if err := lb.Add(request); err != nil {
		t.Errorf("unexpected error for %s", request)
	}

	request = Request{UserId: "C"}
	if err := lb.Add(request); err != nil {
		t.Errorf("unexpected error for %s", request)
	}

	request = Request{UserId: "D"}
	if err := lb.Add(request); err == nil {
		t.Errorf("expected error for %s", request)
	}

	// Wait so request A is sent for processing
	time.Sleep(100*time.Millisecond + 10*time.Millisecond)

	// D should be enqueued after workers consumed the previous ones
	if err := lb.Add(request); err != nil {
		t.Errorf("unexpected error for %s", request)
	}

	// Sleep enough to process all pending requests (B, C, D)
	time.Sleep(300*time.Millisecond + 10*time.Millisecond)

	if len(lb.RequestsQueue) != 0 {
		t.Errorf("expected queue to have been consumed")
	}

	for _, w := range workers {
		w.Stop()
	}
}
