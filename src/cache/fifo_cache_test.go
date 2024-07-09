package cache

import (
	"testing"
)

func TestBasicCase(t *testing.T) {
	cache, _ := NewFIFOCache(1)

	_, ok := cache.Read("test")
	if ok {
		t.Error("Didn't expected the value")
	}

	cache.Write("test", 2)
	val, ok := cache.Read("test")
	if !ok {
		t.Error("Expected value")
	}

	if val != 2 {
		t.Errorf("Expected %d, actual %d", 2, val)
	}
}

func TestMaxSize(t *testing.T) {
	cache, _ := NewFIFOCache(2)

	cache.Write("one", 1)
	cache.Write("two", 2)

	if cache.Size() != 2 {
		t.Error("Expected Size to reach the max size")
	}

	cache.Write("three", 3)

	if cache.Size() != 2 {
		t.Error("Expected Size to stay at the max size")
	}

	_, ok := cache.Read("one")
	if ok {
		t.Error("One should have been pruned")
	}

	_, ok = cache.Read("two")
	if !ok {
		t.Error("Two should be available")
	}

	_, ok = cache.Read("three")
	if !ok {
		t.Error("Three should be available")
	}
}
