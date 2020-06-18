package data

// Response is the response to /results.
type Response struct {
	Items []*Item `json:"items"`
}
