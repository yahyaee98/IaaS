package upstream

import (
	"github.com/prometheus/client_golang/prometheus"
)

func reportDuration(f func(), obs prometheus.Observer) {
	t := prometheus.NewTimer(obs)
	f()
	t.ObserveDuration()
}
