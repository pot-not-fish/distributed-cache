package mutex

import (
	"kv-cache/lru"
	"sync"
)

type Cache struct {
	mu    sync.Mutex
	cache lru.CacheInterface
}

func New(cacheBytes int64, cache lru.CacheInterface) *Cache {
	if cache == nil {
		cache = &lru.Cache{}
	}
	cache.New(cacheBytes)
	return &Cache{
		cache: cache,
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

func (c *Cache) Set(key string, value []byte) {
	c.mu.Lock()
	defer c.mu.Unlock()

	c.cache.Set(key, ByteView(value))
}

func (c *Cache) Get(key string) ([]byte, bool) {
	c.mu.Lock()
	defer c.mu.Unlock()

	if val, ok := c.cache.Get(key); ok {
		return val, ok
	}
	return nil, false
}
