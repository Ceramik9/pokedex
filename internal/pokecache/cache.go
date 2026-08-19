package pokecache

import (
	"sync"
	"time"
)

func NewCache(interval time.Duration) *Cache {
	
	//create new cache
	c := Cache{
		cacheEntries: make(map[string]cacheEntry),
	}
	// start reapLoop
	go c.reapLoop(interval)
	
	// return pointer the the new cache
	return &c
}

type Cache struct {
	mu         sync.RWMutex
	cacheEntries map[string]cacheEntry
}

func (c *Cache) Add(key string, val []byte) {
	
	// cache lock
	c.mu.Lock()
	defer c.mu.Unlock()

	// add new cache entry
	c.cacheEntries[key] = cacheEntry{time.Now(), val}
}

func (c *Cache) Get(key string) ([]byte, bool) {
	
	// cache write lock
	c.mu.RLock()
	defer c.mu.RUnlock()

	// if cache entry exist, return the entry value
	entry, ok := c.cacheEntries[key]
	if !ok {
		return nil, false
	}
	return entry.val, true
}

func (c *Cache) reapLoop(interval time.Duration) {
	
	// set a time interval for the cleanup
	ticker := time.NewTicker(interval)
	
	// for every tick
	for range ticker.C {
		// cache lock
		c.mu.Lock()
		
		// delete cache entries oldaer then the interval
		for key, cacheEntry := range c.cacheEntries {
			if time.Since(cacheEntry.createdAt) > interval {
				delete(c.cacheEntries, key)
			}
		}
		// cache unlock
		c.mu.Unlock()
	}
}

type cacheEntry struct {
	createdAt time.Time
	val       []byte
}

