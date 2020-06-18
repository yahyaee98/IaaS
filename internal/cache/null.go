package cache

import "time"

// NullCache is useful when we decide to not use caches.
type NullCache struct {
}

func NewNullCache() *NullCache {
	return &NullCache{}
}

func (n NullCache) Get(_ string) (cached interface{}, found bool) {
	return nil, false
}

func (n NullCache) Set(_ string, _ interface{}, _ time.Duration) {
	return
}
