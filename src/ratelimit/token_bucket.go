package ratelimit

import (
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

func (b *TokenBucket) Start() {
	b.mu.Lock()
	defer b.mu.Unlock()
	
	b.ticker = time.NewTicker(b.refillInterval)
	b.done = make(chan struct{})

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
}

func (b *TokenBucket) Stop() {
	b.mu.Lock()
	defer b.mu.Unlock()

	b.ticker.Stop()
	close(b.done)
}

func (b *TokenBucket) IsOverflown() bool {
	b.mu.Lock()
	defer b.mu.Unlock()

	if b.availableTokens == 0 {
		return true
	} else {
		b.availableTokens--
		return false
	}
}

func (b *TokenBucket) refill() {
	b.mu.Lock()
	defer b.mu.Unlock()

	b.availableTokens += b.refillDelta
}
