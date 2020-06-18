package itunes

// Response holds the data received from itunes
type Response struct {
	Results []Result `json:"results"`
}

// Result holds a single track
type Result struct {
	TrackName  string `json:"trackName"`
	ArtistName string `json:"artistName"`
}
