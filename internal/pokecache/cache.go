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
	return c
}

func (c *Cache) Add(key string, val []byte) {
	c.cacheEntries[key] = cacheEntry{time.Now(), val}
}

func (c *Cache) Get(key string) ([]byte, bool) {
	value, exists := c.cacheEntries[key]
	return value.val, exists
}
