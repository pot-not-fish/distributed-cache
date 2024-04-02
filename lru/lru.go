package lru

import "container/list"

type Cache struct {
	maxBytes     int64
	currentBytes int64
	ll           *list.List
	cache        map[string]*list.Element
}

type CacheInterface interface {
	New(maxBytes int64)
	Set(key string, value []byte)
	Get(key string) (value []byte, ok bool)
}

type entry struct {
	key   string
	value []byte
}

func (c *Cache) New(maxBytes int64) {
	c.maxBytes = maxBytes
	c.ll = list.New()
	c.cache = make(map[string]*list.Element)
}

func (c *Cache) Get(key string) (value []byte, ok bool) {
	if val, ok := c.cache[key]; ok {
		c.ll.MoveToFront(val)
		kv := val.Value.(*entry)
		return kv.value, ok
	}
	return nil, false
}

func (c *Cache) Del() {
	if val := c.ll.Back(); val != nil {
		c.ll.Remove(val)
		kv := val.Value.(*entry)
		delete(c.cache, kv.key)
		c.currentBytes -= int64(len(kv.key) + len(kv.value))
	}
}

// 更新操作
func (c *Cache) Set(key string, value []byte) {
	if val, ok := c.cache[key]; ok {
		kv := val.Value.(*entry)
		c.currentBytes -= int64(len(kv.key) + len(kv.value))
		c.ll.Remove(val)
	}
	ele := c.ll.PushFront(&entry{key: key, value: value})
	kv := ele.Value.(*entry)
	c.currentBytes += int64(len(kv.key) + len(kv.value))
	c.cache[key] = c.ll.Front()
	for c.maxBytes != 0 && c.maxBytes < c.currentBytes {
		c.Del()
	}
}
