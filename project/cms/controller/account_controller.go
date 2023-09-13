package controller

import (
	"cms/helper"
	"cms/model/dto"
	"cms/service"
	"encoding/json"
	"net/http"
	"strings"

	"github.com/go-playground/validator/v10"
)

// create object AccountController
type AccountController struct {
	AccountService *service.AccountService
}

// function provider to create new object AccountController
func NewAccountController(accService *service.AccountService) *AccountController {
	return &AccountController{
		AccountService: accService,
	}
}

// handler Create Account
func (a *AccountController) CreateAccount(w http.ResponseWriter, r *http.Request) {
	// decode request body
	var request dto.CreateAccountRequest
	err := json.NewDecoder(r.Body).Decode(&request)
	if err != nil {
		// error ketika decode
		helper.ResponseError(w, http.StatusInternalServerError, err.Error())
		return
	}

	// call function createAccount in service
	response, err := a.AccountService.CreateAccount(r.Context(), &request)
	if err != nil {
		errBadReq, ok := err.(validator.ValidationErrors)

		// jika error bad request gagal validasi required
		if ok {
			helper.ResponseError(w, http.StatusBadRequest, errBadReq.Error())
			return
		} else {
			// jika bukan error bad request gagal validasi

			// jika error not found
			if strings.Contains(err.Error(), "not found") {
				helper.ResponseError(w, http.StatusNotFound, err.Error())
				return
			} else if strings.Contains(err.Error(), "same email") || strings.Contains(err.Error(), "same username") {
				helper.ResponseError(w, http.StatusBadRequest, err.Error())
				return
			} else {
				helper.ResponseError(w, http.StatusInternalServerError, err.Error())
				return
			}
		}
	}

	// success create new account
	helper.ResponseSuccess(w, "success create new account", response)
}

// hancler check account
func (a *AccountController) CheckAccount(w http.ResponseWriter, r *http.Request) {
	var requestEmail dto.CheckAccountEmailRequest
	var requestUsername dto.CheckAccountUsernameRequest

	// get url parameter
	if email := r.URL.Query().Get("email"); email != "" {
		requestEmail = dto.CheckAccountEmailRequest{
			Email: email,
		}
	} else if username := r.URL.Query().Get("username"); username != "" {
		requestUsername = dto.CheckAccountUsernameRequest{
			Username: username,
		}
	} else {
		// error bad request -> harap masukkan param
		helper.ResponseError(w, http.StatusBadRequest, "choose parameters with email or username")
		return
	}

	// jika parameter yang diinput email
	if (requestEmail != dto.CheckAccountEmailRequest{}) {
		// call procedure in service
		data, err := a.AccountService.CheckAccountEmail(r.Context(), &requestEmail)
		if err != nil {
			errBadReq, ok := err.(validator.ValidationErrors)

			// jika error karena gagal validasi required
			if ok {
				helper.ResponseError(w, http.StatusBadRequest, errBadReq.Error())
				return
			} else {
				// jika error not found
				if strings.Contains(err.Error(), "not found") {
					helper.ResponseError(w, http.StatusNotFound, err.Error())
					return
				} else {
					helper.ResponseError(w, http.StatusInternalServerError, err.Error())
					return
				}
			}
		}

		// success check account by email
		helper.ResponseSuccess(w, "success get data account by email", data)
	} else {
		// jika parameter yang diinput username
		data, err := a.AccountService.CheckAccountUsername(r.Context(), &requestUsername)
		if err != nil {
			errBadReq, ok := err.(validator.ValidationErrors)

			// jika error gagal validasi required
			if ok {
				helper.ResponseError(w, http.StatusBadRequest, errBadReq.Error())
				return
			} else {
				// jika error data not found
				if strings.Contains(err.Error(), "not found") {
					helper.ResponseError(w, http.StatusNotFound, err.Error())
					return
				} else {
					helper.ResponseError(w, http.StatusInternalServerError, err.Error())
					return
				}
			}
		}

		// success get data account by username
		helper.ResponseSuccess(w, "success get data account by username", data)
	}
}

// handler login
func (a *AccountController) Login(w http.ResponseWriter, r *http.Request) {
	// decode request body
	var request dto.LoginRequest
	err := json.NewDecoder(r.Body).Decode(&request)
	if err != nil {
		helper.ResponseError(w, http.StatusInternalServerError, err.Error())
		return
	}

	// call procedure login in service
	loginResult, err := a.AccountService.Login(r.Context(), &request)
	if err != nil {
		errBadReq, ok := err.(validator.ValidationErrors)

		// jika error gagal validasi
		if ok {
			helper.ResponseError(w, http.StatusBadRequest, errBadReq.Error())
			return
		} else {
			if strings.Contains(err.Error(), "wrong") {
				// error salah password
				helper.ResponseError(w, http.StatusBadRequest, err.Error())
				return
			} else if strings.Contains(err.Error(), "not found") || strings.Contains(err.Error(), "not exist") {
				// error account not found
				helper.ResponseError(w, http.StatusNotFound, err.Error())
				return
			} else {
				// error internal server
				helper.ResponseError(w, http.StatusInternalServerError, err.Error())
				return
			}
		}
	}

	// success login
	helper.ResponseSuccess(w, "success login", loginResult)
}

// handler request otp
func (a *AccountController) RequestOTP(w http.ResponseWriter, r *http.Request) {
	// decode request body
	var request dto.OtpRequest
	err := json.NewDecoder(r.Body).Decode(&request)

	// jika ada kesalahan ketika decode request body
	if err != nil {
		helper.ResponseError(w, http.StatusInternalServerError, err.Error())
		return
	}

	// call procedure in service to
	response, err := a.AccountService.RequestOTP(r.Context(), &request)
	if err != nil {
		errBadReq, ok := err.(validator.ValidationErrors)
		if ok {
			// jika error bad request gagal validasi
			helper.ResponseError(w, http.StatusBadRequest, errBadReq.Error())
			return
		} else {
			// jika error selain gagal validasi

			// jika error email tidak ditemukan
			if strings.Contains(err.Error(), "email") {
				helper.ResponseError(w, http.StatusNotFound, err.Error())
				return
			} else {
				helper.ResponseError(w, http.StatusInternalServerError, err.Error())
				return
			}
		}
	}

	// success request OTP
	helper.ResponseSuccess(w, "success request OTP", response)
}

// handler forgot-password
func (a *AccountController) ForgotPassword(w http.ResponseWriter, r *http.Request) {
	// decode request body
	var request dto.ForgotPasswordRequest
	err := json.NewDecoder(r.Body).Decode(&request)
	if err != nil {
		// jika error ketika decode
		helper.ResponseError(w, http.StatusInternalServerError, err.Error())
		return
	}

	// call procedure ForgotPassword in service
	result, err := a.AccountService.ForgotPassword(r.Context(), &request)
	if err != nil {
		errBadReq, ok := err.(validator.ValidationErrors)
		if ok {
			helper.ResponseError(w, http.StatusBadRequest, errBadReq.Error())
			return
		}

		// jika bukan error validasi
		if strings.Contains(err.Error(), "not exist") || strings.Contains(err.Error(), "not found") {
			// jika error not found
			helper.ResponseError(w, http.StatusNotFound, err.Error())
			return
		}

		// jika error bad request
		if strings.Contains(err.Error(), "wrong") || strings.Contains(err.Error(), "must be same") {
			helper.ResponseError(w, http.StatusBadRequest, err.Error())
			return
		}

		// error internal server
		helper.ResponseError(w, http.StatusInternalServerError, err.Error())
		return
	}

	// success forgot password
	helper.ResponseSuccess(w, "success forgot password", result)
}

// handler Change Password
func (a *AccountController) ChangePassword(w http.ResponseWriter, r *http.Request) {
	// decode data from request body
	var request dto.ChangePasswordRequest
	err := json.NewDecoder(r.Body).Decode(&request)
	if err != nil {
		// error internal server : gagal decode
		helper.ResponseError(w, http.StatusInternalServerError, err.Error())
		return
	}

	// call procedure change_password in service
	result, err := a.AccountService.ChangePassword(r.Context(), &request)

	// jika ada error
	if err != nil {
		errBadReq, ok := err.(validator.ValidationErrors)

		// jika error bad request gagal validasi
		if ok {
			helper.ResponseError(w, http.StatusBadRequest, errBadReq.Error())
			return
		}

		// jika error not found
		if strings.Contains(err.Error(), "not exist") || strings.Contains(err.Error(), "not found") {
			helper.ResponseError(w, http.StatusNotFound, err.Error())
			return
		}

		// jika error bad request
		if strings.Contains(err.Error(), "wrong") || strings.Contains(err.Error(), "must be same") {
			helper.ResponseError(w, http.StatusBadRequest, err.Error())
			return
		}

		// jika error internal server error
		helper.ResponseError(w, http.StatusInternalServerError, err.Error())
		return
	}

	// success update password
	helper.ResponseSuccess(w, "success update password", result)
}
