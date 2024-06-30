package cache

import (
	"time"
)

// Set adds a key-value pair to the cache with an expiration time.
// If the key already exists, it will be overwritten.
//
// Parameters:
// - key: The unique identifier for the value.
// - value: The value to be stored in the cache.
func (c *Cache[K, V]) Set(key K, value V) {
	c.storageLock.Lock()
	defer c.storageLock.Unlock()
	c.storage[key] = item[V]{
		value:  value,
		expire: time.Now().Add(c.ttl).UnixNano(),
	}
}

// Get retrieves the value associated with the given key from the cache.
// If the key does not exist or has expired, it returns false for the 'ok' parameter.
//
// Parameters:
// - key: The unique identifier for the value.
//
// Return values:
// - value: The value associated with the given key.
// - ok: A boolean indicating whether the key exists and has not expired.
func (c *Cache[K, V]) Get(key K) (value V, ok bool) {
	c.storageLock.RLock()
	defer c.storageLock.RUnlock()
	storedItem, ok := c.storage[key]
	if !ok || time.Now().UnixNano() > storedItem.expire {
		return value, false
	}
	value = storedItem.value
	return
}

// Delete removes the key-value pair associated with the given key from the cache.
// If the key does not exist, the function does nothing.
//
// Parameters:
// - key: The unique identifier for the value to be deleted.
func (c *Cache[K, V]) Delete(key K) {
	c.storageLock.Lock()
	defer c.storageLock.Unlock()
	delete(c.storage, key)
}

// Clear removes all key-value pairs from the cache.
// This function is thread-safe and does not return any value.
//
// The function locks the storage mutex before performing the operation to ensure that
// concurrent access to the cache does not lead to data corruption.
// After clearing the cache, it unlocks the mutex to allow other goroutines to access the cache.
//
// The function uses the 'make' function to create a new, empty map for the 'storage' field.
// This effectively clears all existing key-value pairs from the cache.
func (c *Cache[K, V]) Clear() {
	c.storageLock.Lock()
	defer c.storageLock.Unlock()
	c.storage = make(map[K]item[V])
}
