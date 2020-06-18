package data

// Item is the struct that we return in a list for the client's request as the result.
type Item struct {
	Title    string   `json:"title"`
	Type     string   `json:"type"`
	Creators []string `json:"creators"`
}
