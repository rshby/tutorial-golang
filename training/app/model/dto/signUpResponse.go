package dto

type SignUpResponse struct {
	Email     string `json:"email,omitempty"`
	Password  string `json:"password,omitempty"`
	CreatedAt string `json:"created_at,omitempty"`
}
