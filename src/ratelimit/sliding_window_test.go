package ratelimit

import (
	"testing"
	"time"
)

func TestSlidingWindow(t *testing.T) {
	sw := NewSlidingWindow(2, 1*time.Second, BasicApi{})

	start := time.Now()

	request := TimestampedRequest{
		Request: Request{
			UserId: "A",
		},
		Timestamp: start.Add(10 * time.Millisecond),
	}

	if _, err := sw.Process(request); err != nil {
		t.Errorf("Unexpected error for %s", request)
	}

	request = TimestampedRequest{
		Request: Request{
			UserId: "B",
		},
		Timestamp: start.Add(20 * time.Millisecond),
	}

	if _, err := sw.Process(request); err != nil {
		t.Errorf("Unexpected error for %s", request)
	}

	request = TimestampedRequest{
		Request: Request{
			UserId: "C",
		},
		Timestamp: start.Add(900 * time.Millisecond),
	}

	if _, err := sw.Process(request); err == nil {
		t.Errorf("Expected error for %s", request)
	}

	time.Sleep(1100 * time.Millisecond)

	request = TimestampedRequest{
		Request: Request{
			UserId: "D",
		},
		Timestamp: start.Add(1200 * time.Millisecond),
	}

	if _, err := sw.Process(request); err != nil {
		t.Errorf("Unexpected error for %s: %s", request, err)
	}
}
