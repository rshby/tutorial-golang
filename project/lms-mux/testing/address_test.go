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

	"github.com/stretchr/testify/assert"
)

// skenario 1 -> test insert address success
func TestInsertAddressSuccess(t *testing.T) {
	// buat request body
	input := map[string]any{
		"street":         "Jl. Swadaya No.57",
		"subdistrict_id": 1,
	}
	inputJSON, _ := json.Marshal(&input)
	requestBody := strings.NewReader(string(inputJSON))

	// buat request dan recorder
	request := httptest.NewRequest(http.MethodPost, "http://localhost:5000/api/address", requestBody)
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

// skenario 2 -> test insert address failed bad request
func TestInsertAddressFailedBadRequest(t *testing.T) {
	// buat request body
	input := map[string]any{
		"street":         "",
		"subdistrict_id": 1,
	}
	inputJSON, _ := json.Marshal(&input)
	requestBody := strings.NewReader(string(inputJSON))

	// buat request dan recorder
	request := httptest.NewRequest(http.MethodPost, "http://localhost:5000/api/address", requestBody)
	request.Header.Add("content-type", "application/json")

	recorder := httptest.NewRecorder()

	handler_test.ServeHTTP(recorder, request)

	// simpan hasil response
	response := recorder.Result()
	body, _ := io.ReadAll(response.Body)
	responseBody := map[string]any{}
	json.Unmarshal(body, &responseBody)

	// cek hasil test
	assert.Equal(t, 400, response.StatusCode)
	assert.Equal(t, 400, int(responseBody["status_code"].(float64)))
	assert.Equal(t, "bad request", responseBody["status"].(string))
}

// skenario 3 -> test insert adress failed internal server error
func TestInsertAddressFailedServerError(t *testing.T) {
	// buat request body
	input := map[string]any{
		"street":         "jl.Rajawali",
		"subdistrict_id": 99,
	}
	inputJSON, _ := json.Marshal(&input)
	requestBody := strings.NewReader(string(inputJSON))

	// buat request dan recorder
	request := httptest.NewRequest(http.MethodPost, "http://localhost:5000/api/address", requestBody)
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

// skenario 4 -> test update address success
func TestUpdateAddressSuccess(t *testing.T) {
	// siapkan request body
	input := map[string]any{
		"id":             1,
		"street":         "jl. Swadaya No.57",
		"subdistrict_id": 1,
	}
	inputJSON, _ := json.Marshal(&input)
	requestBody := strings.NewReader(string(inputJSON))

	// buat request dan recorder
	request := httptest.NewRequest(http.MethodPut, "http://localhost:5000/api/address", requestBody)
	request.Header.Add("content-type", "application/json")

	recorder := httptest.NewRecorder()

	handler_test.ServeHTTP(recorder, request)

	// simpan hasil response
	response := recorder.Result()
	body, _ := io.ReadAll(response.Body)
	responseBody := map[string]any{}
	json.Unmarshal(body, &responseBody)

	// cek hasil test
	assert.Equal(t, 200, response.StatusCode)
	assert.Equal(t, 200, int(responseBody["status_code"].(float64)))
	assert.Equal(t, "ok", responseBody["status"].(string))
}

// skenario 5 -> test update address failed not found
func TestUpdateAddressFailedNotFound(t *testing.T) {
	// siapkan request body
	input := map[string]any{
		"id":             9999,
		"street":         "jl. Jalan",
		"subdistrict_id": 1,
	}
	inputJSON, _ := json.Marshal(&input)
	requestBody := strings.NewReader(string(inputJSON))

	// buat request dan recorder
	request := httptest.NewRequest(http.MethodPut, "http://localhost:5000/api/address", requestBody)
	request.Header.Add("content-type", "application/json")

	recorder := httptest.NewRecorder()

	handler_test.ServeHTTP(recorder, request)

	// simpan hasil repsonse
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

// skenario 6 -> test update address failed bad request
func TestUpdateAddressFailedBadRequest(t *testing.T) {
	// siapkan request body
	input := map[string]any{
		"id":             1,
		"street":         "",
		"subdistrict_id": 1,
	}
	inputJSON, _ := json.Marshal(&input)
	requestBody := strings.NewReader(string(inputJSON))

	// buat request dan recorder
	request := httptest.NewRequest(http.MethodPut, "http://localhost:5000/api/address", requestBody)
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

// skenario 7 -> test update address failed internal server error
func TestUpdateAddressFailedInternalServer(t *testing.T) {
	// siapkan request body
	input := map[string]any{
		"id":             1,
		"street":         "Jl. Jalan",
		"subdistrict_id": 99999,
	}
	inputJSON, _ := json.Marshal(&input)
	requestBody := strings.NewReader(string(inputJSON))

	// buat request dan recorder
	request := httptest.NewRequest(http.MethodPut, "http://localhost:5000/api/address", requestBody)
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

// skenario 8 -> test delete success
func TestDeleteAddressSuccess(t *testing.T) {
	// tambahkan data menggunakan repository
	addressRepo := repository.NewGeneralRepository(NewDatabaseTesting(), &entity.Address{})
	address := entity.Address{
		Street:        "jalan senopati",
		SubDistrictId: 1,
	}
	addressInsert, _ := addressRepo.Insert(context.Background(), &address)

	// buat request dan recorder
	request := httptest.NewRequest(http.MethodDelete, "http://localhost:5000/api/address/"+strconv.Itoa(addressInsert.Id), nil)
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

// skenario 9 -> test delete failed not found
func TestDeleteAddressFailedNotFound(t *testing.T) {
	// buat request dan recorder
	request := httptest.NewRequest(http.MethodDelete, "http://localhost:5000/api/address/9999", nil)
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

// skenario 10 -> test get all data addresses success
func TestGetAllAddressesSuccess(t *testing.T) {
	// buat request dan recorder
	request := httptest.NewRequest(http.MethodGet, "http://localhost:5000/api/addresses", nil)
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

// skenario 11 -> test get all data addresses failed not found
func TestGetAllAddressesFailed(t *testing.T) {
	// truncate
	db := NewDatabaseTesting()
	sqlDB, _ := db.DB()
	sqlDB.Exec("SET FOREIGN_KEY_CHECKS = 0")
	sqlDB.Exec("TRUNCATE addresses")
	sqlDB.Exec("SET FOREIGN_KEY_CHECKS = 1")

	// buat request dan recorder
	request := httptest.NewRequest(http.MethodGet, "http://localhost:5000/api/addresses", nil)
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

// skenario 12 -> test get by Id success
func TestGetAddressByIdSuccess(t *testing.T) {
	// insert data menggunakan repository
	addressRepo := repository.NewGeneralRepository(NewDatabaseTesting(), &entity.Address{})
	address := entity.Address{
		Street:        "jl. Remaja",
		SubDistrictId: 1,
	}
	addressInsert, _ := addressRepo.Insert(context.Background(), &address)

	// buat request dan recorder
	request := httptest.NewRequest(http.MethodGet, "http://localhost:5000/api/address/"+strconv.Itoa(addressInsert.Id), nil)
	recorder := httptest.NewRecorder()

	handler_test.ServeHTTP(recorder, request)

	// simpan hasil test
	response := recorder.Result()
	body, _ := io.ReadAll(response.Body)
	responseBody := map[string]any{}
	json.Unmarshal(body, &responseBody)

	// cek hasil test
	assert.Equal(t, 200, response.StatusCode)
	assert.Equal(t, 200, int(responseBody["status_code"].(float64)))
	assert.Equal(t, "ok", responseBody["status"].(string))
}

// skenario 13 -> test get by Id failed
func TestGetByIdFailed(t *testing.T) {
	// buat requeest dan recorder
	request := httptest.NewRequest(http.MethodGet, "http://localhost:5000/api/address/999", nil)
	recorder := httptest.NewRecorder()

	handler_test.ServeHTTP(recorder, request)

	// simpan hasil request
	response := recorder.Result()
	body, _ := io.ReadAll(response.Body)
	responseBody := map[string]any{}
	json.Unmarshal(body, &responseBody)

	// cek hasil test
	assert.Equal(t, 404, response.StatusCode)
	assert.Equal(t, 404, int(responseBody["status_code"].(float64)))
	assert.Equal(t, "not found", responseBody["status"].(string))
}

// skenario 14 -> test get full address success
func TestGetFullAddressSuccess(t *testing.T) {
	// buat request dan recorder
	request := httptest.NewRequest(http.MethodGet, "http://localhost:5000/api/full-address/1", nil)
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

// skenario 15 -> test get full address failed not found
func TestGetAllFullAddressFailed(t *testing.T) {
	// buat request dan recorder
	request := httptest.NewRequest(http.MethodGet, "http://localhost:5000/api/full-address/9999", nil)
	recorder := httptest.NewRecorder()

	handler_test.ServeHTTP(recorder, request)

	// simpan hasil test
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
