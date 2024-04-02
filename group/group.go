package group

import (
	"fmt"
	"kv-cache/lru"
	"kv-cache/mutex"
	"kv-cache/singleflight"
	"sync"
)

type Group struct {
	name       string
	mainCache  *mutex.Cache
	load       *singleflight.Group
	getterFunc GetterFunc
}

type GetterFunc func(key string) ([]byte, bool)

var (
	mu     sync.RWMutex
	groups = make(map[string]*Group)
)

func NewGroup(name string, cacheBytes int64, getterFunc GetterFunc, cache lru.CacheInterface) (*Group, error) {
	mu.Lock()
	defer mu.Unlock()

	g := &Group{
		name:       name,
		mainCache:  mutex.New(cacheBytes, cache),
		getterFunc: getterFunc,
		load:       &singleflight.Group{},
	}
	if _, ok := groups[name]; ok {
		return nil, fmt.Errorf("repeat new group")
	}
	return g, nil
}

func GetGroup(name string) (*Group, error) {
	mu.RLock()
	defer mu.RUnlock()

	if val, ok := groups[name]; ok {
		return val, nil
	}
	return nil, fmt.Errorf("no such group")
}

func (g *Group) Get(key string) ([]byte, bool) {
	value, err := g.load.Do(key, func(key string) ([]byte, error) {
		val, ok := g.mainCache.Get(key)
		if !ok {
			return nil, fmt.Errorf("no such key")
		}
		return val, nil
	})
	if err != nil {
		if g.getterFunc != nil {
			return g.getterFunc(key)
		}
		return nil, false
	}
	return value, true
}

func (g *Group) Set(key string, value []byte) {
	g.mainCache.Set(key, value)
}
