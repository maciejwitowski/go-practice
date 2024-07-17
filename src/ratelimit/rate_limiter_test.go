package ratelimit

import (
	"net"
	"testing"
	"time"
)

func TestRateLimiterHappyPath(t *testing.T) {
	bucketFactory := func() *TokenBucket {
		return NewTokenBucket(1, 1, 2*time.Millisecond, 1)
	}

	limiter := NewRateLimiter(UserIdPolicy{}, BasicApi{}, bucketFactory)

	requestFromA := Request{
		UserId:    "John",
		IpAddress: net.IPAddr{},
	}

	requestFromB := Request{
		UserId:    "Ann",
		IpAddress: net.IPAddr{},
	}

	// User A can send 1 request
	if _, err := limiter.Process(requestFromA); err != nil {
		t.Errorf("unexpected error")
	}

	// On 2nd request, user A should get error
	if _, err := limiter.Process(requestFromA); err == nil {
		t.Errorf("expected error")
	}

	// Request from User B should be processed
	if _, err := limiter.Process(requestFromB); err != nil {
		t.Errorf("unexpected error: %s", err)
	}

	// Wait for buckets refill
	time.Sleep(3 * time.Millisecond)

	// Requests for both users should be processed
	if _, err := limiter.Process(requestFromA); err != nil {
		t.Errorf("unexpected error")
	}

	if _, err := limiter.Process(requestFromB); err != nil {
		t.Errorf("unexpected error")
	}

	err := limiter.StopAll()
	if err != nil {
		t.Errorf("unexpected error")
	}
}
