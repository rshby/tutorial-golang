package controller

import (
	"cms/helper"
	"cms/service"
	"net/http"
)

// create object struct UserController
type UserController struct {
	UserService *service.UserService
}

// function provider to create object UserController
func NewUserController(userService *service.UserService) *UserController {
	return &UserController{
		UserService: userService,
	}
}

// handler get user by email
func (u *UserController) GetByEmail(w http.ResponseWriter, r *http.Request) {
	email := r.URL.Query().Get("email")

	// jika parameter email kosong
	if email == "" || email == " " {
		helper.ResponseError(w, http.StatusBadRequest, "email parameter is required")
		return
	}

	// call procedure getByEmail in service
	result, err := u.UserService.GetByEmail(r.Context(), email)

	// jika ada error ketika get data
	if err != nil {
		helper.ResponseError(w, http.StatusNotFound, err.Error())
		return
	}

	// success get data
	helper.ResponseSuccess(w, "success get data user by email", result)
}

// handler get all data users
func (u *UserController) GetAll(w http.ResponseWriter, r *http.Request) {
	// call procedure get all users in service
	results, err := u.UserService.GetAll(r.Context())

	// jika ada error
	if err != nil {
		helper.ResponseError(w, http.StatusNotFound, err.Error())
		return
	}

	// success get all data users
	helper.ResponseSuccess(w, "success get all data users", results)
}
