package ratelimit

import (
	"errors"
	"fmt"
	"sync"
	"time"
)

type LeakingBucket struct {
	mu              sync.Mutex
	RequestsQueue   []Request
	RequestsChannel chan Request
	capacity        int
	ticker          *time.Ticker
	interval        time.Duration
	done            chan struct{}
}

func NewLeakingBucket(interval time.Duration, capacity int) *LeakingBucket {
	return &LeakingBucket{
		RequestsQueue:   make([]Request, 0),
		RequestsChannel: make(chan Request, 1),
		capacity:        capacity,
		mu:              sync.Mutex{},
		interval:        interval,
		done:            make(chan struct{}),
	}
}

func (lb *LeakingBucket) Start() {
	lb.mu.Lock()
	defer lb.mu.Unlock()

	lb.ticker = time.NewTicker(lb.interval)
	lb.done = make(chan struct{})

	go func() {
		for {
			select {
			case <-lb.ticker.C:
				if len(lb.RequestsQueue) == 0 {
					continue
				}

				requestToProcess := lb.RequestsQueue[0]
				fmt.Printf("Sending %s for processing\n", requestToProcess)
				lb.RequestsChannel <- requestToProcess
				lb.RequestsQueue = lb.RequestsQueue[1:]
			case <-lb.done:
				return
			}
		}
	}()
}

func (lb *LeakingBucket) Stop() {
	lb.mu.Lock()
	defer lb.mu.Unlock()

	close(lb.done)
}

func (lb *LeakingBucket) Add(r Request) error {
	lb.mu.Lock()
	defer lb.mu.Unlock()

	if len(lb.RequestsQueue) < lb.capacity {
		lb.RequestsQueue = append(lb.RequestsQueue, r)
		fmt.Printf("%s added. Queue size: %d\n", r, len(lb.RequestsQueue))
		return nil
	} else {
		fmt.Printf("Capacity reached, dropping the request %s\n", r)
		return errors.New("capacity reached")
	}

}

type Worker struct {
	ID   int
	in   <-chan Request
	done chan struct{}
}

func (w Worker) Start() {
	fmt.Printf("%s: Starting\n", w)

	go func() {
		for {
			select {
			case r := <-w.in:
				fmt.Printf("%s: processing %s\n", w, r)
			case <-w.done:
				return
			}
		}
	}()
}

func (w Worker) Stop() {
	fmt.Printf("Stopping %s\n", w)
	close(w.done)
}

func (w Worker) String() string {
	return fmt.Sprintf("Worker[ID=%d]", w.ID)
}
