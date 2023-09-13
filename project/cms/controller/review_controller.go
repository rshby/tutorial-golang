package controller

import (
	"cms/helper"
	"cms/model/dto"
	"cms/service"
	"encoding/json"
	"net/http"
	"strings"

	"github.com/go-playground/validator/v10"
)

// create layer object ReviewController
type ReviewController struct {
	ReviewService *service.ReviewService
}

// function provider
func NewReviewController(reviewService *service.ReviewService) *ReviewController {
	return &ReviewController{
		ReviewService: reviewService,
	}
}

// handler create review
func (r *ReviewController) CreateReview(w http.ResponseWriter, req *http.Request) {
	// decode request body
	var request dto.CreateReviewRequest
	err := json.NewDecoder(req.Body).Decode(&request)

	// jika ada error ketika decode
	if err != nil {
		helper.ResponseError(w, http.StatusInternalServerError, err.Error())
		return
	}

	// call procedure create in service
	result, err := r.ReviewService.Insert(req.Context(), &request)
	if err != nil {
		errBadReq, ok := err.(validator.ValidationErrors)
		if ok {
			helper.ResponseError(w, http.StatusBadRequest, errBadReq.Error())
			return
		}

		// jika error not found
		if strings.Contains(strings.ToLower(err.Error()), strings.ToLower("not found")) {
			helper.ResponseError(w, http.StatusNotFound, err.Error())
			return
		}

		// jika error internal server
		helper.ResponseError(w, http.StatusInternalServerError, err.Error())
		return
	}

	// success create review
	helper.ResponseSuccess(w, "success create review", result)
}

// handler delete review
func (r *ReviewController) Delete(w http.ResponseWriter, req *http.Request) {
	// decode request body
	var request dto.DeleteReviewRequest
	err := json.NewDecoder(req.Body).Decode(&request)

	// jika ada error ketika proses decode
	if err != nil {
		helper.ResponseError(w, http.StatusInternalServerError, err.Error())
		return
	}

	// call procedure delete in service
	err = r.ReviewService.Delete(req.Context(), &request)
	if err != nil {
		errBadReq, ok := err.(validator.ValidationErrors)
		if ok {
			// jika error bad request gagal validasi
			helper.ResponseError(w, http.StatusBadRequest, errBadReq.Error())
			return
		}

		// jika error not found
		if strings.Contains(strings.ToLower(err.Error()), strings.ToLower("not found")) {
			helper.ResponseError(w, http.StatusNotFound, err.Error())
			return
		}

		// jika error internal server
		helper.ResponseError(w, http.StatusInternalServerError, err.Error())
		return
	}

	// success delete data review
	helper.ResponseSuccess(w, "success delete review", nil)
}
