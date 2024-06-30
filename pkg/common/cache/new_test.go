package cache

import (
	"testing"
	"time"
)

func TestNewCacheNilTTL(t *testing.T) {
	cache := New[string, int](0, time.Second)
	if cache.ttl != 0 {
		t.Errorf("Expected TTL to be 0, got %v", cache.ttl)
	}
	if cache.cleanTicker != nil {
		t.Error("Expected cleanTicker to be nil, got a ticker")
	}
}

func TestNewCacheNilCleanTtl(t *testing.T) {
	cache := New[string, int](time.Second, 0)
	if cache.ttl != time.Second {
		t.Errorf("Expected TTL to be %v, got %v", time.Second, cache.ttl)
	}
	if cache.cleanTicker != nil {
		t.Error("Expected cleanTicker to be nil, got a ticker")
	}
}

func TestNewCacheNilTTLAndCleanTtl(t *testing.T) {
	cache := New[string, int](0, 0)
	if cache.ttl != 0 {
		t.Errorf("Expected TTL to be 0, got %v", cache.ttl)
	}
	if cache.cleanTicker != nil {
		t.Error("Expected cleanTicker to be nil, got a ticker")
	}
}

func TestNewCacheWithTTL(t *testing.T) {
	cache := New[string, int](time.Second, 2*time.Second)
	if cache.ttl != time.Second {
		t.Errorf("Expected TTL to be %v, got %v", time.Second, cache.ttl)
	}
	if cache.cleanTicker == nil {
		t.Error("Expected cleanTicker to be a ticker, got nil")
	}
}

func TestNewCacheWithCleanTtl(t *testing.T) {
	cache := New[string, int](2*time.Second, time.Second)
	if cache.ttl != 2*time.Second {
		t.Errorf("Expected TTL to be %v, got %v", 2*time.Second, cache.ttl)
	}
	if cache.cleanTicker == nil {
		t.Error("Expected cleanTicker to be a ticker, got nil")
	}
}
