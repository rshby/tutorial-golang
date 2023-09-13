package entity

import (
	"database/sql"
	"time"
)

type Account struct {
	ID         int64          `json:"id"`
	Email      string         `json:"email"`
	Username   string         `json:"username"`
	Password   string         `json:"password"`
	OTP        sql.NullString `json:"otp"`
	ExpiredOTP sql.NullTime   `json:"expired_otp"`
	CreatedAt  time.Time      `json:"created_at"`
	UserId     int64          `json:"user_id"`
}
