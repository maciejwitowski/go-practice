package ratelimit

import (
	"errors"
	"time"
)

type SlidingWindow struct {
	Log      []TimestampedRequest
	capacity int
	duration time.Duration
	api      Api
}

func NewSlidingWindow(capacity int, duration time.Duration, api Api) *SlidingWindow {
	return &SlidingWindow{
		Log:      make([]TimestampedRequest, 0),
		capacity: capacity,
		duration: duration,
		api:      api,
	}
}

func (sw *SlidingWindow) Process(r TimestampedRequest) (ApiResult, error) {
	sw.Log = append(sw.Log, r)

	sw.cleanupOutdatedLogs()

	if len(sw.Log) > sw.capacity {
		return "", errors.New("capacity exceeded")
	} else {
		return sw.api.execute(r.Request)
	}
}

func (sw *SlidingWindow) cleanupOutdatedLogs() {
	windowStart := time.Now().Add(-sw.duration)
	for _, l := range sw.Log {
		if l.Timestamp.Before(windowStart) {
			// Remove outdated request
			sw.Log = sw.Log[1:]
		}
	}
}

type TimestampedRequest struct {
	Request
	Timestamp time.Time
}
