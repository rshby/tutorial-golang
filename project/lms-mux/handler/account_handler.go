package handler

import (
	"bytes"
	"encoding/json"
	"io"
	"lms-mux/helper"
	"lms-mux/model/web"
	"lms-mux/service"
	"net/http"
	"strconv"
	"strings"

	"github.com/go-playground/validator/v10"
	"github.com/gorilla/mux"
)

type AccountHandler struct {
	AccountService *service.AccountService
}

func NewAccountHandler(accountService *service.AccountService) *AccountHandler {
	return &AccountHandler{
		AccountService: accountService,
	}
}

// handler function Insert
func (a *AccountHandler) Insert(w http.ResponseWriter, r *http.Request) {
	// decode request body
	var request web.RequestAccountInsert

	err := json.NewDecoder(r.Body).Decode(&request)
	if err != nil {
		// error bad request -> gagal decode
		helper.ResponseError(w, http.StatusBadRequest, "bad request", err.Error())
		return
	}

	bodyBytes, _ := json.Marshal(&request)
	r.Body.Close()
	r.Body = io.NopCloser(bytes.NewBuffer(bodyBytes))

	// jalankan perintah Insert yang ada di service
	account, err := a.AccountService.Insert(r.Context(), &request)
	if err != nil {
		excp, ok := err.(validator.ValidationErrors)
		if ok {
			// error bad request -> gagal validasi
			helper.ResponseError(w, http.StatusBadRequest, "bad request", excp.Error())
		} else {
			if strings.Contains(err.Error(), "not found") {
				// error not found
				helper.ResponseError(w, http.StatusNotFound, "not found", err.Error())
			} else {
				// error internal server
				helper.ResponseError(w, http.StatusInternalServerError, "internal server error", err.Error())
			}
		}

		return
	}

	// success insert data account
	helper.ResponseSuccess(w, http.StatusOK, "ok", "success create new account", account)
}

// handler function Update
func (a *AccountHandler) Update(w http.ResponseWriter, r *http.Request) {
	// decode request body
	var request web.RequestUpdateAccount
	err := json.NewDecoder(r.Body).Decode(&request)
	if err != nil {
		// error bad request
		helper.ResponseError(w, http.StatusBadRequest, "bad request", err.Error())
		return
	}

	bodyBytes, _ := json.Marshal(&request)
	r.Body.Close()
	r.Body = io.NopCloser(bytes.NewBuffer(bodyBytes))

	// jalankan perintah Update yang ada di service
	account, err := a.AccountService.Update(r.Context(), &request)
	if err != nil {
		excp, ok := err.(validator.ValidationErrors)
		if ok {
			// error bad request -> gagal validasi
			helper.ResponseError(w, http.StatusBadRequest, "bad request", excp.Error())
		} else {
			if strings.Contains(err.Error(), "not found") {
				// error not found
				helper.ResponseError(w, http.StatusNotFound, "not found", err.Error())
			} else {
				// error internal server
				helper.ResponseError(w, http.StatusInternalServerError, "internal server error", err.Error())
			}
		}

		return
	}

	// success update data account by Id
	helper.ResponseSuccess(w, http.StatusOK, "ok", "success update data account", account)
}

// handler function Delete
func (a *AccountHandler) Delete(w http.ResponseWriter, r *http.Request) {
	// ambil id dari params url
	params := mux.Vars(r)
	id, _ := strconv.Atoi(params["id"])

	// jalankan perintah delete yg ada di service
	result, err := a.AccountService.Delete(r.Context(), id)
	if err != nil {
		// error data not found
		helper.ResponseError(w, http.StatusNotFound, "not found", err.Error())
		return
	}

	// success delete data
	helper.ResponseSuccess(w, http.StatusOK, "ok", result, nil)
}

// handler function GetAll data accounts
func (a *AccountHandler) GetAll(w http.ResponseWriter, r *http.Request) {
	accounts, err := a.AccountService.GetAll(r.Context())
	if err != nil {
		// error data not found
		helper.ResponseError(w, http.StatusNotFound, "not found", err.Error())
		return
	}

	// success get all data accounts
	helper.ResponseSuccess(w, http.StatusOK, "ok", "success get all data accounts", accounts)
}

// handler function get data account by Id
func (a *AccountHandler) GetById(w http.ResponseWriter, r *http.Request) {
	// ambil id dari params url
	params := mux.Vars(r)
	id, _ := strconv.Atoi(params["id"])

	// jalankan perintah get account by Id yang ada di service
	account, err := a.AccountService.GetById(r.Context(), id)
	if err != nil {
		// error data not found
		helper.ResponseError(w, http.StatusNotFound, "not found", err.Error())
		return
	}

	// success get data account by Id
	helper.ResponseSuccess(w, http.StatusOK, "ok", "success get data account by Id", account)
}

// handler function get data account by email
func (a *AccountHandler) GetByEmail(w http.ResponseWriter, r *http.Request) {
	// ambil email
	email := r.URL.Query().Get("email")

	// jalankan perintah GetByEmail yang ada di service
	account, err := a.AccountService.GetAccountByEmail(r.Context(), email)
	if err != nil {
		// error data not found
		helper.ResponseError(w, http.StatusNotFound, "not found", err.Error())
		return
	}

	// success get data account by Email
	helper.ResponseSuccess(w, http.StatusOK, "ok", "success get data account by Email", account)
}
