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

type RoleHandler struct {
	RoleService *service.RoleService
}

func NewRoleHandler(roleService *service.RoleService) *RoleHandler {
	return &RoleHandler{
		RoleService: roleService,
	}
}

// handler function Insert
func (rl *RoleHandler) Insert(w http.ResponseWriter, r *http.Request) {
	// decode request body
	request := web.RequestRoleInsert{}
	err := json.NewDecoder(r.Body).Decode(&request)
	if err != nil {
		// error bad request -> gagal decode
		helper.ResponseError(w, http.StatusBadRequest, "bad request", err.Error())
		return
	}

	// jalankan perintah Insert yang ada di service
	role, err := rl.RoleService.Insert(r.Context(), &request)
	if err != nil {
		excp, ok := err.(validator.ValidationErrors)
		if ok {
			// error bad request
			helper.ResponseError(w, http.StatusBadRequest, "bad request", excp.Error())
		} else {
			// error internal server
			helper.ResponseError(w, http.StatusInternalServerError, "internal server error", err.Error())
		}

		return
	}

	// success insert data role to database
	helper.ResponseSuccess(w, http.StatusOK, "ok", "success insert data role", role)
}

// handler function Update
func (rl *RoleHandler) Update(w http.ResponseWriter, r *http.Request) {
	// decode request body
	request := web.RequestRoleUpdate{}
	err := json.NewDecoder(r.Body).Decode(&request)
	if err != nil {
		// error bad request -> gagal decode
		helper.ResponseError(w, http.StatusBadRequest, "bad request", err.Error())
		return
	}

	bodyBytes, _ := json.Marshal(&request)
	r.Body.Close()
	r.Body = io.NopCloser(bytes.NewBuffer(bodyBytes))

	// jalankan perintah update yang ada di service
	role, err := rl.RoleService.Update(r.Context(), &request)
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

	// success update role by Id
	helper.ResponseSuccess(w, http.StatusOK, "ok", "success update data role", role)
}

// handler function Delete
func (rl *RoleHandler) Delete(w http.ResponseWriter, r *http.Request) {
	// ambil id dari params url
	params := mux.Vars(r)
	id, _ := strconv.Atoi(params["id"])

	// jalankan perintah delete yang ada di service
	result, err := rl.RoleService.Delete(r.Context(), id)
	if err != nil {
		if strings.Contains(err.Error(), "not found") {
			// error not found
			helper.ResponseError(w, http.StatusNotFound, "not found", err.Error())
		} else {
			// error internal server error
			helper.ResponseError(w, http.StatusInternalServerError, "internal server error", err.Error())
		}

		return
	}

	// success delete role by Id
	helper.ResponseSuccess(w, http.StatusOK, "ok", result, nil)
}

// handler functio Get All roles
func (rl *RoleHandler) GetAll(w http.ResponseWriter, r *http.Request) {
	// jalankan perintah getall yang ada di service
	roles, err := rl.RoleService.GetAll(r.Context())
	if err != nil {
		// error not found
		helper.ResponseError(w, http.StatusNotFound, "not found", err.Error())
		return
	}

	// success get all data roles
	helper.ResponseSuccess(w, http.StatusOK, "ok", "success get all data roles", roles)
}

// handler function Get By Id
func (rl *RoleHandler) GetById(w http.ResponseWriter, r *http.Request) {
	// ambil id dari params url
	params := mux.Vars(r)
	id, _ := strconv.Atoi(params["id"])

	// jalankan perintah getbyid yang ada di service
	role, err := rl.RoleService.GetById(r.Context(), id)
	if err != nil {
		// error data not found
		helper.ResponseError(w, http.StatusNotFound, "not found", err.Error())
		return
	}

	// success get by id
	helper.ResponseSuccess(w, http.StatusOK, "ok", "success get data role", role)
}
