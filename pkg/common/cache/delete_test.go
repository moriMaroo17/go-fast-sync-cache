package cache

import (
	"sync"
	"testing"
	"time"
)

func TestDelete(t *testing.T) {
	// Test case 1: Delete an existing key
	cache := New[string, int](10*time.Second, time.Minute)
	cache.Set("key1", 100)
	cache.Delete("key1")
	_, ok := cache.Get("key1")
	if ok {
		t.Errorf("Test case 1 failed: expected key to be deleted, but it still exists")
	}

	// Test case 2: Delete a non-existent key
	cache = New[string, int](10*time.Second, time.Minute)
	cache.Delete("key2")
	_, ok = cache.Get("key2")
	if ok {
		t.Errorf("Test case 2 failed: expected key to be deleted, but it still exists")
	}

	// Test case 3: Delete a key after it expires
	cache = New[string, int](1*time.Second, time.Minute)
	cache.Set("key3", 200)
	time.Sleep(2 * time.Second)
	cache.Delete("key3")
	_, ok = cache.Get("key3")
	if ok {
		t.Errorf("Test case 3 failed: expected key to be deleted, but it still exists")
	}

	// Test case 4: Delete a key while multiple goroutines are accessing the cache
	cache = New[string, int](10*time.Second, time.Minute)
	cache.Set("key4", 300)
	var wg sync.WaitGroup
	wg.Add(2)
	go func() {
		defer wg.Done()
		cache.Delete("key4")
	}()
	go func() {
		defer wg.Done()
		time.Sleep(1 * time.Second) // Wait for the goroutine delete "key4"
		_, ok := cache.Get("key4")
		if ok {
			t.Errorf("Test case 4 failed: expected key to be deleted, but it still exists")
		}
	}()
	wg.Wait()

	// Test case 5: Delete a key multiple times
	cache = New[string, int](10*time.Second, time.Minute)
	cache.Set("key5", 400)
	cache.Delete("key5")
	cache.Delete("key5")
	_, ok = cache.Get("key5")
	if ok {
		t.Errorf("Test case 5 failed: expected key to be deleted, but it still exists")
	}
}
