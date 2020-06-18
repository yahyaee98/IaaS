package repository

import (
	"iaas/internal/data"
	"time"
)

type mockUpstream struct {
	searchFunc func(search string) ([]*data.Item, error)
}

func (m mockUpstream) Search(search string) ([]*data.Item, error) {
	return m.searchFunc(search)
}

type mockCache struct {
	getFunc func(key string) (cached interface{}, found bool)
	setFunc func(key string, value interface{}, expiration time.Duration)
}

func (m mockCache) Get(key string) (cached interface{}, found bool) {
	return m.getFunc(key)
}

func (m mockCache) Set(key string, value interface{}, expiration time.Duration) {
	m.setFunc(key, value, expiration)
}
