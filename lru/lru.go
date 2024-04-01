package lru

import "container/list"

type Cache struct {
	maxBytes     int64
	currentBytes int64
	ll           *list.List
	cache        map[string]*list.Element
}

type entry struct {
	key   string
	value Value
}

type Value interface {
	Len() int
}

func New(maxBytes int64) *Cache {
	return &Cache{
		maxBytes: maxBytes,
		ll:       list.New(),
		cache:    make(map[string]*list.Element),
	}
}

func (c *Cache) Get(key string) (value Value, ok bool) {
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
		c.currentBytes -= int64(len(kv.key) + kv.value.Len())
	}
}

// 更新操作
func (c *Cache) Set(key string, value Value) {
	if val, ok := c.cache[key]; ok {
		kv := val.Value.(*entry)
		c.currentBytes -= int64(len(kv.key) + kv.value.Len())
		c.ll.Remove(val)
	}
	ele := c.ll.PushFront(&entry{key: key, value: value})
	kv := ele.Value.(*entry)
	c.currentBytes += int64(len(kv.key) + kv.value.Len())
	c.cache[key] = c.ll.Front()
	for c.maxBytes != 0 && c.maxBytes < c.currentBytes {
		c.Del()
	}
}
