package dto

// create dto object user response
type UserResponse struct {
	Id        int64  `json:"id,omitempty"`
	FirstName string `json:"first_name,omitempty"`
	LastName  string `json:"last_name,omitempty"`
	FullName  string `json:"full_name,omitempty"`
	Gender    string `json:"gender,omitempty"`
	Address   string `json:"address,omitempty"`
	CreatedAt string `json:"created_at,omitempty"`
}
