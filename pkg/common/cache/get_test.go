package cache

import (
	"testing"
	"time"
)

func TestGet(t *testing.T) {
	// Test case 1: Get an existing key-value pair
	cache := New[string, int](6*time.Second, time.Minute)
	cache.Set("key1", 100)
	value, ok := cache.Get("key1")
	if !ok || value != 100 {
		t.Errorf("Expected (100, true), got (%v, %v)", value, ok)
	}

	// Test case 2: Get a non-existent key
	value, ok = cache.Get("key2")
	if ok || value != 0 {
		t.Errorf("Expected (0, false), got (%v, %v)", value, ok)
	}

	// Test case 3: Get an expired key
	cache.Set("key3", 200)
	time.Sleep(6 * time.Second) // Wait for the key to expire
	value, ok = cache.Get("key3")
	if ok || value != 0 {
		t.Errorf("Expected (0, false), got (%v, %v)", value, ok)
	}

	// Test case 4: Get a key after deleting it
	cache.Set("key4", 300)
	cache.Delete("key4")
	value, ok = cache.Get("key4")
	if ok || value != 0 {
		t.Errorf("Expected (0, false), got (%v, %v)", value, ok)
	}

	// Test case 5: Get a key after clearing the cache
	cache.Set("key5", 400)
	cache.Clear()
	value, ok = cache.Get("key5")
	if ok || value != 0 {
		t.Errorf("Expected (0, false), got (%v, %v)", value, ok)
	}
}
