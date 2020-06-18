package googlebooks

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"time"
)

// GoogleBooks is a separate struct which can be used in other projects as it's not coupled to our code.
// We could even publish this as a library.
type GoogleBooks struct {
	resultLimit int
	baseURL     string
	apiKey      string
	client      *http.Client
	log         func(msg string, keysAndValues ...interface{})
}

// NewGoogleBooks returns a new GoogleBooks instance.
func NewGoogleBooks(resultLimit int, baseURL, apiKey string, timeout time.Duration, log func(msg string, keysAndValues ...interface{})) *GoogleBooks {
	client := http.Client{
		Timeout: timeout,
	}

	return &GoogleBooks{
		resultLimit: resultLimit,
		baseURL:     baseURL,
		apiKey:      apiKey,
		client:      &client,
		log:         log,
	}
}

// Search fetches response from API.
func (gb GoogleBooks) Search(search string) (*Response, error) {
	url := fmt.Sprintf("%s?q=intitle:%s&key=%s", gb.baseURL, search, gb.apiKey)
	resp, err := gb.client.Get(url)

	defer func() {
		if resp != nil {
			if err := resp.Body.Close(); err != nil {
				gb.log("google books has failed to close Response body",
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

	if len(response.Items) > gb.resultLimit {
		response.Items = response.Items[:gb.resultLimit]
	}

	return &response, nil
}
