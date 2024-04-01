package mutex

import (
	"kv-cache/lru"
	"sync"
)

type cache struct {
	mu  sync.Mutex
	lru *lru.Cache
}

func New(cacheBytes int64) *cache {
	return &cache{
		lru: lru.New(cacheBytes),
	}
}

type ByteView []byte

func (b ByteView) Len() int {
	return len(b)
}

func (b ByteView) String() string {
	return string(b)
}

func (b ByteView) ByteSlice() []byte {
	return b
}

func (c *cache) Set(key string, value []byte) {
	c.mu.Lock()
	defer c.mu.Unlock()
	c.lru.Set(key, ByteView(value))
}

func (c *cache) Get(key string) ([]byte, bool) {
	c.mu.Lock()
	defer c.mu.Unlock()
	if val, ok := c.lru.Get(key); ok {
		return val.(ByteView), ok
	}
	return nil, false
}
