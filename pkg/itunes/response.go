package itunes

type Response struct {
	Results []Result `json:"results"`
}

type Result struct {
	TrackName  string `json:"trackName"`
	ArtistName string `json:"artistName"`
}
