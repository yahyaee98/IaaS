package repository

import (
	"crypto/md5"
	"encoding/json"
	"fmt"
	"iaas/internal/cache"
	"iaas/internal/data"
	"iaas/internal/upstream"
	"sort"
	"strings"
	"sync"
	"time"
)

// Repository is responsible for fetching items from upstreams/cache
type Repository interface {
	GetItems(search string) ([]*data.Item, error)
}

type repository struct {
	upstreams []upstream.Upstream
	cache     cache.Cache
}

// NewRepository returns an implementation of the Repository interface
func NewRepository(upstreams []upstream.Upstream, cache cache.Cache) Repository {
	return &repository{
		upstreams: upstreams,
		cache:     cache,
	}
}

func (r repository) GetItems(search string) ([]*data.Item, error) {
	// I initialize "items" this way to have an empty array instead of null when we json marshal it.
	items := make([]*data.Item, 0)

	cached, found := r.cache.Get(r.getKeyForCache(search))
	if found {
		_ = json.Unmarshal([]byte(cached.(string)), &items)
		return items, nil
	}

	items = r.FetchFromThirdParties(search)

	sort.Slice(items, func(i, j int) bool {
		return items[i].Title < items[j].Title
	})

	jsonData, err := json.Marshal(items)
	if err == nil {
		r.cache.Set(r.getKeyForCache(search), jsonData, time.Hour)
	}

	return items, nil
}

func (r repository) FetchFromThirdParties(search string) []*data.Item {
	// I initialize "allItems" this way to have an empty array instead of null when we json marshal it.
	allItems := make([]*data.Item, 0)

	l := &sync.Mutex{}
	wg := &sync.WaitGroup{}

	for _, u := range r.upstreams {
		wg.Add(1)
		go func(u upstream.Upstream) {
			defer wg.Done()

			items, err := u.Search(search)
			if err != nil {
				return
			}

			l.Lock()
			defer l.Unlock()
			allItems = append(allItems, items...)
		}(u)
	}

	// Wait till all items are fetched and also appended together.
	wg.Wait()

	return allItems
}

func (r repository) getKeyForCache(search string) string {
	return fmt.Sprintf(
		"%x",
		md5.Sum(
			[]byte(
				strings.ToLower(strings.TrimSpace(search)),
			),
		),
	)
}
