package upstream

import (
	"iaas/internal/data"
	"iaas/pkg/itunes"
	"time"
)

type ItunesUpstream struct {
	itunes *itunes.Itunes
	mr     MetricReport
}

func NewItunesUpstream(itunes *itunes.Itunes, mr MetricReport) *ItunesUpstream {
	return &ItunesUpstream{itunes: itunes, mr: mr}
}

func (i ItunesUpstream) Search(search string) ([]*data.Item, error) {
	response, err := i.itunes.Search(search)
	if err != nil {
		return nil, err
	}
	i.mr(time.Second) //TODO implement

	var items []*data.Item

	for _, item := range response.Results {
		items = append(items, &data.Item{
			Title:    item.TrackName,
			Type:     "music",
			Creators: []string{item.ArtistName},
		})
	}

	return items, nil
}
