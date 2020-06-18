package data

// ErrorResponse is returned from the HTTP server when an error has happened.
type ErrorResponse struct {
	Error string `json:"error"`
}
