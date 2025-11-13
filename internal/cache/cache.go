package cache

import (
	"sync"
)

type Cache interface {
	Get(key string) (interface{}, bool)
	Set(key string, value interface{})
}

type inMemoryCache struct {
	mu    sync.RWMutex
	store map[string]interface{}
}

func NewCache() Cache {
	return &inMemoryCache{
		store: make(map[string]interface{}),
	}
}

func (i *inMemoryCache) Get(key string) (interface{}, bool) {
	i.mu.RLock()
	defer i.mu.RUnlock()
	val, ok := i.store[key]
	return val, ok
}

func (i *inMemoryCache) Set(key string, value interface{}) {
	i.mu.Lock()
	defer i.mu.Unlock()
	i.store[key] = value
}
