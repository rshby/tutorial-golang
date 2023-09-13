package controller

import (
	"cms/helper"
	"cms/model/dto"
	"cms/service"
	"encoding/json"
	"net/http"
	"strconv"
	"strings"

	"github.com/go-playground/validator/v10"
	"github.com/gorilla/mux"
)

// create object ContentController
type ContentController struct {
	ContentService *service.ContentService
}

// function provider to create new object ContentController
func NewContentController(contentService *service.ContentService) *ContentController {
	return &ContentController{
		ContentService: contentService,
	}
}

// handler Insert
func (c *ContentController) Insert(w http.ResponseWriter, r *http.Request) {
	// decode request body
	var request dto.CreateContentRequest
	err := json.NewDecoder(r.Body).Decode(&request)
	if err != nil {
		helper.ResponseError(w, http.StatusInternalServerError, err.Error())
		return
	}

	// call procedure Insert in service
	result, err := c.ContentService.Insert(r.Context(), &request)
	if err != nil {
		// cek jika error bad request -> gagal validasi
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

	// success insert content
	helper.ResponseSuccess(w, "success insert content", result)
}

// handler edit/update content
func (c *ContentController) Update(w http.ResponseWriter, r *http.Request) {
	// decode request body
	var request dto.UpdateContentRequest
	err := json.NewDecoder(r.Body).Decode(&request)

	// jika ada kesalahan ketika decode
	if err != nil {
		helper.ResponseError(w, http.StatusInternalServerError, err.Error())
		return
	}

	// call procedure Update in service
	result, err := c.ContentService.Update(r.Context(), &request)
	if err != nil {
		errBadReq, ok := err.(validator.ValidationErrors)

		// jika error bad request -> gagal validasi required
		if ok {
			helper.ResponseError(w, http.StatusBadRequest, errBadReq.Error())
			return
		}

		// jika error not found
		if strings.Contains(err.Error(), "not found") || strings.Contains(err.Error(), "not exist") {
			helper.ResponseError(w, http.StatusNotFound, err.Error())
			return
		}

		// jika error internal server
		helper.ResponseError(w, http.StatusInternalServerError, err.Error())
		return
	}

	// success update content by ID
	helper.ResponseSuccess(w, "success update content", result)
}

// handler get all data contents
func (c *ContentController) GetAll(w http.ResponseWriter, r *http.Request) {
	// call procedure get all in service
	contents, err := c.ContentService.GetAll(r.Context())
	if err != nil {
		// jika error not found
		if strings.Contains(strings.ToLower(err.Error()), strings.ToLower("not found")) {
			helper.ResponseError(w, http.StatusNotFound, err.Error())
			return
		}

		// error internal server
		helper.ResponseError(w, http.StatusInternalServerError, err.Error())
		return
	}

	// success get all data contents
	helper.ResponseSuccess(w, "success get all data contents", contents)
}

// handler GetById
func (c *ContentController) GetById(w http.ResponseWriter, r *http.Request) {
	// ambil id
	params := mux.Vars(r)
	id, err := strconv.Atoi(params["id"])
	if err != nil {
		helper.ResponseError(w, http.StatusInternalServerError, err.Error())
		return
	}

	// call procedure GetById in service
	content, err := c.ContentService.GetById(r.Context(), int64(id))
	if err != nil {
		helper.ResponseError(w, http.StatusNotFound, err.Error())
		return
	}

	// success get data content by ID
	helper.ResponseSuccess(w, "success get data content by id", content)
}

// handler get by username
func (c *ContentController) GetByUsername(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	username := params["username"]
	if username == "" || username == " " {
		helper.ResponseError(w, http.StatusBadRequest, "username required")
		return
	}

	// call procedure GetByUsername in service
	contents, err := c.ContentService.GetByUsername(r.Context(), username)
	if err != nil {
		// jika error not found
		if strings.Contains(err.Error(), "not found") {
			helper.ResponseError(w, http.StatusNotFound, err.Error())
			return
		}

		// error internal server
		helper.ResponseError(w, http.StatusInternalServerError, err.Error())
		return
	}

	// success get data content by creator (username)
	helper.ResponseSuccess(w, "success get content by username", contents)
}
