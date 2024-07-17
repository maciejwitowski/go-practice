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
	Started         bool
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

	if b.Started {
		return errors.New("token bucket already Started")
	}

	b.ticker = time.NewTicker(b.refillInterval)
	b.done = make(chan struct{})
	b.Started = true

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

	if !b.Started {
		return errors.New("token bucket already Started")
	}

	b.ticker.Stop()
	close(b.done)
	b.Started = false

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
