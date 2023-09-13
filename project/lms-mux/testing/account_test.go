package testing

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

// skenario 1 -> test insert account success
func TestInsertAccountSuccess(t *testing.T) {
	// buat request body
	input := map[string]any{
		"email":    "reoshby2@gmail.com",
		"password": "123",
		"user_id":  16,
	}
	inputJSON, _ := json.Marshal(&input)
	requestBody := strings.NewReader(string(inputJSON))

	// buat request dan recorder
	request := httptest.NewRequest(http.MethodPost, "http://localhost:5000/api/account", requestBody)
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

// skenario 2 -> test insert account failed email already exist
func TestInsertAccountFailedEmailExist(t *testing.T) {
	// buat request body
	input := map[string]any{
		"email":    "reoshby@gmail.com",
		"password": "123",
		"user_id":  16,
	}
	inputJSON, _ := json.Marshal(&input)
	requestBody := strings.NewReader(string(inputJSON))

	// buat request recorder
	request := httptest.NewRequest(http.MethodPost, "http://localhost:5000/api/account", requestBody)
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
	assert.Equal(t, 500, response.StatusCode)
	assert.Equal(t, 500, int(responseBody["status_code"].(float64)))
	assert.Equal(t, "internal server error", responseBody["status"].(string))
}

// skenario 3 -> test insert account failed bad request
func TestInsertAccountBadRequest(t *testing.T) {
	// buat request body
	input := map[string]any{
		"email":    "",
		"password": "",
		"user_id":  16,
	}
	inputJSON, _ := json.Marshal(&input)
	requestBody := strings.NewReader(string(inputJSON))

	// buat request dan recorder
	request := httptest.NewRequest(http.MethodPost, "http://localhost:5000/api/account", requestBody)
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

// skenario 4 -> test insert account internal server error
func TestInsertAccountInternalServerError(t *testing.T) {
	// buat request body
	input := map[string]any{
		"email":    "reoshby1@gmail.com",
		"password": "123",
		"user_id":  99999,
	}
	inputJSON, _ := json.Marshal(&input)
	requestBody := strings.NewReader(string(inputJSON))

	// buat request dan recorder
	request := httptest.NewRequest(http.MethodPost, "http://localhost:5000/api/account", requestBody)
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

// skenario 5 -> test update account success
func TestUpdateAccountSuccess(t *testing.T) {
	// buat request body
	input := map[string]any{
		"id":       6,
		"email":    "reoshby2@gmail.com",
		"password": "123",
		"user_id":  16,
	}
	inputJSON, _ := json.Marshal(&input)
	requestBody := strings.NewReader(string(inputJSON))

	// buat request dan recorder
	request := httptest.NewRequest(http.MethodPut, "http://localhost:5000/api/account", requestBody)
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

// skenario 6 -> test update account failed bad request
func TestUpdateAccountFailedBadRequest(t *testing.T) {
	input := map[string]any{
		"id":       6,
		"email":    "",
		"password": "",
		"user_id":  16,
	}
	inputJSON, _ := json.Marshal(&input)
	requestBody := strings.NewReader(string(inputJSON))

	// buat request dan recorder
	request := httptest.NewRequest(http.MethodPut, "http://localhost:5000/api/account", requestBody)
	request.Header.Add("content-type", "application/json")

	recorder := httptest.NewRecorder()

	handler_test.ServeHTTP(recorder, request)

	// simpan hasil response
	response := recorder.Result()
	body, _ := io.ReadAll(response.Body)
	var responseBody map[string]any
	json.Unmarshal(body, &responseBody)

	// cek hasil test
	fmt.Println(string(body))
	assert.Equal(t, 400, response.StatusCode)
	assert.Equal(t, 400, int(responseBody["status_code"].(float64)))
	assert.Equal(t, "bad request", responseBody["status"].(string))
}

// skenario 7 -> test update account failed not found
func TestUpdateAccountFailedNotFound(t *testing.T) {
	// buat request body
	input := map[string]any{
		"id":       99999,
		"email":    "reoshby@gmail.com",
		"password": "123",
		"user_id":  16,
	}
	inputJSON, _ := json.Marshal(&input)
	requestBody := strings.NewReader(string(inputJSON))

	// buat request dan recorder
	request := httptest.NewRequest(http.MethodPut, "http://localhost:5000/api/account", requestBody)
	request.Header.Add("content-type", "application/json")

	recorder := httptest.NewRecorder()

	handler_test.ServeHTTP(recorder, request)

	// simpan hasil response
	response := recorder.Result()
	body, _ := io.ReadAll(response.Body)
	var responseBody map[string]any
	json.Unmarshal(body, &responseBody)

	// cek hasil test
	fmt.Println(string(body))
	assert.Equal(t, 404, response.StatusCode)
	assert.Equal(t, 404, int(responseBody["status_code"].(float64)))
	assert.Equal(t, "not found", responseBody["status"].(string))
}

// skenario 8 -> test update account failed internal server error
func TestUpdateAccountFailedInternalError(t *testing.T) {
	// buat request body
	input := map[string]any{
		"id":       6,
		"email":    "reoshby@gmail.com",
		"password": "123",
		"user_id":  9999,
	}
	inputJSON, _ := json.Marshal(&input)
	requestBody := strings.NewReader(string(inputJSON))

	// buat request dan recorder
	request := httptest.NewRequest(http.MethodPut, "http://localhost:5000/api/account", requestBody)
	request.Header.Add("content-type", "application/json")

	recorder := httptest.NewRecorder()

	handler_test.ServeHTTP(recorder, request)

	// simpan hasil response
	response := recorder.Result()
	body, _ := io.ReadAll(response.Body)
	var responseBody map[string]any
	json.Unmarshal(body, &responseBody)

	// cek hasil test
	fmt.Println(string(body))
	assert.Equal(t, 500, response.StatusCode)
	assert.Equal(t, 500, int(responseBody["status_code"].(float64)))
	assert.Equal(t, "internal server error", responseBody["status"].(string))
}

// skenario 9 -> test delete account success
func TestDeleteAccountSuccess(t *testing.T) {
	// buat request dan recorder
	request := httptest.NewRequest(http.MethodDelete, "http://localhost:5000/api/account/6", nil)
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

// skenario 10 -> test delete account failed not found
func TestDeleteAccountFailed(t *testing.T) {
	// buat request dan recorder
	request := httptest.NewRequest(http.MethodDelete, "http://localhost:5000/api/account/99999", nil)
	recorder := httptest.NewRecorder()

	handler_test.ServeHTTP(recorder, request)

	// simpan hasil response
	response := recorder.Result()
	body, _ := io.ReadAll(response.Body)
	var responseBody map[string]any
	json.Unmarshal(body, &responseBody)

	// cek hasil test
	fmt.Println(string(body))
	assert.Equal(t, 404, response.StatusCode)
	assert.Equal(t, 404, int(responseBody["status_code"].(float64)))
	assert.Equal(t, "not found", responseBody["status"].(string))
}

// skenario 11 -> test get all data accounts succcess
func TestGetAllAccountsSuccess(t *testing.T) {
	// buat response dan recorder
	request := httptest.NewRequest(http.MethodGet, "http://localhost:5000/api/accounts", nil)
	recorder := httptest.NewRecorder()

	handler_test.ServeHTTP(recorder, request)

	// simpan hasil response
	response := recorder.Result()
	body, _ := io.ReadAll(response.Body)
	var responseBody map[string]any
	json.Unmarshal(body, &responseBody)

	// cek hasil test
	fmt.Println(string(body))
	assert.Equal(t, 200, response.StatusCode)
	assert.Equal(t, 200, int(responseBody["status_code"].(float64)))
	assert.Equal(t, "ok", responseBody["status"].(string))
}

// skenario 12 -> test get data account by Id success
func TestGetAccountById(t *testing.T) {
	// buat request dan recorder
	request := httptest.NewRequest(http.MethodGet, "http://localhost:5000/api/account/5", nil)
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

// skenario 13 -> test get data account by Id failed not found
func TestGetAccountByIdFailed(t *testing.T) {
	// buat request dan recorder
	request := httptest.NewRequest(http.MethodGet, "http://localhost:5000/api/account/999", nil)
	recorder := httptest.NewRecorder()

	handler_test.ServeHTTP(recorder, request)

	// simpan hasil test
	response := recorder.Result()
	body, _ := io.ReadAll(response.Body)
	var responseBody map[string]any
	json.Unmarshal(body, &responseBody)

	// cek hasil test
	fmt.Println(string(body))
	assert.Equal(t, 404, response.StatusCode)
	assert.Equal(t, 404, int(responseBody["status_code"].(float64)))
	assert.Equal(t, "not found", responseBody["status"].(string))
}

// skenario 14 -> test get data account by email success
func TestGetAccountByEmailSuccess(t *testing.T) {
	// buat request dan recorder
	request := httptest.NewRequest(http.MethodGet, "http://localhost:5000/api/account?email=reoshby@gmail.com", nil)
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

// skenario 15 -> test get data account by email failed not found
func TestGetAccountByEmailFailed(t *testing.T) {
	// buat request dan recorder
	request := httptest.NewRequest(http.MethodGet, "http://localhost:5000/api/account?email=reo@gmail.com", nil)
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
