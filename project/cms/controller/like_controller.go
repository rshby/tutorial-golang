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

// create object LikeController
type LikeController struct {
	LikeService *service.LikeService
}

// function provider to create object LikeController
func NewLikeController(likeService *service.LikeService) *LikeController {
	return &LikeController{
		LikeService: likeService,
	}
}

// handler like
func (l *LikeController) Like(w http.ResponseWriter, r *http.Request) {
	// decode request body
	var request dto.LikeRequest
	err := json.NewDecoder(r.Body).Decode(&request)
	if err != nil {
		helper.ResponseError(w, http.StatusInternalServerError, err.Error())
		return
	}

	// call procedure Like in service
	result, err := l.LikeService.Like(r.Context(), &request)
	if err != nil {
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

	// success like
	helper.ResponseSuccess(w, "success like content", result)
}

// handler unlike
func (l *LikeController) Unlike(w http.ResponseWriter, r *http.Request) {
	// decode request body
	var request dto.LikeRequest
	err := json.NewDecoder(r.Body).Decode(&request)

	// jika ada kesalahan ketika decode request body
	if err != nil {
		helper.ResponseError(w, http.StatusInternalServerError, err.Error())
		return
	}

	// call procedure Unlike in service
	result, err := l.LikeService.Unlike(r.Context(), &request)
	if err != nil {
		errBadReq, ok := err.(validator.ValidationErrors)

		// jika error bad request -> gagal valdasi required
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

	// success unlike
	helper.ResponseSuccess(w, "success unlike this content", result)
}

// handler get content liked by username
func (l *LikeController) GetContentLikedByUsername(w http.ResponseWriter, r *http.Request) {
	// get params
	username := mux.Vars(r)["username"]

	if username == "" || username == " " {
		helper.ResponseError(w, http.StatusBadRequest, "username is required")
		return
	}

	// call procedure GetContentLikedByUsername in service
	result, err := l.LikeService.GetContentLikedByUsername(r.Context(), username)
	if err != nil {
		// jika error data tidak ditemukan
		if strings.Contains(strings.ToLower(err.Error()), strings.ToLower("not found")) {
			helper.ResponseError(w, http.StatusNotFound, err.Error())
			return
		}

		// jika error internal server
		helper.ResponseError(w, http.StatusInternalServerError, err.Error())
		return
	}

	// success get data liked by username
	helper.ResponseSuccess(w, "success get contents liked by username", result)
}
