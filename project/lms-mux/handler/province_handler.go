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

type ProvinceHandler struct {
	ProvinceService *service.ProvinceService
}

func NewProvinceHandler(provinceService *service.ProvinceService) *ProvinceHandler {
	return &ProvinceHandler{
		ProvinceService: provinceService,
	}
}

// handler function Insert province
func (p *ProvinceHandler) Insert(w http.ResponseWriter, r *http.Request) {
	// decode request body
	request := web.RequestProvinceInsert{}
	err := json.NewDecoder(r.Body).Decode(&request)
	if err != nil {
		panic(err.Error())
	}

	bodyBytes, _ := json.Marshal(&request)
	r.Body.Close()
	r.Body = io.NopCloser(bytes.NewBuffer(bodyBytes))

	// jalankan perintah Insert province yang ada di service
	province, err := p.ProvinceService.Insert(r.Context(), &request)
	if err != nil {
		excp, ok := err.(validator.ValidationErrors)
		if ok {
			// error bad request -> gagal validasi
			helper.ResponseError(w, http.StatusBadRequest, "bad request", excp.Error())
			return
		} else {
			// error internal server error
			helper.ResponseError(w, http.StatusInternalServerError, "internal server error", err.Error())
			return
		}
	}

	// success insert data province
	helper.ResponseSuccess(w, http.StatusOK, "ok", "success insert data province", province)
}

// handler Update province
func (p *ProvinceHandler) Update(w http.ResponseWriter, r *http.Request) {
	// ambil id
	params := mux.Vars(r)
	id, _ := strconv.Atoi(params["id"])

	// decode request body
	request := web.RequestProvinceUpdate{}
	err := json.NewDecoder(r.Body).Decode(&request)
	if err != nil {
		panic(err.Error())
	}

	requestJSON, _ := json.Marshal(&request)
	r.Body.Close()
	r.Body = io.NopCloser(bytes.NewBuffer(requestJSON))

	// jalankan perintah update province yang ada di service
	request.Id = id
	province, err := p.ProvinceService.Update(r.Context(), &request)
	if err != nil {
		excp, ok := err.(validator.ValidationErrors)
		if ok {
			// error bad request -> gagal validasi
			helper.ResponseError(w, http.StatusBadRequest, "bad request", excp.Error())
			return
		} else {
			if strings.Contains(err.Error(), "record not found") {
				// error not found
				helper.ResponseError(w, http.StatusNotFound, "not found", err.Error())
				return
			} else {
				// error internal server
				helper.ResponseError(w, http.StatusInternalServerError, "internal server error", err.Error())
				return
			}
		}
	}

	// success update data
	helper.ResponseSuccess(w, http.StatusOK, "ok", "success update data", province)
}

// handler function Delete
func (p *ProvinceHandler) Delete(w http.ResponseWriter, r *http.Request) {
	// ambil id
	params := mux.Vars(r)
	id, _ := strconv.Atoi(params["id"])

	// jalankan perintah delete province yang ada di service
	response, err := p.ProvinceService.Delete(r.Context(), id)
	if err != nil {
		if strings.Contains(err.Error(), "record not found") {
			// error data not found
			helper.ResponseError(w, http.StatusNotFound, "not found", err.Error())
			return
		} else {
			// error internal server -> gagal delete
			helper.ResponseError(w, http.StatusInternalServerError, "internal server error", err.Error())
			return
		}
	}

	// success delete data province
	helper.ResponseSuccess(w, http.StatusOK, "ok", response, nil)
}

// handler function Get All provinces
func (p *ProvinceHandler) GetAll(w http.ResponseWriter, r *http.Request) {
	// jalankan perintah get all provinces yang ada di service
	provinces, err := p.ProvinceService.GetAll(r.Context())
	if err != nil {
		// error not found
		helper.ResponseError(w, http.StatusNotFound, "not found", err.Error())
		return
	}

	// success get all data provinces
	helper.ResponseSuccess(w, http.StatusOK, "ok", "success get all data provinces", provinces)
}

// handler function get province by Id
func (p *ProvinceHandler) GetById(w http.ResponseWriter, r *http.Request) {
	// ambil id
	params := mux.Vars(r)
	id, _ := strconv.Atoi(params["id"])

	// jalankan perintah GetById yang ada di service
	province, err := p.ProvinceService.GetById(r.Context(), id)
	if err != nil {
		// error data not found
		helper.ResponseError(w, http.StatusNotFound, "not found", err.Error())
		return
	}

	// success get data province by id
	helper.ResponseSuccess(w, http.StatusOK, "ok", "success get data by province by id", province)
}
