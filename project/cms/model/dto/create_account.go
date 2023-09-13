package dto

import (
	"cms/model/entity"
)

// create struct untuk DTO insert user
type CreateAccountRequest struct {
	Email      string `json:"email" validate:"required"`
	Password   string `json:"password" validate:"required"`
	Username   string `json:"username" validate:"required"`
	FirstName  string `json:"first_name" validate:"required"`
	LastName   string `json:"last_name" validate:"required"`
	IdentityId string `json:"identity_id" validate:"required"`
	Gender     string `json:"gender" validate:"required"`
	Address    string `json:"address" validate:"required"`
}

type CreateAccountResponse struct {
	ID_User       int64          `json:"id_user,omitempty"`
	ID_Account    int64          `json:"id_account,omitempty"`
	AccountDetail *AccountDetail `json:"account_detail,omitempty"`
	UserDetail    *entity.User   `json:"user_detail,omitempty"`
}
