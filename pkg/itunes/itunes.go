package itunes

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"time"
)

// Itunes is a separate struct which can be used in other projects as it's not coupled to our code.
// We could even publish this as a library.
type Itunes struct {
	resultLimit int
	baseUrl     string
	client      *http.Client
	log         func(msg string, keysAndValues ...interface{})
}

func NewItunes(resultLimit int, baseUrl string, timeout time.Duration, log func(msg string, keysAndValues ...interface{})) *Itunes {
	client := http.Client{
		Timeout: timeout,
	}

	return &Itunes{
		resultLimit: resultLimit,
		baseUrl:     baseUrl,
		client:      &client,
		log:         log,
	}
}

func (i Itunes) Search(search string) (*Response, error) {
	url := fmt.Sprintf("%s?term=%s&limit=%d&entity=musicTrack", i.baseUrl, search, i.resultLimit)
	resp, err := i.client.Get(url)

	defer func() {
		if resp != nil {
			if err := resp.Body.Close(); err != nil {
				i.log("Itunes has failed to close Response body",
					"error", err,
				)
			}
		}
	}()

	if err != nil {
		_, _ = ioutil.ReadAll(resp.Body)
		return nil, err
	}

	var response Response
	if err := json.NewDecoder(resp.Body).Decode(&response); err != nil {
		return nil, err
	}

	return &response, nil
}
