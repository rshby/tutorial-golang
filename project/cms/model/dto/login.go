package dto

// buat object struct untuk login
type LoginRequest struct {
	Username string `json:"username" validate:"required"`
	Password string `json:"password" validate:"required"`
}

// buat object struct untuk response setelah login success
type LoginResponse struct {
	IsLoggin bool   `json:"is_loggin,omitempty"`
	Token    string `json:"token,omitempty"`
}
