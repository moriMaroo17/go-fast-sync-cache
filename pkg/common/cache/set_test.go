package cache

import (
	"sync"
	"testing"
	"time"
)

func TestSetWithNilStorageLock(t *testing.T) {
	// Test case: Set an item with a nil storageLock
	// Expected behavior: The function should panic with a nil pointer dereference error

	// Create a new cache with a nil storageLock
	cache := &Cache[string, int]{
		storage:     make(map[string]item[int]),
		ttl:         10 * time.Second,
		storageLock: nil, // Set the storageLock to nil
	}

	// Run the test in a separate goroutine to catch the panic
	go func() {
		defer func() {
			if r := recover(); r == nil {
				t.Error("The function did not panic with a nil pointer dereference error")
			}
		}()

		// Attempt to set an item in the cache
		cache.Set("key", 100)
	}()

	// Wait for the test to complete
	time.Sleep(100 * time.Millisecond)
}

func TestSetWithValidStorageLock(t *testing.T) {
	// Test case: Set an item with a valid storageLock
	// Expected behavior: The item should be added to the cache successfully

	// Create a new cache with a valid storageLock
	//cache := New[string, int](1)
	cache := &Cache[string, int]{
		storage:     make(map[string]item[int]),
		ttl:         10 * time.Second,
		storageLock: &sync.RWMutex{}, // Set the storageLock to a valid RWMutex
	}

	// Set an item in the cache
	cache.Set("key", 100)

	// Verify that the item was added to the cache
	value, ok := cache.Get("key")
	if !ok {
		t.Error("The item was not added to the cache")
	}
	if value != 100 {
		t.Errorf("The value in the cache is incorrect. Expected: %d, Actual: %d", 100, value)
	}
}

func TestSetWithExpiredTTL(t *testing.T) {
	// Test case: Set an item with an expired TTL
	// Expected behavior: The item should not be added to the cache

	// Create a new cache with a valid storageLock and an expired TTL
	cache := &Cache[string, int]{
		storage:     make(map[string]item[int]),
		ttl:         0 * time.Second, // Set the TTL to 0 seconds
		storageLock: &sync.RWMutex{},
	}

	// Set an item in the cache
	cache.Set("key", 100)

	// Verify that the item was not added to the cache
	_, ok := cache.Get("key")
	if ok {
		t.Error("The item was added to the cache with an expired TTL")
	}
}

func TestSetWithConcurrentAccess(t *testing.T) {
	// Test case: Set an item while concurrently accessing the cache
	// Expected behavior: The function should not panic and the item should be added to the cache successfully

	// Create a new cache with a valid storageLock
	cache := &Cache[string, int]{
		storage:     make(map[string]item[int]),
		ttl:         10 * time.Second,
		storageLock: &sync.RWMutex{},
	}

	// Set an item in the cache while concurrently accessing it
	var wg sync.WaitGroup
	wg.Add(2)

	go func() {
		defer wg.Done()
		cache.Set("key", 100)
	}()

	go func() {
		defer wg.Done()
		_, _ = cache.Get("key")
	}()

	wg.Wait()

	// Verify that the item was added to the cache
	value, ok := cache.Get("key")
	if !ok {
		t.Error("The item was not added to the cache")
	}
	if value != 100 {
		t.Errorf("The value in the cache is incorrect. Expected: %d, Actual: %d", 100, value)
	}
}

func TestSetWithNilValue(t *testing.T) {
	// Test case: Set an item with a nil value
	// Expected behavior: The item should be added to the cache successfully with a nil value

	// Create a new cache with a valid storageLock
	cache := &Cache[string, *int]{
		storage:     make(map[string]item[*int]),
		ttl:         10 * time.Second,
		storageLock: &sync.RWMutex{},
	}

	// Set a nil item in the cache
	var nilValue *int
	cache.Set("key", nilValue)

	// Verify that the nil item was added to the cache
	value, ok := cache.Get("key")
	if !ok {
		t.Error("The item was not added to the cache")
	}
	if value != nil {
		t.Error("The value in the cache is not nil")
	}
}
