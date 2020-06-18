package cache

import (
	"context"
	"github.com/go-redis/redis/v8"
	"iaas/internal/log"
	"time"
)

type RedisCache struct {
	rdb *redis.Client
}

func NewRedisCache(addr, password string, db int) *RedisCache {
	rdb := redis.NewClient(&redis.Options{
		Addr:     addr,
		Password: password,
		DB:       db,
	})

	return &RedisCache{
		rdb: rdb,
	}
}

func (r RedisCache) Get(key string) (cached interface{}, found bool) {
	val, err := r.rdb.Get(context.Background(), key).Result()
	if err != nil {
		return nil, false
	}

	if len(val) < 1 {
		return nil, false
	}

	return val, true
}

func (r RedisCache) Set(key string, value interface{}, expiration time.Duration) {
	err := r.rdb.Set(context.Background(), key, value, expiration).Err()
	if err != nil {
		log.Errorw("redis cache has failed to set",
			"error", err,
		)
	}
}
