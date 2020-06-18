package itunes

import (
	"github.com/stretchr/testify/assert"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestItunesCallsRightUrl(t *testing.T) {
	server := httptest.NewServer(
		http.HandlerFunc(
			func(rw http.ResponseWriter, req *http.Request) {
				assert.Equal(
					t,
					"/?term=some_search&limit=1&entity=musicTrack",
					req.URL.String(),
				)
				_, _ = rw.Write([]byte(`{}`))
			},
		),
	)
	defer server.Close()

	gb := &Itunes{
		resultLimit: 1,
		baseURL:     server.URL,
		client:      server.Client(),
		log: func(msg string, keysAndValues ...interface{}) {

		},
	}

	_, _ = gb.Search("some_search")
}
