package testing

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"lms-mux/model/entity"
	"lms-mux/repository"
	"net/http"
	"net/http/httptest"
	"strconv"
	"strings"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

// skenario 1 -> test insert user success
func TestInsertUserSuccess(t *testing.T) {
	// buat request body
	input := map[string]any{
		"first_name":   "Reo",
		"last_name":    "Sahobby",
		"gender":       "L",
		"birth_date":   "1999-02-05T00:00:00Z",
		"address_id":   2,
		"education_id": 1,
	}
	inputJSON, _ := json.Marshal(&input)
	requestBody := strings.NewReader(string(inputJSON))

	// buat request dan recorder
	request := httptest.NewRequest(http.MethodPost, "http://localhost:5000/api/user", requestBody)
	request.Header.Add("content-type", "application/json")

	recorder := httptest.NewRecorder()

	handler_test.ServeHTTP(recorder, request)

	// simpan hasil response
	response := recorder.Result()
	body, _ := io.ReadAll(response.Body)
	responseBody := map[string]any{}
	json.Unmarshal(body, &responseBody)

	// cek hasil test
	fmt.Println(string(body))
	assert.Equal(t, 200, response.StatusCode)
	assert.Equal(t, 200, int(responseBody["status_code"].(float64)))
	assert.Equal(t, "ok", responseBody["status"].(string))
}

// skenario 2 -> test insert user error bad request
func TestInsertUserFailedBadRequest(t *testing.T) {
	// buat request body
	input := map[string]any{
		"first_name":   "",
		"last_name":    "Sahobby",
		"gender":       "L",
		"birth_date":   "1999-02-05T00:00:00Z",
		"address_id":   2,
		"education_id": 1,
	}
	inputJSON, _ := json.Marshal(&input)
	requestBody := strings.NewReader(string(inputJSON))

	// buat request dan recorder
	request := httptest.NewRequest(http.MethodPost, "http://localhost:5000/api/user", requestBody)
	request.Header.Add("content-type", "application/json")

	recorder := httptest.NewRecorder()

	handler_test.ServeHTTP(recorder, request)

	// simpan hasil response
	response := recorder.Result()
	body, _ := io.ReadAll(response.Body)
	responseBody := map[string]any{}
	json.Unmarshal(body, &responseBody)

	// cek hasil test
	fmt.Println(string(body))
	assert.Equal(t, 400, response.StatusCode)
	assert.Equal(t, 400, int(responseBody["status_code"].(float64)))
	assert.Equal(t, "bad request", responseBody["status"].(string))
}

// skenario 3 -> test insert user error internal server
func TestInsertUserFailedInternalServer(t *testing.T) {
	// buat request body
	input := map[string]any{
		"first_name":   "Reo",
		"last_name":    "Sahobby",
		"gender":       "L",
		"birth_date":   "1999-02-05T00:00:00Z",
		"address_id":   9999,
		"education_id": 9999,
	}
	inputJSON, _ := json.Marshal(&input)
	requestBody := strings.NewReader(string(inputJSON))

	// buat request dan recorder
	request := httptest.NewRequest(http.MethodPost, "http://localhost:5000/api/user", requestBody)
	request.Header.Add("content-type", "application/json")

	recorder := httptest.NewRecorder()

	handler_test.ServeHTTP(recorder, request)

	// simpan hasil response
	response := recorder.Result()
	body, _ := io.ReadAll(response.Body)
	responseBody := map[string]any{}
	json.Unmarshal(body, &responseBody)

	// cek hasil test
	fmt.Println(string(body))
	assert.Equal(t, 500, response.StatusCode)
	assert.Equal(t, 500, int(responseBody["status_code"].(float64)))
	assert.Equal(t, "internal server error", responseBody["status"].(string))
}

// skenario 4 -> test get by id success
func TestGetByIdSuccess(t *testing.T) {
	// buat request dan recorder
	request := httptest.NewRequest(http.MethodGet, "http://localhost:5000/api/user/2", nil)
	recorder := httptest.NewRecorder()

	handler_test.ServeHTTP(recorder, request)

	// simpan hasil response
	response := recorder.Result()
	body, _ := io.ReadAll(response.Body)
	responseBody := map[string]any{}
	json.Unmarshal(body, &responseBody)

	// cekk hasil test
	fmt.Println(string(body))
	assert.Equal(t, 200, response.StatusCode)
	assert.Equal(t, 200, int(responseBody["status_code"].(float64)))
	assert.Equal(t, "ok", responseBody["status"].(string))
}

// skenario 5 -> test get by id failed data not found
func TestGetUserByIdFailed(t *testing.T) {
	// buat request dan recorder
	request := httptest.NewRequest(http.MethodGet, "http://localhost:5000/api/user/9999", nil)
	recorder := httptest.NewRecorder()

	handler_test.ServeHTTP(recorder, request)

	// simpan hasil response
	response := recorder.Result()
	body, _ := io.ReadAll(response.Body)
	responseBody := map[string]any{}
	json.Unmarshal(body, &responseBody)

	// cek hasil test
	fmt.Println(string(body))
	assert.Equal(t, 404, response.StatusCode)
	assert.Equal(t, 404, int(responseBody["status_code"].(float64)))
	assert.Equal(t, "not found", responseBody["status"].(string))
}

// skenario 6 -> test update user by id success
func TestUpdateUserSuccess(t *testing.T) {
	// buat request body
	input := map[string]any{
		"id":           2,
		"first_name":   "Reo",
		"last_name":    "Sahobby",
		"gender":       "L",
		"birth_date":   "1999-02-05T00:00:00Z",
		"address_id":   2,
		"education_id": 1,
	}
	inputJSON, _ := json.Marshal(&input)
	requestBody := strings.NewReader(string(inputJSON))

	// buat request dan recorder
	request := httptest.NewRequest(http.MethodPut, "http://localhost:5000/api/user", requestBody)
	request.Header.Add("content-type", "application/json")

	recorder := httptest.NewRecorder()

	handler_test.ServeHTTP(recorder, request)

	// simpan hasil response
	response := recorder.Result()
	body, _ := io.ReadAll(response.Body)
	responseBody := map[string]any{}
	json.Unmarshal(body, &responseBody)

	// cek hasil test
	fmt.Println(string(body))
	assert.Equal(t, 200, response.StatusCode)
	assert.Equal(t, 200, int(responseBody["status_code"].(float64)))
	assert.Equal(t, "ok", responseBody["status"].(string))
}

// skenario 7 -> test update user failed not found
func TestUpdateUserFailedNotFound(t *testing.T) {
	// buat request body
	input := map[string]any{
		"id":           9999,
		"first_name":   "Reo",
		"last_name":    "Sahobby",
		"gender":       "L",
		"birth_date":   "1999-02-05T00:00:00Z",
		"address_id":   2,
		"education_id": 1,
	}
	inputJSON, _ := json.Marshal(&input)
	requestBody := strings.NewReader(string(inputJSON))

	// buat request dan recorder
	request := httptest.NewRequest(http.MethodPut, "http://localhost:5000/api/user", requestBody)
	request.Header.Add("content-type", "application/json")

	recorder := httptest.NewRecorder()

	handler_test.ServeHTTP(recorder, request)

	// simpan response
	response := recorder.Result()
	body, _ := io.ReadAll(response.Body)
	responseBody := map[string]any{}
	json.Unmarshal(body, &responseBody)

	// cek hasil test
	fmt.Println(string(body))
	assert.Equal(t, 404, response.StatusCode)
	assert.Equal(t, 404, int(responseBody["status_code"].(float64)))
	assert.Equal(t, "not found", responseBody["status"].(string))
}

// skenario 8 -> test update user failed bad request
func TestUpdateUserFailedBadRequest(t *testing.T) {
	// buat request body
	input := map[string]any{
		"id":           2,
		"first_name":   "",
		"last_name":    "Sahobby",
		"gender":       "L",
		"birth_date":   "1999-02-05T00:00:00Z",
		"address_id":   2,
		"education_id": 1,
	}
	inputJSON, _ := json.Marshal(&input)
	requestBody := strings.NewReader(string(inputJSON))

	// buat request dan recorder
	request := httptest.NewRequest(http.MethodPut, "http://localhost:5000/api/user", requestBody)
	request.Header.Add("content-type", "application/json")

	recorder := httptest.NewRecorder()

	handler_test.ServeHTTP(recorder, request)

	// simpan hasil test
	response := recorder.Result()
	body, _ := io.ReadAll(response.Body)
	responseBody := map[string]any{}
	json.Unmarshal(body, &responseBody)

	// cek hasil test
	fmt.Println(string(body))
	assert.Equal(t, 400, response.StatusCode)
	assert.Equal(t, 400, int(responseBody["status_code"].(float64)))
	assert.Equal(t, "bad request", responseBody["status"].(string))
}

// skenario 9 -> test update user failed internal server
func TestUpdateuserFailedInternalServerError(t *testing.T) {
	// buat request body
	input := map[string]any{
		"id":           2,
		"first_name":   "Reo",
		"last_name":    "Sahobby",
		"gender":       "L",
		"birth_date":   "1999-02-05T00:00:00Z",
		"address_id":   99999,
		"education_id": 99999,
	}
	inputJSON, _ := json.Marshal(&input)
	requestBody := strings.NewReader(string(inputJSON))

	// buat request dan recorder
	request := httptest.NewRequest(http.MethodPut, "http://localhost:5000/api/user", requestBody)
	request.Header.Add("content-type", "application/json")

	recorder := httptest.NewRecorder()

	handler_test.ServeHTTP(recorder, request)

	// simpan hasil response
	response := recorder.Result()
	body, _ := io.ReadAll(response.Body)
	responseBody := map[string]any{}
	json.Unmarshal(body, &responseBody)

	// cek hasil test
	fmt.Println(string(body))
	assert.Equal(t, 500, response.StatusCode)
	assert.Equal(t, 500, int(responseBody["status_code"].(float64)))
	assert.Equal(t, "internal server error", responseBody["status"].(string))
}

// skenario 10 -> test delete user by id success
func TestDeleteUserSuccess(t *testing.T) {
	// tambahkan data menggunakan repository
	userRepo := repository.NewGeneralRepository(NewDatabaseTesting(), &entity.User{})
	user := entity.User{
		FirstName:   "Reo",
		LastName:    "Sahobby",
		Gender:      "L",
		BirthDate:   time.Now(),
		AddressId:   2,
		EducationId: 1,
	}
	userInsert, _ := userRepo.Insert(context.TODO(), &user)

	// buat request dan recorder
	request := httptest.NewRequest(http.MethodDelete, "http://localhost:5000/api/user/"+strconv.Itoa(userInsert.Id), nil)
	//request := httptest.NewRequest(http.MethodDelete, "http://localhost:5000/api/user/2", nil)
	recorder := httptest.NewRecorder()

	handler_test.ServeHTTP(recorder, request)

	// simpan hasil response
	response := recorder.Result()
	body, _ := io.ReadAll(response.Body)
	responseBody := map[string]any{}
	json.Unmarshal(body, &responseBody)

	// cek hasil test
	fmt.Println(string(body))
	assert.Equal(t, 200, response.StatusCode)
	assert.Equal(t, 200, int(responseBody["status_code"].(float64)))
	assert.Equal(t, "ok", responseBody["status"].(string))
}

// skenario 11 -> test delete user by id failed not found
func TestDeleteUserFailed(t *testing.T) {
	// buat request dan recorder
	request := httptest.NewRequest(http.MethodDelete, "http://localhost:5000/api/user/999", nil)
	recorder := httptest.NewRecorder()

	handler_test.ServeHTTP(recorder, request)

	// simpan hasil response
	response := recorder.Result()
	body, _ := io.ReadAll(response.Body)
	responseBody := map[string]any{}
	json.Unmarshal(body, &responseBody)

	// cek hasil test
	fmt.Println(string(body))
	assert.Equal(t, 404, response.StatusCode)
	assert.Equal(t, 404, int(responseBody["status_code"].(float64)))
	assert.Equal(t, "not found", responseBody["status"].(string))
}
