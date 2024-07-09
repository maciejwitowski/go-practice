package cache

import "errors"

// Implement an LRU (Least Recently Used) cache with a given capacity
type LRUCache struct {
	items      map[string]interface{}
	linkedList *LinkedList
	MaxSize    int
}

func NewLRUCache(capacity int) (*LRUCache, error) {
	if capacity < 0 {
		return nil, errors.New("expected capacity > 0")
	}

	return &LRUCache{
		items:      make(map[string]interface{}),
		linkedList: EmptyLinkedList(),
		MaxSize:    capacity,
	}, nil
}

func (c *LRUCache) Read(key string) (interface{}, bool) {
	val, ok := c.items[key]
	if ok {
		c.linkedList.makeLatest(key)
	}
	return val, ok
}

// Write adds or updates a value in the cache
// It returns true if the value was added, false if it was updated
func (c *LRUCache) Write(key string, val interface{}) bool {
	if _, exists := c.items[key]; exists {
		c.linkedList.makeLatest(key)
		c.items[key] = val
		return false
	}

	if len(c.items) == c.MaxSize {
		previousOldest := c.linkedList.dropOldest()
		delete(c.items, previousOldest)
	}

	c.items[key] = val
	c.linkedList.makeLatest(key)
	return true
}

func (c *LRUCache) Size() int {
	return len(c.items)
}

type LinkedList struct {
	lookup map[string]*ListItem
	latest *ListItem
	oldest *ListItem
}

func (l *LinkedList) makeLatest(key string) {
	item, exists := l.lookup[key]
	if !exists {
		newItem := &ListItem{val: key}
		l.lookup[key] = newItem
		if l.latest == nil { // Empty list
			l.latest = newItem
			l.oldest = newItem
		} else {
			newItem.previous = l.latest
			l.latest.next = newItem
			l.latest = newItem
		}
	} else if item != l.latest {
		// Remove item from its current position
		if item.previous != nil {
			item.previous.next = item.next
		}
		if item.next != nil {
			item.next.previous = item.previous
		}
		if item == l.oldest {
			l.oldest = item.next
		}
		// Move item to the front
		item.previous = l.latest
		item.next = nil
		l.latest.next = item
		l.latest = item
	}
}

func (l *LinkedList) dropOldest() string {
	if l.oldest == nil {
		return ""
	}

	oldest := l.oldest
	if oldest.next != nil {
		oldest.next.previous = nil
	}
	l.oldest = oldest.next
	if l.oldest == nil {
		l.latest = nil
	}
	delete(l.lookup, oldest.val)
	return oldest.val
}

func EmptyLinkedList() *LinkedList {
	return &LinkedList{
		lookup: make(map[string]*ListItem),
		latest: nil,
		oldest: nil,
	}
}

type ListItem struct {
	val      string
	previous *ListItem
	next     *ListItem
}
