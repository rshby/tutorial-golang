package dto

// request insert logger
type LoggerInsertRequest struct {
	IpAddress  string `json:"ip_address" validate:"required"`
	UrlPath    string `json:"url_path" validate:"required"`
	Method     string `json:"method" validate:"required"`
	StatusCode int    `json:"status_code" validate:"required"`
	Status     string `json:"status" validate:"required"`
	Duration   string `json:"duration" validate:"required"`
}
