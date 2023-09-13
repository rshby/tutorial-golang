package dto

// create object model API response JSON
type ApiResponse struct {
	StatusCode int    `json:"status_code,omitempty"`
	Status     string `json:"status,omitempty"`
	Message    string `json:"message,omitempty"`
	Data       any    `json:"data,omitempty"`
}
