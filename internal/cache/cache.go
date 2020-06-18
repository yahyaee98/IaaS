package cache

import "time"

type Cache interface {
	Get(key string) (cached interface{}, found bool)
	Set(key string, value interface{}, expiration time.Duration)
}
