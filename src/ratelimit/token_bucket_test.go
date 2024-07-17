package ratelimit

import (
	"testing"
	"time"
)

func TestTokenBucketHappyPath(t *testing.T) {
	tb := NewTokenBucket(4, 2*time.Millisecond, 2)

	tb.Start()
	defer tb.Stop()

	for i := 1; i <= 4; i++ {
		if tb.IsOverflown() {
			t.Errorf("Request %d shouldn't have overflown", i)
		}
	}

	if !tb.IsOverflown() {
		t.Errorf("Request should have overflown")
	}

	// advance time somehow by 2 seconds to get 2 more tokens
	time.Sleep(2*time.Millisecond + 1*time.Millisecond)

	if tb.IsOverflown() {
		t.Errorf("Request shouldn't have overflown")
	}

	if tb.IsOverflown() {
		t.Errorf("Request shouldn't have overflown")
	}

	if !tb.IsOverflown() {
		t.Errorf("Request should have overflown")
	}
}

func TestTokenBucketStartAndStop(t *testing.T) {
	tb := NewTokenBucket(4, 2*time.Millisecond, 2)

	tb.Start()

	// Use all tokens
	for i := 1; i <= 4; i++ {
		tb.IsOverflown()
	}

	if !tb.IsOverflown() {
		t.Errorf("All tokens should have been used")
	}

	// Stopped not to refill
	tb.Stop()

	// Advance shouldn't have an effect since the times has been stopped
	time.Sleep(2*time.Millisecond + 1*time.Millisecond)

	if !tb.IsOverflown() {
		t.Errorf("Tokens shouldn't have been refilled becaue it's stopped")
	}

	tb.Start()
	defer tb.Stop()
	if !tb.IsOverflown() {
		t.Errorf("Tokens shouldn't have been refilled becaue it just started again")
	}

	// This advance will have an effect
	time.Sleep(2*time.Millisecond + 1*time.Millisecond)
	if tb.IsOverflown() {
		t.Errorf("A token should be available")
	}

	if tb.IsOverflown() {
		t.Errorf("A token should be available")
	}

	if !tb.IsOverflown() {
		t.Errorf("There should be no tokens")
	}
}
