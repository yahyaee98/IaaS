package google_books

type Response struct {
	Items []Item `json:"items"`
}

type Item struct {
	VolumeInfo volumeInfo `json:"volumeInfo"`
}

type volumeInfo struct {
	Title   string   `json:"title"`
	Authors []string `json:"authors"`
}
