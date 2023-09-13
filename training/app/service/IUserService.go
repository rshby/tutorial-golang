package service

import (
	"training/app/model/dto"
)

type IUserService interface {
	SignUp(request *dto.SignUpRequest) (*dto.SignUpResponse, error)
	Login(request *dto.LoginRequest) (*dto.LoginResponse, error)
}
