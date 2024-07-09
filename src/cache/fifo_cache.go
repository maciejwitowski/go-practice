package cache

import (
	"errors"
	"sync"
)

type FIFOCache struct {
	items   map[string]interface{}
	queue   *Queue
	MaxSize int
	mu      sync.RWMutex
}

func NewFIFOCache(maxSize int) (*FIFOCache, error) {
	if maxSize < 0 {
		return nil, errors.New("expected maxSize > 0")
	}

	return &FIFOCache{
		items:   make(map[string]interface{}),
		queue:   EmptyQueue(),
		MaxSize: maxSize,
	}, nil
}

func (c *FIFOCache) Read(key string) (interface{}, bool) {
	c.mu.RLock()
	defer c.mu.RUnlock()

	val, ok := c.items[key]
	return val, ok
}

// Write adds or updates a value in the cache
// It returns true if the value was added, false if it was updated
func (c *FIFOCache) Write(key string, val interface{}) bool {
	c.mu.Lock()
	defer c.mu.Unlock()

	if _, exists := c.items[key]; exists {
		c.items[key] = val
		return false
	}

	if len(c.items) == c.MaxSize {
		previousHead := c.queue.dropHead()
		delete(c.items, previousHead)
	}

	c.items[key] = val
	c.queue.add(key)
	return true
}

func (c *FIFOCache) Size() int {
	return len(c.items)
}
