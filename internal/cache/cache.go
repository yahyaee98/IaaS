package cache

import "time"

// Cache is a simple interface responsible for getting/setting keys into/from the cache service
type Cache interface {
	Get(key string) (cached interface{}, found bool)
	Set(key string, value interface{}, expiration time.Duration)
}
