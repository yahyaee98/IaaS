package upstream

import (
	"github.com/prometheus/client_golang/prometheus"
	"iaas/internal/data"
	"iaas/pkg/googlebooks"
)

// GoogleBooksUpstream is the adapter to GoogleBooks.
type GoogleBooksUpstream struct {
	gb *googlebooks.GoogleBooks
	o  prometheus.Observer
}

// NewGoogleBooksUpstream returns a new GoogleBooksUpstream instance.
func NewGoogleBooksUpstream(gb *googlebooks.GoogleBooks, o prometheus.Observer) *GoogleBooksUpstream {
	return &GoogleBooksUpstream{gb: gb, o: o}
}

// Search returns data fetched from the GoogleBooks client and also cares about metrics.
func (g GoogleBooksUpstream) Search(search string) ([]*data.Item, error) {
	var err error
	var response *googlebooks.Response

	reportDuration(
		func() {
			response, err = g.gb.Search(search)
		},
		g.o,
	)

	if err != nil {
		return nil, err
	}

	var items []*data.Item

	for _, item := range response.Items {
		items = append(items, &data.Item{
			Title:    item.VolumeInfo.Title,
			Type:     "book",
			Creators: item.VolumeInfo.Authors,
		})
	}

	return items, nil
}
