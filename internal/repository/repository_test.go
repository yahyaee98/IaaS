package repository

import (
	"github.com/stretchr/testify/assert"
	"iaas/internal/data"
	"iaas/internal/upstream"
	"testing"
	"time"
)

func TestGetItemsDoesNotFetchDataFromUpstreamsWhenCachedResultIsAvailable(t *testing.T) {
	mockedCache := &mockCache{
		getFunc: func(key string) (cached interface{}, found bool) {
			return "something_cached", true
		},
		setFunc: func(key string, value interface{}, expiration time.Duration) {

		},
	}

	upstreamUsed := false
	mockedUpstream1 := &mockUpstream{
		searchFunc: func(search string) ([]*data.Item, error) {
			upstreamUsed = true
			return nil, nil
		},
	}

	r := &repository{
		upstreams: []upstream.Upstream{
			mockedUpstream1,
		},
		cache: mockedCache,
	}

	_, _ = r.GetItems("some_search")

	assert.False(t, upstreamUsed)
}

func TestGetItemsGathersDataFromAllUpstreamsIfNoCacheIsAvailable(t *testing.T) {
	mockedCache := &mockCache{
		getFunc: func(key string) (cached interface{}, found bool) {
			return nil, false
		},
		setFunc: func(key string, value interface{}, expiration time.Duration) {

		},
	}

	firstUpstreamUsed := false
	mockedUpstream1 := &mockUpstream{
		searchFunc: func(search string) ([]*data.Item, error) {
			firstUpstreamUsed = true
			return nil, nil
		},
	}

	secondUpstreamUsed := false
	mockedUpstream2 := &mockUpstream{
		searchFunc: func(search string) ([]*data.Item, error) {
			secondUpstreamUsed = true
			return nil, nil
		},
	}

	r := &repository{
		upstreams: []upstream.Upstream{
			mockedUpstream1,
			mockedUpstream2,
		},
		cache: mockedCache,
	}

	_, _ = r.GetItems("some_search")

	assert.True(t, firstUpstreamUsed)
	assert.True(t, secondUpstreamUsed)
}
