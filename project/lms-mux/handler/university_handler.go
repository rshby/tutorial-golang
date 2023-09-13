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

type UniversityHandler struct {
	UniversityService *service.UniversityService
}

func NewUniversityHandler(univService *service.UniversityService) *UniversityHandler {
	return &UniversityHandler{
		UniversityService: univService,
	}
}

// handler function Insert
func (u *UniversityHandler) Insert(w http.ResponseWriter, r *http.Request) {
	// decode request body
	request := web.RequestUniversityInsert{}
	err := json.NewDecoder(r.Body).Decode(&request)
	if err != nil {
		// error bad request
		helper.ResponseError(w, http.StatusBadRequest, "bad request", err.Error())
		return
	}

	bodyBytes, _ := json.Marshal(&request)
	r.Body.Close()
	r.Body = io.NopCloser(bytes.NewBuffer(bodyBytes))

	// jalankan perintah Insert yang ada di service
	university, err := u.UniversityService.Insert(r.Context(), &request)
	if err != nil {
		excp, ok := err.(validator.ValidationErrors)
		if ok {
			// error bad request -> gagal validasi
			helper.ResponseError(w, http.StatusBadRequest, "bad request", excp.Error())
		} else {
			// error internal server
			helper.ResponseError(w, http.StatusInternalServerError, "internal server error", err.Error())
		}

		return
	}

	// success insert data university
	helper.ResponseSuccess(w, http.StatusOK, "ok", "success insert data university", university)
}

// handler function Update by Id
func (u *UniversityHandler) Update(w http.ResponseWriter, r *http.Request) {
	// decode request body
	request := web.RequestuniversityUpdate{}
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
	university, err := u.UniversityService.Update(r.Context(), &request)
	if err != nil {
		excp, ok := err.(validator.ValidationErrors)
		if ok {
			// error bad request -> gagal validasi
			helper.ResponseError(w, http.StatusBadRequest, "bad request", excp.Error())
		} else {
			if strings.Contains(err.Error(), "not found") {
				// error data not found
				helper.ResponseError(w, http.StatusNotFound, "not found", err.Error())
			} else {
				// error internal server
				helper.ResponseError(w, http.StatusInternalServerError, "internal server error", err.Error())
			}
		}

		return
	}

	// success insert data university
	helper.ResponseSuccess(w, http.StatusOK, "ok", "success update data university", university)
}

// handler function delete university by Id
func (u *UniversityHandler) Delete(w http.ResponseWriter, r *http.Request) {
	// ambil id dari params url
	params := mux.Vars(r)
	id, _ := strconv.Atoi(params["id"])

	// jalankan perintah delete yang ada di service
	result, err := u.UniversityService.Delete(r.Context(), id)
	if err != nil {
		if strings.Contains(err.Error(), "not found") {
			// error data not found
			helper.ResponseError(w, http.StatusNotFound, "not found", err.Error())
		} else {
			// error internal server
			helper.ResponseError(w, http.StatusInternalServerError, "internal server error", err.Error())
		}

		return
	}

	// success delete data university by Id
	helper.ResponseSuccess(w, http.StatusOK, "ok", result, nil)
}

// handler functio get all universities
func (u *UniversityHandler) GetAll(w http.ResponseWriter, r *http.Request) {
	// jalankan perintah GetAll yang ada di service
	universities, err := u.UniversityService.GetAll(r.Context())
	if err != nil {
		// error data not found
		helper.ResponseError(w, http.StatusNotFound, "not found", err.Error())
		return
	}

	// success get all data universities
	helper.ResponseSuccess(w, http.StatusOK, "ok", "success get all data universities", universities)
}

// handler functio get university by Id
func (u *UniversityHandler) GetById(w http.ResponseWriter, r *http.Request) {
	// ambil id dari params url
	params := mux.Vars(r)
	id, _ := strconv.Atoi(params["id"])

	// jalankan perintah GetById yang ada di service
	university, err := u.UniversityService.GetById(r.Context(), id)
	if err != nil {
		// error data not found
		helper.ResponseError(w, http.StatusNotFound, "not found", err.Error())
		return
	}

	// success get full university information by id
	helper.ResponseSuccess(w, http.StatusOK, "ok", "success get full university information by Id", university)
}

// handler function get all educations by university_Id
func (u *UniversityHandler) GetAllEducationsByUniversityId(w http.ResponseWriter, r *http.Request) {
	// ambil university_Id yang ada di params url
	params := mux.Vars(r)
	universityId, _ := strconv.Atoi(params["id"])

	// jalankan perintah get all educations by university_id yg ada di repository
	educations, err := u.UniversityService.GetAllEducationsByUniversityId(r.Context(), universityId)
	if err != nil {
		// error data not found
		helper.ResponseError(w, http.StatusNotFound, "not found", err.Error())
		return
	}

	// success get all educations by university_Id
	helper.ResponseSuccess(w, http.StatusOK, "ok", "success get all educations by university_id", educations)
}
