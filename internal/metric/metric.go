package metric

import (
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promauto"
)

var (
	buckets = []float64{
		0.01, // 10ms
		0.02,
		0.05,
		0.1, // 100 ms
		0.2,
		0.5,
		1.0, // 1s
		2.0,
		5.0,
		10.0,
	}

	// ItunesResponseTime is the histogram holding response times for iTunes.
	ItunesResponseTime = promauto.NewHistogram(prometheus.HistogramOpts{
		Name:    "iaas_external_itunes_responsetime",
		Help:    "iTunes response time",
		Buckets: buckets,
	})

	// GoogleBooksResponseTime is the histogram holding response times for Google Books.
	GoogleBooksResponseTime = promauto.NewHistogram(prometheus.HistogramOpts{
		Name:    "iaas_external_googlebooks_responsetime",
		Help:    "Google Books response time",
		Buckets: buckets,
	})
)
