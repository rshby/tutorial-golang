package controller

import (
	"cms/helper"
	"cms/model/dto"
	"cms/service"
	"encoding/json"
	"net/http"
	"strings"

	"github.com/go-playground/validator/v10"
	"github.com/gorilla/mux"
)

// create object DislikeController
type DislikeController struct {
	DislikeService *service.DislikeService
}

// function provider to create new object DislikeService
func NewDislikeController(dislikeService *service.DislikeService) *DislikeController {
	return &DislikeController{
		DislikeService: dislikeService,
	}
}

// handler dislike
func (d *DislikeController) Dislike(w http.ResponseWriter, r *http.Request) {
	// decode request body
	var request dto.DislikeRequest
	err := json.NewDecoder(r.Body).Decode(&request)

	// jika ada error ketika decode request_body
	if err != nil {
		helper.ResponseError(w, http.StatusInternalServerError, err.Error())
		return
	}

	// call procedure Dislike in service
	result, err := d.DislikeService.Dislike(r.Context(), &request)
	if err != nil {
		errBadReq, ok := err.(validator.ValidationErrors)

		// jika error bad request -> gagal validasi required
		if ok {
			helper.ResponseError(w, http.StatusBadRequest, errBadReq.Error())
			return
		}

		// jika error not found
		if strings.Contains(err.Error(), "not found") {
			helper.ResponseError(w, http.StatusNotFound, err.Error())
			return
		}

		// error internal server
		helper.ResponseError(w, http.StatusInternalServerError, err.Error())
		return
	}

	// success dislike
	helper.ResponseSuccess(w, "success dislike to this content", result)
}

// handler undislike
func (d *DislikeController) Undislike(w http.ResponseWriter, r *http.Request) {
	// decode request body
	var request dto.DislikeRequest
	err := json.NewDecoder(r.Body).Decode(&request)

	// jika ada error ketika decode request_body
	if err != nil {
		helper.ResponseError(w, http.StatusInternalServerError, err.Error())
		return
	}

	// call procedure Undislike in service
	result, err := d.DislikeService.Undislike(r.Context(), &request)
	if err != nil {
		// cek apabila error bad request -> gagal validasi
		errBadReq, ok := err.(validator.ValidationErrors)
		if ok {
			helper.ResponseError(w, http.StatusBadRequest, errBadReq.Error())
			return
		}

		// jika error not found
		if strings.Contains(err.Error(), "not found") {
			helper.ResponseError(w, http.StatusNotFound, err.Error())
			return
		}

		// error internal server
		helper.ResponseError(w, http.StatusInternalServerError, err.Error())
		return
	}

	// success undislike
	helper.ResponseSuccess(w, "success undislike", result)
}

// handler get contents disliked by username
func (d *DislikeController) GetContentDislikedByUsername(w http.ResponseWriter, r *http.Request) {
	// get username params
	username := mux.Vars(r)["username"]

	if username == "" || username == " " {
		helper.ResponseError(w, http.StatusBadRequest, "username is required")
		return
	}

	// call procedure get content disliked by username in service
	result, err := d.DislikeService.GetContentDislikedByUsername(r.Context(), username)
	if err != nil {
		// jika error not found
		if strings.Contains(strings.ToLower(err.Error()), strings.ToLower("not found")) {
			helper.ResponseError(w, http.StatusNotFound, err.Error())
			return
		}

		// jika error internal server
		helper.ResponseError(w, http.StatusInternalServerError, err.Error())
		return
	}

	// success get
	helper.ResponseSuccess(w, "success get content disliked by username", result)
}
