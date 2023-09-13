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

type EducationHandler struct {
	EducationService *service.EducationService
}

func NewEducationHandler(educationService *service.EducationService) *EducationHandler {
	return &EducationHandler{
		EducationService: educationService,
	}
}

// handler function insert
func (e *EducationHandler) Insert(w http.ResponseWriter, r *http.Request) {
	// decode request body
	request := web.RequestEducationInsert{}
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
	education, err := e.EducationService.Insert(r.Context(), &request)
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

	// success insert data education
	helper.ResponseSuccess(w, http.StatusOK, "ok", "success insert data education", education)
}

// handler function update
func (e *EducationHandler) Update(w http.ResponseWriter, r *http.Request) {
	// decode request body
	request := web.RequestEducationUpdate{}
	err := json.NewDecoder(r.Body).Decode(&request)
	if err != nil {
		// error bad request
		helper.ResponseError(w, http.StatusBadRequest, "bad request", err.Error())
		return
	}

	bodyBytes, _ := json.Marshal(&request)
	r.Body.Close()
	r.Body = io.NopCloser(bytes.NewBuffer(bodyBytes))

	// jalankan perintah update yang ada di service
	education, err := e.EducationService.Update(r.Context(), &request)
	if err != nil {
		excp, ok := err.(validator.ValidationErrors)
		if ok {
			// error bad request -> gagal validasi
			helper.ResponseError(w, http.StatusBadRequest, "bad request", excp.Error())
		} else {
			if strings.Contains(err.Error(), "record not found") {
				// error data not found
				helper.ResponseError(w, http.StatusNotFound, "not found", err.Error())
			} else {
				// error internal server
				helper.ResponseError(w, http.StatusInternalServerError, "internal server error", err.Error())
			}
		}

		return
	}

	// success update data education
	helper.ResponseSuccess(w, http.StatusOK, "ok", "success update data education", education)
}

// handler function delete]
func (e *EducationHandler) Delete(w http.ResponseWriter, r *http.Request) {
	// ambil data id dari params url
	params := mux.Vars(r)
	id, _ := strconv.Atoi(params["id"])

	// jalankan perintah delete yang ada di service
	result, err := e.EducationService.Delete(r.Context(), id)
	if err != nil {
		if strings.Contains(err.Error(), "not found") {
			// error not found
			helper.ResponseError(w, http.StatusNotFound, "not found", err.Error())
		} else {
			// error internal server
			helper.ResponseError(w, http.StatusInternalServerError, "internal server error", err.Error())
		}

		return
	}

	// success delete education
	helper.ResponseSuccess(w, http.StatusOK, "ok", result, nil)
}

// handler function get all education include university
func (e *EducationHandler) GetAll(w http.ResponseWriter, r *http.Request) {
	// jalankan perintah get all yang ada di service
	educations, err := e.EducationService.GetAll(r.Context())
	if err != nil {
		// error data not found
		helper.ResponseError(w, http.StatusNotFound, "not found", err.Error())
		return
	}

	// success get all data educations
	helper.ResponseSuccess(w, http.StatusOK, "ok", "success get all data educations", educations)
}

// handler function get education include university by udication_id
func (e *EducationHandler) GetById(w http.ResponseWriter, r *http.Request) {
	// ambil id yang ada di params url
	params := mux.Vars(r)
	id, _ := strconv.Atoi(params["id"])

	// jalankan perintah getbyId yang ada di service
	education, err := e.EducationService.GetById(r.Context(), id)
	if err != nil {
		// error data not found
		helper.ResponseError(w, http.StatusNotFound, "not found", err.Error())
		return
	}

	// success get data education by Id
	helper.ResponseSuccess(w, http.StatusOK, "ok", "success get data education by Id", education)
}
