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

type SubDistrictHandler struct {
	SubDistrictService *service.SubDistrictService
}

func NewSubDistrictHandler(sdService *service.SubDistrictService) *SubDistrictHandler {
	return &SubDistrictHandler{
		SubDistrictService: sdService,
	}
}

// handler function insert data from general
func (s *SubDistrictHandler) Insert(w http.ResponseWriter, r *http.Request) {
	// decode request body
	request := web.RequestSubDistrictInsert{}
	err := json.NewDecoder(r.Body).Decode(&request)
	if err != nil {
		helper.ResponseError(w, http.StatusInternalServerError, "internal server error", err.Error())
		return
	}

	bodyBytes, _ := json.Marshal(&request)
	r.Body.Close()
	r.Body = io.NopCloser(bytes.NewBuffer(bodyBytes))

	// jalankan perintah Insert yang ada di SubDistrictService
	subdistrict, err := s.SubDistrictService.Insert(r.Context(), &request)
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

	// success insert data subdistrict
	helper.ResponseSuccess(w, http.StatusOK, "ok", "success insert data subdistrict", subdistrict)
}

// handler function update data subdistrict
func (s *SubDistrictHandler) Update(w http.ResponseWriter, r *http.Request) {
	// decode request body
	request := web.RequestSubDistrictUpdate{}
	err := json.NewDecoder(r.Body).Decode(&request)
	if err != nil {
		// error bad request -> gagal decode
		helper.ResponseError(w, http.StatusBadRequest, "bad request", err.Error())
		return
	}

	bodyBytes, _ := json.Marshal(&request)
	r.Body.Close()
	r.Body = io.NopCloser(bytes.NewBuffer(bodyBytes))

	// jalankan perintah Update yang ada di SubDistrictService
	subdistrict, err := s.SubDistrictService.Update(r.Context(), &request)
	if err != nil {
		excp, ok := err.(validator.ValidationErrors)
		if ok {
			// error bad request
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

	// success update data subdistrict by Id
	helper.ResponseSuccess(w, http.StatusOK, "ok", "success update data subdistrict by Id", subdistrict)
}

// handler function delete data subdistrict
func (s *SubDistrictHandler) Delete(w http.ResponseWriter, r *http.Request) {
	// ambil id yang ada di params url
	params := mux.Vars(r)
	id, _ := strconv.Atoi(params["id"])

	// jalankan printah Delete yang ada di DistrictService
	result, err := s.SubDistrictService.Delete(r.Context(), id)
	if err != nil {
		if strings.Contains(err.Error(), "record not found") {
			// error data not found
			helper.ResponseError(w, http.StatusNotFound, "not found", err.Error())
		} else {
			// error internal server
			helper.ResponseError(w, http.StatusInternalServerError, "internal server error", err.Error())
		}

		return
	}

	// success delete data subdistrict by Id
	helper.ResponseSuccess(w, http.StatusOK, "ok", result, nil)
}

// handler function get all data subdistricts
func (s *SubDistrictHandler) GetAll(w http.ResponseWriter, r *http.Request) {
	// jalankan perintah Getall yang ada di service
	subDistricts, err := s.SubDistrictService.GetAll(r.Context())
	if err != nil {
		// error data not found
		helper.ResponseError(w, http.StatusNotFound, "not found", err.Error())
		return
	}

	// sucess get all data subdistricts
	helper.ResponseSuccess(w, http.StatusOK, "ok", "success get all data subdistricts", subDistricts)
}

// handler function get data subdistrict by Id
func (s *SubDistrictHandler) GetById(w http.ResponseWriter, r *http.Request) {
	// ambil id dari params url
	params := mux.Vars(r)
	id, _ := strconv.Atoi(params["id"])

	// jalankan perintah GetById yang ada di service
	subdistrict, err := s.SubDistrictService.GetById(r.Context(), id)
	if err != nil {
		// error data not found
		helper.ResponseError(w, http.StatusNotFound, "not found", err.Error())
		return
	}

	// success get data by id
	helper.ResponseSuccess(w, http.StatusOK, "ok", "success get subdistrict by Id", subdistrict)
}

// handlre function get all data subdistricts by district_id
func (s *SubDistrictHandler) GetByDistrictId(w http.ResponseWriter, r *http.Request) {
	// ambil id dari params url
	params := mux.Vars(r)
	id, _ := strconv.Atoi(params["id"])

	// jalankan perintah GetByDistrictId yang ada di service
	subdistricts, err := s.SubDistrictService.GetByDistrictId(r.Context(), id)
	if err != nil {
		// error data not found
		helper.ResponseError(w, http.StatusNotFound, "not found", err.Error())
		return
	}

	// success get all data subdistricts by district_id
	helper.ResponseSuccess(w, http.StatusOK, "ok", "success get all data subdistricts by district_id", subdistricts)
}
