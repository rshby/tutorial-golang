package dto

// create object to request otp
type OtpRequest struct {
	Email string `json:"email" validate:"required"`
}

// create object to response otp
type OtpResponse struct {
	OTP        string `json:"otp,omitempty"`
	ExpiredOtp string `json:"expired_otp,omitempty"`
}
