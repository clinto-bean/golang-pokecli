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
	interval time.Duration
}

func NewCache(interval time.Duration) *Cache {
	c := &Cache{}
	c.mu = &sync.Mutex{}
	c.entries = make(map[string]CacheEntry)
	c.interval = interval
	go c.reapLoop()
	return c
}

func (cache *Cache) Add(key string, val []byte) {
	cache.mu.Lock()
	defer cache.mu.Unlock()
	var entry CacheEntry
	entry.createdAt = time.Now()
	entry.val = val
	cache.entries[key] = entry

}

func (cache *Cache) Get(key string) ([]byte, bool) {
	cache.mu.Lock()
	defer cache.mu.Unlock()
	entry, ok := cache.entries[key]
	if !ok {
		return nil, false
	}
	return entry.val, true

}

func (cache *Cache) reapLoop() {
	ticker := time.NewTicker(cache.interval)
	
	go func() {
		defer ticker.Stop()
		
		for {
		
		<- ticker.C

		curr := time.Now()
		
		var discard []string

		for key, entry := range cache.entries {
			expired := curr.After(entry.createdAt.Add(cache.interval))
			if expired {
				discard = append(discard, key)
			}
		}

		cache.mu.Lock()
		for _, key := range discard {
			delete(cache.entries, key)
		}
		cache.mu.Unlock()}
	}()
}