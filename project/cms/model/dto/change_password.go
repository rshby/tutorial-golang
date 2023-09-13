package dto

// create struct dto request change password
type ChangePasswordRequest struct {
	Email           string `json:"email" validate:"required"`
	OldPassword     string `json:"old_password" validate:"required"`
	NewPassword     string `json:"new_password" validate:"required"`
	ConfirmPassword string `json:"confirm_password" validate:"required"`
}

type ChangePasswordResponse struct {
	Email    string `json:"email,omitempty"`
	Password string `json:"password,omitempty"`
}
