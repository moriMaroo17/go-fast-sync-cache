package cache

import (
	"sync"
	"time"
)

// New creates a new instance of Cache with the specified time-to-live (TTL) and clean-up interval.
// The TTL is the duration after which an item will be automatically removed from the cache.
// The cleanTtl is the interval at which the cache will be cleaned up to remove expired items.
// If cleanTtl is set to 0 or less, the cache will not be cleaned up automatically.
// If ttl is set to 0 or less, items will never expire.
// The returned Cache instance is safe for concurrent use.
func New[K comparable, V any](ttl, cleanTtl time.Duration) *Cache[K, V] {
	c := Cache[K, V]{
		storage:     make(map[K]item[V]),
		storageLock: &sync.RWMutex{},
		ttl:         ttl,
	}
	if cleanTtl > 0 && ttl > 0 {
		c.cleanTicker = time.NewTicker(cleanTtl)
		go func() {
			for range c.cleanTicker.C {
				c.storageLock.Lock()
				for key, value := range c.storage {
					if time.Now().UnixNano() > value.expire {
						delete(c.storage, key)
					}
				}
				c.storageLock.Unlock()
			}
		}()
	}

	return &c
}
