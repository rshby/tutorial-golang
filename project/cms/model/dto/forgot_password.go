package dto

// create request untuk forgot password
type ForgotPasswordRequest struct {
	Email              string `json:"email" validate:"required"`
	OldPassword        string `json:"old_password" validate:"required"`
	NewPassword        string `json:"new_password" validate:"required"`
	ConfirmNewPassword string `json:"confirm_new_password" validate:"required"`
	OTP                string `json:"otp" validate:"required"`
}

type ForgotPasswordResponse struct {
	Email    string `json:"email,omitempty"`
	Password string `json:"password,omitempty"`
}
