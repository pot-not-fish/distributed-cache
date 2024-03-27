package lru

import "container/list"

type Cache struct {
	maxBytes     int64
	currentBytes int64
	ll           *list.List
	cache        map[string]*list.Element
	OnEvicted    func(key string, value Value)
}

type entry struct {
	key   string
	value Value
}

type Value interface {
	Len() int
}

func New(maxBytes int64, onEvicted func(string, Value)) *Cache {
	return &Cache{
		maxBytes:  maxBytes,
		ll:        list.New(),
		cache:     make(map[string]*list.Element),
		OnEvicted: onEvicted,
	}
}

func (c *Cache) Get(key string) (value Value, ok bool) {
	if ele, ok := c.cache[key]; ok {
		c.ll.MoveToFront(ele)
		kv := ele.Value.(*entry)
		return kv.value, ok
	}
	return nil, false
}

func (c *Cache) Del() (value Value, ok bool) {
	if ele := c.ll.Back(); ele != nil {
		c.ll.Remove(ele)
		kv := ele.Value.(*entry)
		delete(c.cache, kv.key)
		c.currentBytes -= int64(len(kv.key) + kv.value.Len())
		return kv.value, true
	}
	return nil, false
}

func (c *Cache) Set(key string, value Value) {
	if _, ok := c.Get(key); !ok {
		ele := c.ll.PushFront(&entry{key: key, value: value})
		kv := ele.Value.(*entry)
		c.currentBytes += int64(len(kv.key) + kv.value.Len())
	}
	c.cache[key] = c.ll.Front()
	for c.maxBytes != 0 && c.maxBytes < c.currentBytes {
		c.Del()
	}
}
