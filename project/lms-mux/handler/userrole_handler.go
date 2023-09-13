package handler

import (
	"bytes"
	"encoding/json"
	"io"
	"lms-mux/helper"
	"lms-mux/model/entity"
	"lms-mux/model/web"
	"lms-mux/service"
	"net/http"
	"strconv"
	"strings"

	"github.com/go-playground/validator/v10"
	"github.com/gorilla/mux"
)

type UserRoleHandler struct {
	UserRoleService *service.UserRoleService
}

func NewUserRoleHandler(userRoleService *service.UserRoleService) *UserRoleHandler {
	return &UserRoleHandler{
		UserRoleService: userRoleService,
	}
}

// handler function Insert
func (u *UserRoleHandler) Insert(w http.ResponseWriter, r *http.Request) {
	// decode request body
	request := web.RequestUserRoleInsert{}
	err := json.NewDecoder(r.Body).Decode(&request)
	if err != nil {
		// error bad request
		helper.ResponseError(w, http.StatusBadRequest, "bad request", err.Error())
		return
	}

	// jalankan perintah insert yang ada di service
	userRole, err := u.UserRoleService.Insert(r.Context(), &request)
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

	// success insert data user-role
	helper.ResponseSuccess(w, http.StatusOK, "ok", "success insert data userRole", userRole)
}

// handler function Update
func (u *UserRoleHandler) Update(w http.ResponseWriter, r *http.Request) {
	// decode request body
	request := web.RequestUserRoleUpdate{}
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
	userRole, err := u.UserRoleService.Update(r.Context(), &request)
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

	// success update data user-role by Id
	helper.ResponseSuccess(w, http.StatusOK, "ok", "success update data user-role by id", userRole)
}

// handler Function Delete
func (u *UserRoleHandler) Delete(w http.ResponseWriter, r *http.Request) {
	// ambil id yang ada di params url
	params := mux.Vars(r)
	id, _ := strconv.Atoi(params["id"])

	// jalankan perintah delete yang ada di service
	result, err := u.UserRoleService.Delete(r.Context(), id)
	if err != nil {
		if strings.Contains(err.Error(), "not found") {
			// error data not found
			helper.ResponseError(w, http.StatusNotFound, "not found", err.Error())
		} else {
			// error internal server
			helper.ResponseError(w, http.StatusInternalServerError, "internal server", err.Error())
		}

		return
	}

	// success delete data user-role by Id
	helper.ResponseSuccess(w, http.StatusOK, "ok", result, nil)
}

// handler function GetAll
func (u *UserRoleHandler) GetAll(w http.ResponseWriter, r *http.Request) {
	// jalankan perintah get all yang ada di service
	userRoles, err := u.UserRoleService.GetAll(r.Context())
	if err != nil {
		// error not found
		helper.ResponseError(w, http.StatusNotFound, "not found", err.Error())
		return
	}

	// success get all data user-roles
	helper.ResponseSuccess(w, http.StatusOK, "ok", "success get all data user-roles", userRoles)
}

// handler function GetById
func (u *UserRoleHandler) GetById(w http.ResponseWriter, r *http.Request) {
	// ambil id dari params url
	params := mux.Vars(r)
	id, _ := strconv.Atoi(params["id"])

	// jalankan perintah getbyid yang ada di service
	userRole, err := u.UserRoleService.GetById(r.Context(), id)
	if err != nil {
		// error data not found
		helper.ResponseError(w, http.StatusNotFound, "not found", err.Error())
		return
	}

	// success get data user-role by id
	helper.ResponseSuccess(w, http.StatusOK, "ok", "success get data user-role by id", userRole)
}

// handler function get by user_id or and role_id
func (u *UserRoleHandler) GetByUserIdAndRoleId(w http.ResponseWriter, r *http.Request) {
	// ambil parameter
	userId := r.URL.Query().Get("userid")
	roleId := r.URL.Query().Get("roleid")

	var response []entity.UserRole

	if userId != "" && roleId == "" {
		// get data by user_id
		userID, _ := strconv.Atoi(userId)
		users, err := u.UserRoleService.GetByUserId(r.Context(), userID)
		if err != nil {
			// error data not found
			helper.ResponseError(w, http.StatusNotFound, "not found", err.Error())
			return
		}

		// success
		response = users
	}

	if roleId != "" && userId == "" {
		// get data by role_id
		roleID, _ := strconv.Atoi(roleId)
		users, err := u.UserRoleService.GetByRoleId(r.Context(), roleID)
		if err != nil {
			// error data not found
			helper.ResponseError(w, http.StatusNotFound, "not found", err.Error())
			return
		}

		// success
		response = users
	}

	if userId != "" && roleId != "" {
		// get data by user_id and role_id
		userID, _ := strconv.Atoi(userId)
		roleID, _ := strconv.Atoi(roleId)

		user, err := u.UserRoleService.GetByUserIdAndRoleId(r.Context(), userID, roleID)
		if err != nil {
			// error data not found
			helper.ResponseError(w, http.StatusNotFound, "not found", err.Error())
			return
		}

		// success
		helper.ResponseSuccess(w, http.StatusOK, "ok", "success get data", user)
		return
	}

	// success get data by user_id OR by role_id
	helper.ResponseSuccess(w, http.StatusOK, "ok", "success get data", response)
}
