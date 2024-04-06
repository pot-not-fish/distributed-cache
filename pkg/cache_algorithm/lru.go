package cache_algorithm

import (
	"container/list"
	"fmt"
)

type LRUCache struct {
	maxBytes     int64
	currentBytes int64
	ll           *list.List
	cache        map[string]*list.Element
}

type CacheInterface interface {
	New(maxBytes int64)
	Set(key string, value []byte)
	Get(key string) (value []byte, ok bool)
	Del(key string) error
	GetAll() ([]string, [][]byte)
}

type lruEntry struct {
	key   string
	value []byte
}

func (c *LRUCache) New(maxBytes int64) {
	c.maxBytes = maxBytes
	c.ll = list.New()
	c.cache = make(map[string]*list.Element)
}

func (c *LRUCache) Get(key string) (value []byte, ok bool) {
	if val, ok := c.cache[key]; ok {
		c.ll.MoveToFront(val)
		kv := val.Value.(*lruEntry)
		return kv.value, ok
	}
	return nil, false
}

func (c *LRUCache) Del(key string) error {
	if val, ok := c.cache[key]; !ok {
		return fmt.Errorf("nil value")
	} else {
		c.ll.Remove(val)
		delete(c.cache, key)
		return nil
	}
}

func (c *LRUCache) DelBack() {
	if val := c.ll.Back(); val != nil {
		c.ll.Remove(val)
		kv := val.Value.(*lruEntry)
		delete(c.cache, kv.key)
		c.currentBytes -= int64(len(kv.key) + len(kv.value))
	}
}

func (c *LRUCache) Set(key string, value []byte) {
	if val, ok := c.cache[key]; ok {
		kv := val.Value.(*lruEntry)
		c.currentBytes -= int64(len(kv.key) + len(kv.value))
		c.ll.Remove(val)
	}
	ele := c.ll.PushFront(&lruEntry{key: key, value: value})
	kv := ele.Value.(*lruEntry)
	c.currentBytes += int64(len(kv.key) + len(kv.value))
	c.cache[key] = c.ll.Front()
	for c.maxBytes != 0 && c.maxBytes < c.currentBytes {
		c.DelBack()
	}
}

func (c *LRUCache) GetAll() ([]string, [][]byte) {
	keys := make([]string, 0, c.ll.Len())
	values := make([][]byte, 0, c.ll.Len())
	iter := c.ll.Front()
	for iter != nil {
		keys = append(keys, iter.Value.(*lruEntry).key)
		values = append(values, iter.Value.(*lruEntry).value)
		iter = iter.Next()
	}
	return keys, values
}
