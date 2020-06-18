package data

// HealthResponse is the response to /health.
type HealthResponse struct {
	Health string `json:"health"`
}
