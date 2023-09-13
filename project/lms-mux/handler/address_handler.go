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

type AddressHandler struct {
	AddressService *service.AddressService
}

func NewAddressHandler(adService *service.AddressService) *AddressHandler {
	return &AddressHandler{
		AddressService: adService,
	}
}

// handler function Insert
func (a *AddressHandler) Insert(w http.ResponseWriter, r *http.Request) {
	// decode request body
	request := web.RequestAddressInsert{}
	err := json.NewDecoder(r.Body).Decode(&request)
	if err != nil {
		// error bad request
		helper.ResponseError(w, http.StatusBadRequest, "bad request", err.Error())
		return
	}

	bodyBytes, _ := json.Marshal(&request)
	r.Body.Close()
	r.Body = io.NopCloser(bytes.NewBuffer(bodyBytes))

	// jalankan perintah insert yang ada di service
	address, err := a.AddressService.Insert(r.Context(), &request)
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

	// success insert data address
	helper.ResponseSuccess(w, http.StatusOK, "ok", "success insert data address", address)
}

// handler function update
func (a *AddressHandler) Update(w http.ResponseWriter, r *http.Request) {
	// decode request body
	request := web.RequestAddressUpdate{}
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
	address, err := a.AddressService.Update(r.Context(), &request)
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
				// error internal server error
				helper.ResponseError(w, http.StatusInternalServerError, "internal server error", err.Error())
			}
		}

		return
	}

	// success update data address by id
	helper.ResponseSuccess(w, http.StatusOK, "ok", "success update data address by Id", address)
}

// handler function delete
func (a *AddressHandler) Delete(w http.ResponseWriter, r *http.Request) {
	// ambil id dari params url
	params := mux.Vars(r)
	id, _ := strconv.Atoi(params["id"])

	// jalankan perintah Delete yang ada di service
	result, err := a.AddressService.Delete(r.Context(), id)
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

	// success delete data address by Id
	helper.ResponseSuccess(w, http.StatusOK, "ok", result, nil)
}

// handler functio get all addresses
func (a *AddressHandler) GetAll(w http.ResponseWriter, r *http.Request) {
	// jalankan perintah GetAll yang ada di service
	addresses, err := a.AddressService.GetAll(r.Context())
	if err != nil {
		// error not found
		helper.ResponseError(w, http.StatusNotFound, "not found", err.Error())
		return
	}

	// success get all data addresses
	helper.ResponseSuccess(w, http.StatusOK, "ok", "success get all data addresses", addresses)
}

// handler function get address by Id
func (a *AddressHandler) GetById(w http.ResponseWriter, r *http.Request) {
	// ambil params id
	params := mux.Vars(r)
	id, _ := strconv.Atoi(params["id"])

	// jalankan perintah GetByid yang ada di service
	address, err := a.AddressService.GetById(r.Context(), id)
	if err != nil {
		helper.ResponseError(w, http.StatusNotFound, "not found", err.Error())
		return
	}

	// success get data address by Id
	helper.ResponseSuccess(w, http.StatusOK, "ok", "success get data address by Id", address)
}

// handler function get full address by Id
func (a *AddressHandler) GetFullAddressById(w http.ResponseWriter, r *http.Request) {
	// ambil params id
	params := mux.Vars(r)
	id, _ := strconv.Atoi(params["id"])

	// jalankan perintah GetFullAddressById yang ada di service
	address, err := a.AddressService.GetFullAddressById(r.Context(), id)
	if err != nil {
		// error data not found
		helper.ResponseError(w, http.StatusNotFound, "not found", err.Error())
		return
	}

	// success get full address by Id
	helper.ResponseSuccess(w, http.StatusOK, "ok", "success get full address by Id", address)
}
