package cache

import (
	"fmt"
	"testing"
	"time"
)

func TestCache_Clear(t *testing.T) {
	// Test case 1: Clear an empty cache
	cache := New[string, int](10*time.Second, time.Minute)
	cache.Clear()
	if len(cache.storage) != 0 {
		t.Errorf("Test case 1 failed: expected empty cache, got %v", cache.storage)
	}

	// Test case 2: Clear a cache with multiple items
	cache = New[string, int](10*time.Second, time.Minute)
	cache.Set("key1", 1)
	cache.Set("key2", 2)
	cache.Set("key3", 3)
	cache.Clear()
	if len(cache.storage) != 0 {
		t.Errorf("Test case 2 failed: expected empty cache, got %v", cache.storage)
	}

	// Test case 3: Clear a cache with expired items
	cache = New[string, int](1*time.Millisecond, time.Minute)
	cache.Set("key1", 1)
	time.Sleep(2 * time.Millisecond) // Wait for the item to expire
	cache.Clear()
	if len(cache.storage) != 0 {
		t.Errorf("Test case 3 failed: expected empty cache, got %v", cache.storage)
	}

	// Test case 4: Clear a cache with a mix of expired and non-expired items
	cache = New[string, int](10*time.Second, time.Minute)
	cache.Set("key1", 1)
	cache.Set("key2", 2)
	time.Sleep(5 * time.Millisecond) // Wait for some items to expire
	cache.Set("key3", 3)
	cache.Clear()
	if len(cache.storage) != 0 {
		t.Errorf("Test case 4 failed: expected empty cache, got %v", cache.storage)
	}

	// Test case 5: Clear a cache with a large number of items
	cache = New[string, int](10*time.Second, time.Minute)
	for i := 0; i < 1000; i++ {
		cache.Set(fmt.Sprintf("key%d", i), i)
	}
	cache.Clear()
	if len(cache.storage) != 0 {
		t.Errorf("Test case 5 failed: expected empty cache, got %v", cache.storage)
	}
}
