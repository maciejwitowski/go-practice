package ratelimit

import (
	"errors"
	"sync"
	"time"
)

type TokenBucket struct {
	mu              sync.Mutex
	availableTokens int
	maxTokens       int
	refillInterval  time.Duration
	refillDelta     int
	ticker          *time.Ticker
	done            chan struct{}
	started         bool
}

func NewTokenBucket(availableTokens int, maxTokens int, refillInterval time.Duration, refillDelta int) *TokenBucket {
	if maxTokens < availableTokens {
		maxTokens = availableTokens
	}

	return &TokenBucket{
		availableTokens: availableTokens,
		maxTokens:       maxTokens,
		refillInterval:  refillInterval,
		refillDelta:     refillDelta,
	}
}

func (b *TokenBucket) Start() error {
	b.mu.Lock()
	defer b.mu.Unlock()

	if b.started {
		return errors.New("token bucket already started")
	}

	b.ticker = time.NewTicker(b.refillInterval)
	b.done = make(chan struct{})
	b.started = true

	go func() {
		for {
			select {
			case <-b.ticker.C:
				b.refill()
			case <-b.done:
				return
			}
		}
	}()

	return nil
}

func (b *TokenBucket) Stop() error {
	b.mu.Lock()
	defer b.mu.Unlock()

	if !b.started {
		return errors.New("token bucket already started")
	}

	b.ticker.Stop()
	close(b.done)
	b.started = false

	return nil
}

func (b *TokenBucket) IsOverflown() bool {
	b.mu.Lock()
	defer b.mu.Unlock()

	if b.availableTokens == 0 {
		return true
	}

	b.availableTokens--
	return false
}

func (b *TokenBucket) refill() {
	b.mu.Lock()
	defer b.mu.Unlock()

	b.availableTokens = min(b.availableTokens+b.refillDelta, b.maxTokens)
}
