package dto

// create object request email
type CheckAccountEmailRequest struct {
	Email string `json:"email" validate:"required"`
}

type CheckAccountUsernameRequest struct {
	Username string `json:"username" validate:"required"`
}

// create object response email
type CheckAccountResponse struct {
	AlreadyExist bool           `json:"already_exist,omitempty"`
	Account      *AccountDetail `json:"account,omitempty"`
}

type AccountDetail struct {
	Id        int64  `json:"id,omitempty"`
	Email     string `json:"email,omitempty"`
	Username  string `json:"username,omitempty"`
	Password  string `json:"password,omitempty"`
	CreatedAt string `json:"created_at,omitempty"`
}
