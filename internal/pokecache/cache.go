package pokecache

import (
	"sync"
	"time"
)

type cacheEntry struct {
	createdAt time.Time
	val       []byte
}

type Cache struct {
	cacheEntries map[string]cacheEntry
	mut          *sync.Mutex
}

func NewCache(interval time.Duration) Cache {
	c := Cache{
		cacheEntries: make(map[string]cacheEntry),
		mut:          &sync.Mutex{},
	}
	go c.reapLoop(interval)
	return c
}

func (c *Cache) Add(key string, val []byte) {
	c.mut.Lock()
	defer c.mut.Unlock()
	c.cacheEntries[key] = cacheEntry{time.Now(), val}
}

func (c *Cache) Get(key string) ([]byte, bool) {
	c.mut.Lock()
	defer c.mut.Unlock()
	value, exists := c.cacheEntries[key]
	return value.val, exists
}

func (c *Cache) reapLoop(interval time.Duration) {
	ticker := time.NewTicker(interval)
	for range ticker.C {
		c.reap(time.Now().UTC(), interval)
	}
}

func (c *Cache) reap(now time.Time, last time.Duration) {
	c.mut.Lock()
	defer c.mut.Unlock()
	for k, v := range c.cacheEntries {
		if v.createdAt.Before(now.Add(-last)) {
			delete(c.cacheEntries, k)
		}
	}
}

// func (c *Cache) reapLoop(interval time.Duration) {
// 	// this function will be called when the cache
// 	// is created by the NewCache function
// 	// each time an interval passes, remove any entries
// 	// from cacheEntries that are older than interval
// 	ticker := time.NewTicker(interval)
// 	// channel that sends current time on channel after
// 	// each tick; period b/w ticks specified by interval
// 	for range ticker.C {
// 		c.mut.Lock()
// 		defer c.mut.Unlock()
// 		for key, value := range c.cacheEntries {
// 			if value.createdAt.Before(<-time.Tick(interval)) {
// 				delete(c.cacheEntries, key)
// 			}
// 		}
// 		close(time.Tick(interval))
// 	}
// }
