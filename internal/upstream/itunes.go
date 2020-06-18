package upstream

import (
	"github.com/prometheus/client_golang/prometheus"
	"iaas/internal/data"
	"iaas/pkg/itunes"
)

// ItunesUpstream is the adapter to Itunes.
type ItunesUpstream struct {
	itunes *itunes.Itunes
	o      prometheus.Observer
}

// NewItunesUpstream returns a new ItunesUpstream.
func NewItunesUpstream(itunes *itunes.Itunes, o prometheus.Observer) *ItunesUpstream {
	return &ItunesUpstream{itunes: itunes, o: o}
}

// Search returns data fetched from the Itunes client and also cares about metrics.
func (i ItunesUpstream) Search(search string) ([]*data.Item, error) {
	var err error
	var response *itunes.Response

	reportDuration(
		func() {
			response, err = i.itunes.Search(search)
		},
		i.o,
	)

	if err != nil {
		return nil, err
	}

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
