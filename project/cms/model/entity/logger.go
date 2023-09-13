package entity

import "time"

type Logger struct {
	ID         int64     `json:"id,omitempty"`
	IpAddress  string    `json:"ip_address,omitempty"`
	UrlPath    string    `json:"url_path,omitempty"`
	Method     string    `json:"method,omitempty"`
	StatusCode int       `json:"status_code,omitempty"`
	Status     string    `json:"status,omitempty"`
	Duration   string    `json:"duration,omitempty"`
	CreatedAt  time.Time `json:"created_id,omitempty"`
}
