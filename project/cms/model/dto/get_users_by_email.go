package dto

// create struct object response Users by email
type UsersByEmail struct {
	Account struct {
		Id       int64  `json:"id,omitempty"`
		Email    string `json:"email,omitempty"`
		Username string `json:"username,omitempty"`
		Password string `json:"password,omitempty"`
	} `json:"account,omitempty"`
	User struct {
		Id        int64  `json:"id,omitempty"`
		FirstName string `json:"first_name,omitempty"`
		LastName  string `json:"last_name,omitempty"`
		Gender    string `json:"gender,omitempty"`
		Address   string `json:"address,omitempty"`
	} `json:"user,omitempty"`
}
