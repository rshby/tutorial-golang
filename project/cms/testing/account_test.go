package testing

import (
	"cms/model/dto"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

// skenario 1 -> create account success
func TestCreateAccountSuccess(t *testing.T) {
	// create request body
	input := map[string]any{
		"email":       "reoshby@gmail.com",
		"password":    "123",
		"username":    "rshby",
		"first_name":  "Reo",
		"last_name":   "Sahobby",
		"identity_id": "3310250502990002",
		"gender":      "M",
		"address":     "Jl. Cilandak KKO No.57, Ragunan, Pasar Minggu, Jakarta Selatan, DKI Jakarta",
	}

	inputJSON, _ := json.Marshal(&input)
	requestBody := strings.NewReader(string(inputJSON))

	// create request dan recorder
	request := httptest.NewRequest(http.MethodPost, "http://localhost:5000/api/account", requestBody)
	request.Header.Add("content-type", "application/json")

	recorder := httptest.NewRecorder()

	routerTest.ServeHTTP(recorder, request)

	// simpan hasil response body
	response := recorder.Result()
	body, _ := io.ReadAll(response.Body)
	responseBody := map[string]any{}

	json.Unmarshal(body, &responseBody)

	// cek hasil test
	assert.Equal(t, 200, response.StatusCode)
	assert.Equal(t, 200, int(responseBody["status_code"].(float64)))
	assert.Equal(t, "ok", responseBody["status"].(string))
}

// skenario 2 -> create account failed email sudah ada
func TestCreateAccountFailedEmailAlreadyExist(t *testing.T) {
	// create request body
	input := map[string]any{
		"email":       "reoshby@gmail.com",
		"password":    "123",
		"username":    "rshby",
		"first_name":  "Reo",
		"last_name":   "Sahobby",
		"identity_id": "3310250502990002",
		"gender":      "M",
		"address":     "Jl. Cilandak KKO No.57, Ragunan, Pasar Minggu, Jakarta Selatan, DKI Jakarta",
	}
	inputJson, _ := json.Marshal(&input)
	requestBody := strings.NewReader(string(inputJson))

	// create request dan recorder
	request := httptest.NewRequest(http.MethodPost, "http://localhost:5000/api/account", requestBody)
	request.Header.Add("content-type", "application/json")

	recorder := httptest.NewRecorder()

	// execute api
	routerTest.ServeHTTP(recorder, request)

	// simpan hasil execute
	response := recorder.Result()
	body, _ := io.ReadAll(response.Body)
	responseBody := map[string]any{}
	json.Unmarshal(body, &responseBody)

	// cek test
	assert.True(t, strings.Contains(strings.ToLower(responseBody["message"].(string)), strings.ToLower("exist")))
	assert.Equal(t, 400, response.StatusCode)
	assert.Equal(t, 400, int(responseBody["status_code"].(float64)))
}

// skenario 3 -> check account success
func TestCheckAccountSuccess(t *testing.T) {
	// create request & recorder
	request := httptest.NewRequest(http.MethodGet, "http://localhost:5000/api/account-check?username=rshby", nil)
	request.Header.Add("authorization", fmt.Sprintf("Bearer %v", GetTokenLogin()))
	recorder := httptest.NewRecorder()

	// execute api
	routerTest.ServeHTTP(recorder, request)

	// simpan hasil execute
	response := recorder.Result()
	body, _ := io.ReadAll(response.Body)
	responseBody := map[string]any{}
	json.Unmarshal(body, &responseBody)

	// cek hasil test
	assert.Equal(t, 200, response.StatusCode)
	assert.Equal(t, 200, int(responseBody["status_code"].(float64)))
	assert.Equal(t, "ok", responseBody["status"].(string))
}

// skenario 4 -> check account failed not found
func TestCheckAccountFailedNotFound(t *testing.T) {
	// buat request dan recorder
	request := httptest.NewRequest(http.MethodGet, "http://localhost:5000/api/account-check?username=qwerty", nil)
	request.Header.Add("authorization", fmt.Sprintf("Bearer %v", GetTokenLogin()))

	recorder := httptest.NewRecorder()

	// execute api
	routerTest.ServeHTTP(recorder, request)

	// simpan hasil execute
	response := recorder.Result()
	body, _ := io.ReadAll(response.Body)
	responseBody := map[string]any{}
	json.Unmarshal(body, &responseBody)

	// cek hasil
	assert.Equal(t, 404, response.StatusCode)
	assert.Equal(t, 404, int(responseBody["status_code"].(float64)))
	assert.Equal(t, strings.ToLower("not found"), responseBody["status"].(string))
}

// ckenario 5 -> check acount wihthout login before
func TestCheckAccountNotLogin(t *testing.T) {
	// buat request dan recorder
	request := httptest.NewRequest(http.MethodGet, "http://localhost:5000/api/account-check?username=rshby", nil)
	recorder := httptest.NewRecorder()

	// execute api
	routerTest.ServeHTTP(recorder, request)

	// simpan hasil execute
	response := recorder.Result()
	body, _ := io.ReadAll(response.Body)
	responseBody := map[string]any{}
	json.Unmarshal(body, &responseBody)

	// cek hasil
	assert.Equal(t, 401, response.StatusCode)
	assert.Equal(t, 401, int(responseBody["status_code"].(float64)))
	assert.True(t, strings.Contains(strings.ToLower(responseBody["message"].(string)), strings.ToLower("token")))
}

// skenario 6 -> Generate OTP success
func TestGenerateOTPSuccess(t *testing.T) {
	// buat request body
	input := map[string]any{
		"email": "reoshby@gmail.com",
	}
	inputJSON, _ := json.Marshal(&input)
	requestBody := strings.NewReader(string(inputJSON))

	// create request dan recorder
	request := httptest.NewRequest(http.MethodPost, "http://localhost:5000/api/otp", requestBody)
	request.Header.Add("content-type", "application/json")

	recorder := httptest.NewRecorder()

	// execute api
	routerTest.ServeHTTP(recorder, request)

	// simpan hasil execute
	response := recorder.Result()
	body, _ := io.ReadAll(response.Body)
	responseBody := map[string]any{}
	json.Unmarshal(body, &responseBody)

	// cek hasil test
	assert.Equal(t, 200, response.StatusCode)
	assert.Equal(t, 200, int(responseBody["status_code"].(float64)))
	assert.Equal(t, "ok", responseBody["status"].(string))
}

// skenario 7 -> Generate OTP Failed email not found
func TestGenerateOTPFailed(t *testing.T) {
	// buat request body
	input := map[string]any{
		"email": "qwerty@gmail.com",
	}
	inputJSON, _ := json.Marshal(&input)
	requestBody := strings.NewReader(string(inputJSON))

	// create request dan recorder
	request := httptest.NewRequest(http.MethodPost, "http://localhost:5000/api/otp", requestBody)
	recorder := httptest.NewRecorder()

	// execute api
	routerTest.ServeHTTP(recorder, request)

	// simpan hasil
	response := recorder.Result()
	body, _ := io.ReadAll(response.Body)
	responseBody := map[string]any{}
	json.Unmarshal(body, &responseBody)

	// cek hasil
	assert.Equal(t, 404, response.StatusCode)
	assert.Equal(t, 404, int(responseBody["status_code"].(float64)))
	assert.Equal(t, "not found", responseBody["status"].(string))
	assert.True(t, strings.Contains(strings.ToLower(responseBody["message"].(string)), strings.ToLower("not found")))
}

// skenario 8 -> change password success
func TestChangePasswordSuccess(t *testing.T) {
	// siapkan request body
	input := map[string]any{
		"email":            "reoshby@gmail.com",
		"old_password":     "123",
		"new_password":     "111",
		"confirm_password": "111",
	}

	inputJSON, _ := json.Marshal(&input)
	requestBody := strings.NewReader(string(inputJSON))

	// create request dan recorder
	request := httptest.NewRequest(http.MethodPut, "http://localhost:5000/api/change-password", requestBody)
	request.Header.Add("authorization", fmt.Sprintf("Bearer %v", GetTokenLogin()))

	recorder := httptest.NewRecorder()

	// execute api
	routerTest.ServeHTTP(recorder, request)

	// simpan hasil execute
	response := recorder.Result()
	body, _ := io.ReadAll(response.Body)
	responseBody := map[string]any{}
	json.Unmarshal(body, &responseBody)

	// cek hasil
	assert.Equal(t, 200, response.StatusCode)
	assert.Equal(t, 200, int(responseBody["status_code"].(float64)))
	assert.Equal(t, "ok", responseBody["status"].(string))
}

// skenario 9 -> change password wrong old password
func TestChangePasswordFailedWrongOldPassword(t *testing.T) {
	// buat request body
	input := map[string]any{
		"email":            "reoshby@gmail.com",
		"old_password":     "123",
		"new_password":     "111",
		"confirm_password": "111",
	}

	inputJSON, _ := json.Marshal(&input)
	log.Printf("inputJSON : %v\n", inputJSON)
	requestBody := strings.NewReader(string(inputJSON))

	// create request dan recorder
	request := httptest.NewRequest(http.MethodPut, "http://localhost:5000/api/change-password", requestBody)
	request.Header.Add("authorization", fmt.Sprintf("Bearer %v", GetTokenLogin()))

	recorder := httptest.NewRecorder()

	// execute api
	routerTest.ServeHTTP(recorder, request)

	// simpan hasil execute
	response := recorder.Result()
	body, _ := io.ReadAll(response.Body)
	responseBody := map[string]any{}
	json.Unmarshal(body, &responseBody)

	// cek hasil test
	assert.Equal(t, 400, response.StatusCode)
	assert.Equal(t, 400, int(responseBody["status_code"].(float64)))
	assert.Equal(t, "bad request", responseBody["status"].(string))
	assert.True(t, strings.Contains(strings.ToLower(responseBody["message"].(string)), strings.ToLower("wrong")))
}

// skenario 10 -> change password confirm password not same
func TestChangePasswordFailedNotSame(t *testing.T) {
	// create request body
	input := dto.ChangePasswordRequest{
		Email:           "reoshby@gmail.com",
		OldPassword:     "111",
		NewPassword:     "123",
		ConfirmPassword: "456",
	}

	inputJSON, _ := json.Marshal(&input)
	requestBody := strings.NewReader(string(inputJSON))

	// create request dan recorder
	request := httptest.NewRequest(http.MethodPut, "http://localhost:5000/api/change-password", requestBody)
	request.Header.Add("authorization", fmt.Sprintf("Bearer %v", GetTokenLogin()))

	recorder := httptest.NewRecorder()

	// execute api
	routerTest.ServeHTTP(recorder, request)

	// simpan hasil execute
	response := recorder.Result()
	body, _ := io.ReadAll(response.Body)
	responseBody := map[string]any{}
	json.Unmarshal(body, &responseBody)

	// cek hasil test
	assert.Equal(t, 400, response.StatusCode)
	assert.Equal(t, 400, int(responseBody["status_code"].(float64)))
	assert.True(t, strings.Contains(strings.ToLower(responseBody["message"].(string)), strings.ToLower("same")))
}
