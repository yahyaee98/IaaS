package googlebooks

// Response is the struct returned by the Google Books API.
type Response struct {
	Items []Item `json:"items"`
}

// Item is an entity returned by Google Books API.
type Item struct {
	VolumeInfo volumeInfo `json:"volumeInfo"`
}

type volumeInfo struct {
	Title   string   `json:"title"`
	Authors []string `json:"authors"`
}
