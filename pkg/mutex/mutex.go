package mutex

import (
	"fmt"
	"kv-cache/pkg/cache_algorithm"
	"sync"
)

type Cache struct {
	mu    sync.Mutex
	cache cache_algorithm.CacheInterface
}

func New(cacheBytes int64, cache cache_algorithm.CacheInterface) *Cache {
	if cache == nil {
		cache = &cache_algorithm.LRUCache{}
	}
	cache.New(cacheBytes)
	return &Cache{
		cache: cache,
	}
}

func (c *Cache) Del(key string) error {
	c.mu.Lock()
	defer c.mu.Unlock()

	return c.cache.Del(key)
}

func (c *Cache) Set(key string, value []byte) {
	c.mu.Lock()
	defer c.mu.Unlock()

	c.cache.Set(key, value)
}

func (c *Cache) Get(key string) ([]byte, bool) {
	c.mu.Lock()
	defer c.mu.Unlock()

	if val, ok := c.cache.Get(key); ok {
		ret := make([]byte, len(val))
		copy(ret, val)
		fmt.Println(string(ret))
		fmt.Println(string(val))
		return ret, ok
	}
	return nil, false
}

func (c *Cache) GetAll() ([]string, [][]byte) {
	return c.cache.GetAll()
}
