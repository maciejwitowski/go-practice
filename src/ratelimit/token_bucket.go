package ratelimit

import (
	"time"
)

type TokenBucket struct {
	availableTokens int
	refillInterval  time.Duration
	refillDelta     int
	ticker          *time.Ticker
	done            chan struct{}
}

func NewTokenBucket(tokens int, refillInterval time.Duration, refillDelta int) *TokenBucket {
	return &TokenBucket{
		availableTokens: tokens,
		refillInterval:  refillInterval,
		refillDelta:     refillDelta,
	}
}

func (b *TokenBucket) Start() {
	b.ticker = time.NewTicker(b.refillInterval)
	b.done = make(chan struct{})

	go func() {
		for {
			select {
			case <-b.ticker.C:
				b.availableTokens += b.refillDelta
			case <-b.done:
				return
			}
		}
	}()
}

func (b *TokenBucket) Stop() {
	b.ticker.Stop()
	close(b.done)
}

func (b *TokenBucket) IsOverflown() bool {
	if b.availableTokens == 0 {
		return true
	} else {
		b.availableTokens--
		return false
	}
}
