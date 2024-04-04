package pokecache

import (
	"sync"
	"time"
)

type cacheEntry struct {
	createdAt time.Time
	data      []byte
}

type Cache struct {
	mux     *sync.RWMutex
	entries map[string]cacheEntry
}

func Create(interval time.Duration) Cache {
	cache := Cache{}
	cache.mux = &sync.RWMutex{}
	cache.entries = make(map[string]cacheEntry)
	go cache.reapLoop(interval)
	return cache
}

func (c *Cache) Add(key string, val []byte) {
	c.mux.Lock()
	c.entries[key] = cacheEntry{
		time.Now(),
		val,
	}
	c.mux.Unlock()
}

func (c *Cache) Get(key string) ([]byte, bool) {
	c.mux.RLock()
	defer c.mux.RUnlock()
	entry, present := c.entries[key]
	if !present {
		return []byte{}, false
	}
	return entry.data, true
}

func (c *Cache) reapLoop(interval time.Duration) {
	ticker := time.NewTicker(interval)
	for {
		<-ticker.C
		c.mux.Lock()
		for name, data := range c.entries {
			if time.Since(data.createdAt) >= interval {
				delete(c.entries, name)
			}
		}
		c.mux.Unlock()
	}
}
