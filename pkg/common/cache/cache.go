package cache

import (
	"sync"
	"time"
)

// Cache is a generic type that represents a key-value cache with support for expiration.
// It is safe for concurrent use by multiple goroutines.
//
// The type parameters are:
// - K: the type of keys in the cache. It must be comparable.
// - V: the type of values in the cache.
type Cache[K comparable, V any] struct {
	// storage is a map that stores the key-value pairs in the cache.
	storage map[K]item[V]

	// storageLock is a read-write mutex to synchronize access to the storage map.
	storageLock *sync.RWMutex

	// ttl is the time-to-live duration for cache entries. Entries will be automatically removed after this duration.
	ttl time.Duration

	// cleanTicker is a ticker that triggers cleaning up expired entries in the cache.
	cleanTicker *time.Ticker
}

// item represents a single entry in the cache.
type item[V any] struct {
	// value is the value associated with the cache entry.
	value V

	// expire is the timestamp when the cache entry will expire.
	expire int64
}
