package dto

type LoginResponse struct {
	Email   string `json:"email,omitempty"`
	LoginAt string `json:"login_at,omitempty"`
	Token   string `json:"token,omitempty"`
}
