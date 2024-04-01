package pokecache

import (
	"sync"
	"time"
)

type CacheEntry struct {
	createdAt time.Time
	val []byte
}

type Cache struct {
	entries map[string]CacheEntry
	mu *sync.Mutex
}

func NewCache(interval time.Duration) *Cache {
	cache := &Cache{}
	cache.mu = &sync.Mutex{}
	cache.entries = make(map[string]CacheEntry)
	go cache.reapLoop(interval)
	return cache
}

func (cache *Cache) Add(key string, val []byte) {
	cache.mu.Lock()
	defer cache.mu.Unlock()
	var entry CacheEntry
	entry.createdAt = time.Now().UTC()
	entry.val = val
	cache.entries[key] = entry

}

func (cache *Cache) Get(key string) ([]byte, bool) {
	cache.mu.Lock()
	defer cache.mu.Unlock()
	entry, ok := cache.entries[key]
	return entry.val, ok

}

func (cache *Cache) reapLoop(interval time.Duration) {
	
	ticker := time.NewTicker(interval)
	for range ticker.C {
		cache.mu.Lock()
		for key, val := range cache.entries {
			if val.createdAt.Before(time.Now().UTC().Add(-interval)) {
				delete(cache.entries, key)
			}
		}
	}
	defer cache.mu.Unlock()

}