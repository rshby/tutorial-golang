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

type CityHandler struct {
	CityService *service.CityService
}

func NewCityHandler(cityService *service.CityService) *CityHandler {
	return &CityHandler{
		CityService: cityService,
	}
}

// handler function insert data city
func (c *CityHandler) Insert(w http.ResponseWriter, r *http.Request) {
	// decode request body
	request := web.RequestCityInsert{}
	err := json.NewDecoder(r.Body).Decode(&request)
	if err != nil {
		panic(err.Error())
	}

	bodyBytes, _ := json.Marshal(&request)
	r.Body.Close()
	r.Body = io.NopCloser(bytes.NewBuffer(bodyBytes))

	// jalankan perintah insert yang ada di service
	city, err := c.CityService.Insert(r.Context(), &request)
	if err != nil {
		excp, ok := err.(validator.ValidationErrors)
		if ok {
			// error bad request -> gagal validasi
			helper.ResponseError(w, http.StatusBadRequest, "bad request", excp.Error())
			return
		} else {
			// error internal server
			helper.ResponseError(w, http.StatusInternalServerError, "internal server error", err.Error())
			return
		}
	}

	// success insert data city
	helper.ResponseSuccess(w, http.StatusOK, "ok", "success insert data city", city)
}

// handler function update city by id
func (c *CityHandler) Update(w http.ResponseWriter, r *http.Request) {
	// decode request body
	request := web.RequestCityUpdate{}
	err := json.NewDecoder(r.Body).Decode(&request)
	if err != nil {
		panic(err.Error())
	}
	bodyBytes, _ := json.Marshal(&request)
	r.Body.Close()
	r.Body = io.NopCloser(bytes.NewBuffer(bodyBytes))

	// jalankan perintah update yang ada di service
	result, err := c.CityService.Update(r.Context(), &request)
	if err != nil {
		excp, ok := err.(validator.ValidationErrors)
		if ok {
			// error bad request -> gagal validasi
			helper.ResponseError(w, http.StatusBadRequest, "bad request", excp.Error())
			return
		} else {
			if strings.Contains(err.Error(), "record not found") {
				// error data not found
				helper.ResponseError(w, http.StatusNotFound, "not found", err.Error())
				return
			} else {
				// error internal server -> gagal update
				helper.ResponseError(w, http.StatusInternalServerError, "internal server error", err.Error())
				return
			}
		}
	}

	// success update data city by Id
	helper.ResponseSuccess(w, http.StatusOK, "ok", "success update data city by id", result)
}

// handler function delete
func (c *CityHandler) Delete(w http.ResponseWriter, r *http.Request) {
	//ambil id
	params := mux.Vars(r)
	id, _ := strconv.Atoi(params["id"])

	// jalankan perintah Delete city yang ada di service
	result, err := c.CityService.Delete(r.Context(), id)
	if err != nil {
		if strings.Contains(err.Error(), "record not found") {
			// error data not found
			helper.ResponseError(w, http.StatusNotFound, "not found", err.Error())
			return
		} else {
			// error internal server error
			helper.ResponseError(w, http.StatusInternalServerError, "internal server error", err.Error())
			return
		}
	}

	// success delete data city by id
	helper.ResponseSuccess(w, http.StatusOK, "ok", result, nil)
}

// handler function Get all cities
func (c *CityHandler) GetAll(w http.ResponseWriter, r *http.Request) {
	// jalankan perintah get all cities yang ada di service
	cities, err := c.CityService.GetAll(r.Context())
	if err != nil {
		// error data not found
		helper.ResponseError(w, http.StatusNotFound, "not found", err.Error())
		return
	}

	// success get all data cities
	helper.ResponseSuccess(w, http.StatusOK, "ok", "success get all data cities", cities)
}

// handler function Get City by Id
func (c *CityHandler) GetById(w http.ResponseWriter, r *http.Request) {
	// ambil id
	params := mux.Vars(r)
	id, _ := strconv.Atoi(params["id"])

	// jalankan perintah get by id yang ada di service
	city, err := c.CityService.GetById(r.Context(), id)
	if err != nil {
		// error data not found
		helper.ResponseError(w, http.StatusNotFound, "not found", err.Error())
		return
	}

	// success get data city by Id
	helper.ResponseSuccess(w, http.StatusOK, "ok", "success get data city by Id", city)
}

// handler function get all cities by province_id
func (c *CityHandler) GetAllCitiesByProvinceId(w http.ResponseWriter, r *http.Request) {
	// ambil province_id yang ada di params url
	params := mux.Vars(r)
	provinceId, _ := strconv.Atoi(params["province_id"])

	// jalankan perintah get all cities by province_id yang ada di service
	cities, err := c.CityService.GetAllCitiesByProvinceId(r.Context(), provinceId)
	if err != nil {
		// error data not found
		helper.ResponseError(w, http.StatusNotFound, "not found", err.Error())
		return
	}

	// success get all cities by province_id
	helper.ResponseSuccess(w, http.StatusOK, "ok", "success get all cities by province_id", cities)
}
