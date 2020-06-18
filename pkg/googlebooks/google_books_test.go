package googlebooks

import (
	"github.com/stretchr/testify/assert"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestGoogleBooksCallsRightUrl(t *testing.T) {
	server := httptest.NewServer(
		http.HandlerFunc(
			func(rw http.ResponseWriter, req *http.Request) {
				assert.Equal(
					t,
					"/?q=intitle:some_search&key=the_key",
					req.URL.String(),
				)
				_, _ = rw.Write([]byte(`{}`))
			},
		),
	)
	defer server.Close()

	gb := &GoogleBooks{
		resultLimit: 1,
		baseURL:     server.URL,
		apiKey:      "the_key",
		client:      server.Client(),
		log: func(msg string, keysAndValues ...interface{}) {

		},
	}

	_, _ = gb.Search("some_search")
}
