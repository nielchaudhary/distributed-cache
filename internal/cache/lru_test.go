package cache

import (
	"testing"
)

func TestLRUCache(t *testing.T) {
	cache := NewLRUCache(2)

	cache.Put("key1", "value1")
	if value, ok := cache.Get("key1"); !ok || value != "value1" {
		t.Errorf("Expected value1, got %v", value)
	}

	cache.Put("key2", "value2")
	cache.Put("key3", "value3")
	if _, ok := cache.Get("key1"); ok {
		t.Error("Expected key1 to be evicted")
	}
}
