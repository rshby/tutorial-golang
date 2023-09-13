package handler

import (
	"bytes"
	"encoding/json"
	"io"
	"lms-mux/helper"
	"lms-mux/model/entity"
	"lms-mux/model/web"
	"lms-mux/service"
	"net/http"
	"strconv"
	"strings"
	"sync"
	"time"

	"github.com/go-playground/validator/v10"
	"github.com/gorilla/mux"
	"github.com/sirupsen/logrus"
)

type UserHandler struct {
	UserService *service.UserService
}

func NewUserHandler(userService *service.UserService) *UserHandler {
	return &UserHandler{
		UserService: userService,
	}
}

// handler function Insert
func (u *UserHandler) Insert(w http.ResponseWriter, r *http.Request) {
	// decode request body
	request := web.RequestUserInsert{}
	err := json.NewDecoder(r.Body).Decode(&request)
	if err != nil {
		// error bad request -> gagal decode
		helper.ResponseError(w, http.StatusBadRequest, "bad request", err.Error())
		return
	}

	bodyBytes, _ := json.Marshal(&request)
	r.Body.Close()
	r.Body = io.NopCloser(bytes.NewBuffer(bodyBytes))

	// jalankan perintah Insert yang ada di Service
	user, err := u.UserService.Insert(r.Context(), &request)
	if err != nil {
		excp, ok := err.(validator.ValidationErrors)
		if ok {
			// error bad request -> gagal validasi
			helper.ResponseError(w, http.StatusBadRequest, "bad request", excp.Error())
		} else {
			if strings.Contains(err.Error(), "not found") {
				// error data not found
				helper.ResponseError(w, http.StatusNotFound, "not found", err.Error())
			} else {
				// error internal server
				helper.ResponseError(w, http.StatusInternalServerError, "internal server error", err.Error())
			}
		}

		return
	}

	// success insert data user
	helper.ResponseSuccess(w, http.StatusOK, "ok", "success insert data user", user)
}

// handler function Update data user by Id
func (u *UserHandler) Update(w http.ResponseWriter, r *http.Request) {
	// decode request body
	request := web.RequestUserUpdate{}
	err := json.NewDecoder(r.Body).Decode(&request)
	if err != nil {
		// error bad request -> gagal decode
		helper.ResponseError(w, http.StatusBadRequest, "bad request", err.Error())
		return
	}

	bodyBytes, _ := json.Marshal(&request)
	r.Body.Close()
	r.Body = io.NopCloser(bytes.NewBuffer(bodyBytes))

	// jalankan perintah Update yang ada di service
	user, err := u.UserService.Update(r.Context(), &request)
	if err != nil {
		excp, ok := err.(validator.ValidationErrors)
		if ok {
			// error bad request -> gagal validasi
			helper.ResponseError(w, http.StatusBadRequest, "bad request", excp.Error())
		} else {
			if strings.Contains(err.Error(), "not found") {
				// error not found
				helper.ResponseError(w, http.StatusNotFound, "not found", err.Error())
			} else {
				// error internal server
				helper.ResponseError(w, http.StatusInternalServerError, "internal server error", err.Error())
			}
		}

		return
	}

	// success update data user by Id
	helper.ResponseSuccess(w, http.StatusOK, "ok", "success update data user by Id", user)
}

// handler function delete user by Id
func (u *UserHandler) Delete(w http.ResponseWriter, r *http.Request) {
	// ambil id dari params url
	params := mux.Vars(r)
	id, _ := strconv.Atoi(params["id"])

	// jalankan perintah Delete yang ada di service
	result, err := u.UserService.Delete(r.Context(), id)
	if err != nil {
		if strings.Contains(err.Error(), "not found") {
			// error data not found
			helper.ResponseError(w, http.StatusNotFound, "not found", err.Error())
		} else {
			// error internal server
			helper.ResponseError(w, http.StatusInternalServerError, "internal server error", err.Error())
		}

		return
	}

	// success delete data user by Id
	helper.ResponseSuccess(w, http.StatusOK, "ok", result, nil)
}

// handler function Get All users
func (u *UserHandler) GetAll(w http.ResponseWriter, r *http.Request) {
	// jalankan perintah GetAll yang ada di service
	users, err := u.UserService.GetAllByBatch(r.Context())
	if err != nil {
		// error data not found
		helper.ResponseError(w, http.StatusNotFound, "not found", err.Error())
		return
	}

	// success get all data users information
	helper.ResponseSuccess(w, http.StatusOK, "ok", "success get all data users information", users)
}

// Handler function GetById
func (u *UserHandler) GetById(w http.ResponseWriter, r *http.Request) {
	// ambil id dari params url
	params := mux.Vars(r)
	id, _ := strconv.Atoi(params["id"])

	// jalankan perintah GetById yang ada di service
	user, err := u.UserService.GetUserInformationById(r.Context(), id)
	if err != nil {
		// error data not found
		helper.ResponseError(w, http.StatusNotFound, "not found", err.Error())
		return
	}

	// success get data user by Id
	helper.ResponseSuccess(w, http.StatusOK, "ok", "success get data user by Id", user)
}

// handler function get by id
func (u *UserHandler) GetByIdFromEntity(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	id, _ := strconv.Atoi(params["id"])

	// jalankan perintah getbyId yang ada di service
	user, err := u.UserService.GetById(r.Context(), id)
	if err != nil {
		helper.ResponseError(w, http.StatusNotFound, "not found", err.Error())
		return
	}

	// success get data
	helper.ResponseSuccess(w, http.StatusOK, "ok", "success get data by Id", user)
}

// handler function create mass user
func (u *UserHandler) CreateMassUser(w http.ResponseWriter, r *http.Request) {
	wg := &sync.WaitGroup{}
	mtx := &sync.RWMutex{}
	for i := 0; i < 100000; i++ {
		wg.Add(1)
		go func(i int) {
			mtx.Lock()
			user := entity.User{
				FirstName:   "reo" + strconv.Itoa(i),
				LastName:    "sahobby" + strconv.Itoa(i),
				Gender:      "L",
				BirthDate:   time.Now(),
				AddressId:   6,
				EducationId: 2,
			}

			// insert ke database menggunakan service
			_, err := u.UserService.UserRepository.General.Insert(r.Context(), &user)
			if err != nil {
				panic(err.Error())
			}

			logrus.WithField("event", "UserHandler-CreateMassUser").Info("success create user-" + strconv.Itoa(i))
			mtx.Unlock()
			wg.Done()
		}(i)
	}

	wg.Wait()
	helper.ResponseSuccess(w, http.StatusOK, "ok", "success insert mass user", nil)
}
