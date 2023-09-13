package handler

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"net/http"
	"strings"
	"training/app/model/dto"
	"training/app/service"
)

type UserHandler struct {
	UserService service.IUserService
}

// create function provider
func NewUserHandler(userService service.IUserService) *UserHandler {
	return &UserHandler{
		UserService: userService,
	}
}

// method to signup
func (u *UserHandler) SignUp(c *gin.Context) {
	// decode request_body
	var request dto.SignUpRequest
	err := c.ShouldBindJSON(&request)
	if err != nil {
		msg, ok := err.(validator.ValidationErrors)
		if ok {
			// error validasi
			var errMessage []string
			for _, item := range msg {
				errMessage = append(errMessage, fmt.Sprintf("error field: %v, condition: %v", item.Field(), item.ActualTag()))
			}
			c.JSON(http.StatusBadRequest, &dto.ApiMessage{
				StatusCode: http.StatusBadRequest,
				Status:     "bad request",
				Message:    strings.Join(errMessage, ". "),
			})
			return
		}

		c.JSON(http.StatusBadRequest, &dto.ApiMessage{
			StatusCode: http.StatusBadRequest,
			Status:     "bad request",
			Message:    err.Error(),
		})
		return
	}

	// call procedure signup in service
	result, err := u.UserService.SignUp(&request)
	if err != nil {
		if strings.Contains(strings.ToLower(err.Error()), "already exist in database") || strings.Contains(strings.ToLower(err.Error()), "hash password") {
			// bad request
			c.JSON(http.StatusBadRequest, &dto.ApiMessage{
				StatusCode: http.StatusBadRequest,
				Status:     "bad request",
				Message:    err.Error(),
			})
			return
		}

		c.JSON(http.StatusInternalServerError, &dto.ApiMessage{
			StatusCode: http.StatusInternalServerError,
			Status:     "internal server error",
			Message:    err.Error(),
		})
		return
	}

	// success create user
	c.JSON(http.StatusOK, &dto.ApiMessage{
		StatusCode: http.StatusOK,
		Status:     "ok",
		Message:    "success create new user",
		Data:       result,
	})
}

// method to login
func (u *UserHandler) Login(c *gin.Context) {
	// decode request body
	var request dto.LoginRequest
	err := c.ShouldBindJSON(&request)
	if err != nil {
		// jika terkena vaidasi
		msg, ok := err.(validator.ValidationErrors)
		if ok {
			var errMessage []string
			for _, item := range msg {
				errMessage = append(errMessage, fmt.Sprintf("error field: %v, condition: %v", item.Field(), item.ActualTag()))
			}
			c.JSON(http.StatusBadRequest, &dto.ApiMessage{
				StatusCode: http.StatusBadRequest,
				Status:     "bad request",
				Message:    strings.Join(errMessage, ". "),
			})
			return
		}

		c.JSON(http.StatusBadRequest, &dto.ApiMessage{
			StatusCode: http.StatusBadRequest,
			Status:     "bad request",
			Message:    err.Error(),
		})
		return
	}

	// call procedure login in service layer
	result, err := u.UserService.Login(&request)
	if err != nil {
		if strings.Contains(strings.ToLower(err.Error()), "not match") {
			// bad request
			c.JSON(http.StatusBadRequest, &dto.ApiMessage{
				StatusCode: http.StatusBadRequest,
				Status:     "bad request",
				Message:    err.Error(),
			})
			return
		}

		if strings.Contains(strings.ToLower(err.Error()), "not found") {
			// not found
			c.JSON(http.StatusNotFound, &dto.ApiMessage{
				StatusCode: http.StatusNotFound,
				Status:     "not found",
				Message:    err.Error(),
			})
			return
		}

		// internal server error
		c.JSON(http.StatusInternalServerError, &dto.ApiMessage{
			StatusCode: http.StatusInternalServerError,
			Status:     "internal server error",
			Message:    err.Error(),
		})
		return
	}

	// success login
	c.JSON(http.StatusOK, &dto.ApiMessage{
		StatusCode: http.StatusOK,
		Status:     "ok",
		Message:    "success login",
		Data:       result,
	})
}
