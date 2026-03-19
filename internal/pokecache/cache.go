package pokecache

import (
	"sync"
	"time"
)

type CacheEntry struct {
	createdAt time.Time
	val       []byte
}

type Cache struct {
	cache map[string]CacheEntry
	mu    *sync.Mutex
}

func NewCache(interval time.Duration) Cache {
	cache := Cache{
		cache: make(map[string]CacheEntry),
		mu:    &sync.Mutex{},
	}
	go cache.reapLoop(interval)

	return cache
}

func (c *Cache) Add(key string, val []byte) {
	c.mu.Lock()
	defer c.mu.Unlock()
	c.cache[key] = CacheEntry{
		createdAt: time.Now(),
		val:       val,
	}
}

func (c *Cache) Get(key string) ([]byte, bool) {
	c.mu.Lock()
	defer c.mu.Unlock()
	val, ok := c.cache[key]
	if !ok {
		return nil, ok
	}

	return val.val, ok
}

func (c *Cache) reapLoop(interval time.Duration) {
	ticker := time.NewTicker(interval)
	for range ticker.C {
		c.reap(time.Now().UTC(), interval)
	}
}

func (c *Cache) reap(now time.Time, last time.Duration) {
	c.mu.Lock()
	defer c.mu.Unlock()
	for key, val := range c.cache {
		if val.createdAt.Before(now.Add(-last)) {
			delete(c.cache, key)
		}
	}
}
