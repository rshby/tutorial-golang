package service

import (
	"errors"
	"fmt"
	"golang.org/x/crypto/bcrypt"
	"time"
	"training/app/helper"
	"training/app/model/dto"
	"training/app/model/entity"
	"training/app/repository"
)

type UserService struct {
	UserRepo repository.IUserRepository
}

// create function provider
func NewUserRepository(userRepo repository.IUserRepository) *UserService {
	return &UserService{
		UserRepo: userRepo,
	}
}

// method to sign up (insert)
func (u *UserService) SignUp(request *dto.SignUpRequest) (*dto.SignUpResponse, error) {
	// 1. cek email apakah sudah ada di database
	_, err := u.UserRepo.GetByEmail(request.Email)
	if err == nil {
		return nil, errors.New("email already exist in database")
	}

	// 2. hash password
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(request.Password), bcrypt.DefaultCost)
	if err != nil {
		return nil, errors.New(fmt.Sprintf("cant hash password: %v", err.Error()))
	}

	// 3. create entity users
	newUser := &entity.User{
		Email:    request.Email,
		Password: string(hashedPassword),
	}

	// 4. call procedure insert in repository
	resultInsert, err := u.UserRepo.InsertUser(newUser)

	// if error
	if err != nil {
		return nil, err
	}

	// 5. success insert
	response := &dto.SignUpResponse{
		Email:     resultInsert.Email,
		Password:  resultInsert.Password,
		CreatedAt: resultInsert.CreatedAt.Format("2006-01-02 15:04:05"),
	}

	return response, nil
}

// method to login
func (u *UserService) Login(request *dto.LoginRequest) (*dto.LoginResponse, error) {
	// 1. cek data email apakah ada di database
	user, err := u.UserRepo.GetByEmail(request.Email)
	if err != nil {
		return nil, errors.New("account not found in database")
	}

	// 2. cek password
	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(request.Password))
	if err != nil {
		return nil, errors.New("password not match")
	}

	// generate token jwt
	token, err := helper.CreateToken(user.Email, int(user.ID))
	if err != nil {
		return nil, err
	}

	// success login
	response := &dto.LoginResponse{
		Email:   user.Email,
		LoginAt: time.Now().Format("2006-01-02 15:04:05"),
		Token:   token,
	}

	return response, nil
}
