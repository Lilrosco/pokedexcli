package pokecache

import (
	"sync"
	"time"
)

type Cache struct {
	entries map[string]CacheEntry
	mu sync.Mutex
}

type CacheEntry struct {
	createdAt time.Time
	val []byte
}

func New(interval time.Duration) *Cache {
	cache := Cache{entries: make(map[string]CacheEntry)}

	// Delete entries in cache older than passed interval in a subroutine
	go cache.reapLoop(interval)

	return &cache
}

func (cache *Cache) Add(key string, val []byte) {
	cache.mu.Lock()
	defer cache.mu.Unlock()

	cache.entries[key] = CacheEntry{createdAt: time.Now(), val: val}
}

func (cache *Cache) Get(key string) ([]byte, bool) {
	cache.mu.Lock()
	defer cache.mu.Unlock()

	entry, found := cache.entries[key]
	return entry.val, found
}

func (cache *Cache) reapLoop(interval time.Duration) {
	ticker := time.NewTicker(interval)
	defer ticker.Stop()

	for {
		select {
		case t := <-ticker.C:
			cache.mu.Lock()

			for key, entry := range cache.entries {
				if t.After(entry.createdAt) {
					delete(cache.entries, key)
				}
			}

			cache.mu.Unlock()
		}
	}
}
