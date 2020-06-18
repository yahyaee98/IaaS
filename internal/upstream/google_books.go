package upstream

import (
	"iaas/internal/data"
	"iaas/pkg/google_books"
	"time"
)

type GoogleBooksUpstream struct {
	gb *google_books.GoogleBooks
	mr MetricReport
}

func NewGoogleBooksUpstream(gb *google_books.GoogleBooks, mr MetricReport) *GoogleBooksUpstream {
	return &GoogleBooksUpstream{gb: gb, mr: mr}
}

func (g GoogleBooksUpstream) Search(search string) ([]*data.Item, error) {
	response, err := g.gb.Search(search)
	if err != nil {
		return nil, err
	}
	g.mr(time.Second) //TODO implement

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
