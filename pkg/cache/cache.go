package cache

import (
	"net/http"
	"sync"
)

type Cache interface {
	Get(key string) (responseBytes []byte, ok bool)
	Set(key string, responseBytes []byte)
	CacheKey(r *http.Request) string
}
type MemoryCache struct {
	mu    sync.RWMutex
	items map[string][]byte
}

func NewMemoryCache() *MemoryCache {
	c := &MemoryCache{items: map[string][]byte{}}
	return c
}

func (c *MemoryCache) CacheKey(r *http.Request) string {
	return r.Method + "_" + r.URL.String()
}

func (c *MemoryCache) Get(key string) (resp []byte, ok bool) {
	c.mu.RLock()
	resp, ok = c.items[key]
	c.mu.RUnlock()
	return resp, ok
}

func (c *MemoryCache) Set(key string, resp []byte) {
	c.mu.Lock()
	c.items[key] = resp
	c.mu.Unlock()
}
