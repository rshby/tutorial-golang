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

type DistrictHandler struct {
	DistrictService *service.DistrictService
}

func NewDistrictHandler(districtService *service.DistrictService) *DistrictHandler {
	return &DistrictHandler{
		DistrictService: districtService,
	}
}

// handler function Insert data district
func (d *DistrictHandler) Insert(w http.ResponseWriter, r *http.Request) {
	// decode request body
	request := web.RequestDistrictInsert{}
	err := json.NewDecoder(r.Body).Decode(&request)
	if err != nil {
		panic(err.Error())
	}

	bodyBytes, _ := json.Marshal(&request)
	r.Body.Close()
	r.Body = io.NopCloser(bytes.NewBuffer(bodyBytes))

	// jalankan perintah insert yang ada di DistrictService
	district, err := d.DistrictService.Insert(r.Context(), &request)
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

	// success insert data district
	helper.ResponseSuccess(w, http.StatusOK, "ok", "success insert data district", district)
}

// handler function update data district
func (d *DistrictHandler) Update(w http.ResponseWriter, r *http.Request) {
	// decode request body
	request := web.RequestDistrictUpdate{}
	err := json.NewDecoder(r.Body).Decode(&request)
	if err != nil {
		panic(err.Error())
	}

	bodyBytes, _ := json.Marshal(&request)
	r.Body.Close()
	r.Body = io.NopCloser(bytes.NewBuffer(bodyBytes))

	// jalankan perintah update yang ada di DistrictService
	district, err := d.DistrictService.Update(r.Context(), &request)
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
				// error inernal server error
				helper.ResponseError(w, http.StatusInternalServerError, "internal server error", err.Error())
			}
		}

		return
	}

	// success update data district
	helper.ResponseSuccess(w, http.StatusOK, "ok", "success update data district", district)
}

// handler function delete district by Id
func (d *DistrictHandler) Delete(w http.ResponseWriter, r *http.Request) {
	// ambil id dari params url
	params := mux.Vars(r)
	id, _ := strconv.Atoi(params["id"])

	// jalankan perintah Delete yang ada di DistrictService
	result, err := d.DistrictService.Delete(r.Context(), id)
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

	// success delete data district by Id
	helper.ResponseSuccess(w, http.StatusOK, "ok", result, nil)
}

// handler function get all
func (d *DistrictHandler) GetAll(w http.ResponseWriter, r *http.Request) {
	// jalankan perintah get all districts yang ada di DistrictService
	districts, err := d.DistrictService.GetAll(r.Context())
	if err != nil {
		// error data not found
		helper.ResponseError(w, http.StatusNotFound, "not found", err.Error())
		return
	}

	// success get all data districts
	helper.ResponseSuccess(w, http.StatusOK, "ok", "success get all data districts", districts)
}

// handler function get data district by id
func (d *DistrictHandler) GetById(w http.ResponseWriter, r *http.Request) {
	// ambil id yang ada di params url
	params := mux.Vars(r)
	id, _ := strconv.Atoi(params["id"])

	// jalankan perintah get by id yang ada di DistrictService
	district, err := d.DistrictService.GetById(r.Context(), id)
	if err != nil {
		// error data not found
		helper.ResponseError(w, http.StatusNotFound, "not found", err.Error())
		return
	}

	// success get data district by id
	helper.ResponseSuccess(w, http.StatusOK, "ok", "success get data district by id", district)
}

// handler function get all districts by city_id
func (d *DistrictHandler) GetAllDistrictsByCityId(w http.ResponseWriter, r *http.Request) {
	// ambil data city_id dari params url
	params := mux.Vars(r)
	id, _ := strconv.Atoi(params["id"])

	// jalankan perintah get all district by city_id yang ada di DistrictService
	districts, err := d.DistrictService.GetAllDistrictsByCityId(r.Context(), id)
	if err != nil {
		// error data not found
		helper.ResponseError(w, http.StatusNotFound, "not found", err.Error())
		return
	}

	// success get all data districts by city_id
	helper.ResponseSuccess(w, http.StatusOK, "ok", "success get all data districts by city_id", districts)
}
