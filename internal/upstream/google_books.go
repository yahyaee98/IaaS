package upstream

import (
	"github.com/prometheus/client_golang/prometheus"
	"iaas/internal/data"
	"iaas/pkg/google_books"
)

type GoogleBooksUpstream struct {
	gb *google_books.GoogleBooks
	o  prometheus.Observer
}

func NewGoogleBooksUpstream(gb *google_books.GoogleBooks, o prometheus.Observer) *GoogleBooksUpstream {
	return &GoogleBooksUpstream{gb: gb, o: o}
}

func (g GoogleBooksUpstream) Search(search string) ([]*data.Item, error) {
	var err error
	var response *google_books.Response

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
