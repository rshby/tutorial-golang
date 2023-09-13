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

// skenario 1 -> test insert sub district success
func TestInsertSubDistrictSuccess(t *testing.T) {
	// siapkan request body
	input := map[string]any{
		"name":        "Ragunan",
		"zip_code":    "12550",
		"district_id": 1,
	}

	inputJSON, _ := json.Marshal(&input)
	requestBody := strings.NewReader(string(inputJSON))

	// buat request dan recorder
	request := httptest.NewRequest(http.MethodPost, "http://localhost:5000/api/subdistrict", requestBody)
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

// skenario 2 -> test insert sub district failed bad request
func TestInsertSubDistrictFailedBadRequest(t *testing.T) {
	// siapkan request body
	input := map[string]any{
		"name":        "",
		"zip_code":    "",
		"district_id": 1,
	}

	inputJSON, _ := json.Marshal(&input)
	requestBody := strings.NewReader(string(inputJSON))

	// buat request dan recorder
	request := httptest.NewRequest(http.MethodPost, "http://localhost:5000/api/subdistrict", requestBody)
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

// skenario 3 -> test insert subdistrict failed internal server error
func TestInsertSubDistrictFailedInternalServer(t *testing.T) {
	// siapkan request body
	input := map[string]any{
		"name":        "Ragunan",
		"zip_code":    "12550",
		"district_id": 99999,
	}

	inputJSON, _ := json.Marshal(&input)
	requestBody := strings.NewReader(string(inputJSON))

	// buat request dan recorder
	request := httptest.NewRequest(http.MethodPost, "http://localhost:5000/api/subdistrict", requestBody)
	request.Header.Add("content-type", "application/json")

	recorder := httptest.NewRecorder()

	handler_test.ServeHTTP(recorder, request)

	// simpan hasil response
	response := recorder.Result()
	body, _ := io.ReadAll(response.Body)
	responseBody := map[string]any{}
	json.Unmarshal(body, &responseBody)

	// cek hasil test
	assert.Equal(t, 500, response.StatusCode)
	assert.Equal(t, 500, int(responseBody["status_code"].(float64)))
	assert.Equal(t, "internal server error", responseBody["status"].(string))
}

// skenario 4 -> test update subdistrict success
func TestUpdateSubDistrictSuccess(t *testing.T) {
	// siapkan request body
	input := map[string]any{
		"id":          1,
		"name":        "Ragunan",
		"zip_code":    "12550",
		"district_id": 1,
	}

	inputJSON, _ := json.Marshal(&input)
	requestBody := strings.NewReader(string(inputJSON))

	// buat request dan recorder
	request := httptest.NewRequest(http.MethodPut, "http://localhost:5000/api/subdistrict", requestBody)
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

// skenario 5 -> test update subdistrict failed not found
func TestUpdateSubDistrictFailedNotFound(t *testing.T) {
	// siapkan request body
	input := map[string]any{
		"id":          999,
		"name":        "Ragunan",
		"zip_code":    "12550",
		"district_id": 1,
	}

	inputJSON, _ := json.Marshal(&input)
	requestBody := strings.NewReader(string(inputJSON))

	// buat request dan recorder
	request := httptest.NewRequest(http.MethodPut, "http://localhost:5000/api/subdistrict", requestBody)
	request.Header.Add("content-type", "application/json")

	recorder := httptest.NewRecorder()

	handler_test.ServeHTTP(recorder, request)

	// simpan hasil response
	response := recorder.Result()
	body, _ := io.ReadAll(response.Body)
	responseBody := map[string]any{}
	json.Unmarshal(body, &responseBody)

	// cek hasil test
	assert.Equal(t, 404, response.StatusCode)
	assert.Equal(t, 404, int(responseBody["status_code"].(float64)))
	assert.Equal(t, "not found", responseBody["status"].(string))
}

// skenario 6 -> test update subdistrict failed internal server error
func TestUpdateSubDistrictFailedInternalServer(t *testing.T) {
	// siapkan request body
	input := map[string]any{
		"id":          1,
		"name":        "Ragunan",
		"zip_code":    "12550",
		"district_id": 99999,
	}

	inputJSON, _ := json.Marshal(&input)
	requestBody := strings.NewReader(string(inputJSON))

	// buat request dan recorder
	request := httptest.NewRequest(http.MethodPut, "http://localhost:5000/api/subdistrict", requestBody)
	request.Header.Add("content-type", "application/json")

	recorder := httptest.NewRecorder()

	handler_test.ServeHTTP(recorder, request)

	// simpan hasil response
	response := recorder.Result()
	body, _ := io.ReadAll(response.Body)
	responseBody := map[string]any{}
	json.Unmarshal(body, &responseBody)

	// cek hasil test
	assert.Equal(t, 500, response.StatusCode)
	assert.Equal(t, 500, int(responseBody["status_code"].(float64)))
	assert.Equal(t, "internal server error", responseBody["status"].(string))
}

// skenario 7 -> test update subdistrict failed bad request gagal validasi
func TestUpdateSubDistrictFailedBadRequest(t *testing.T) {
	// buat request body
	input := map[string]any{
		"id":          1,
		"name":        "",
		"zip_code":    "",
		"district_id": 1,
	}

	inputJSON, _ := json.Marshal(&input)
	requestBody := strings.NewReader(string(inputJSON))

	// buat request dan recorder
	request := httptest.NewRequest(http.MethodPut, "http://localhost:5000/api/subdistrict", requestBody)
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

// skenario 8 -> test delete data subdistrict success
func TestDeleteSubDistrictSuccess(t *testing.T) {
	// insert data ke database menggunakan repository
	sdRepo := repository.NewSubDistrictRepository(NewDatabaseTesting())
	subdistrict := entity.SubDistrict{
		Name:       "jati padang",
		ZipCode:    "12551",
		DistrictId: 1,
	}
	sdInsert, _ := sdRepo.Insert(context.TODO(), &subdistrict)

	// buat request dan recorder
	request := httptest.NewRequest(http.MethodDelete, "http://localhost:5000/api/subdistrict/"+strconv.Itoa(sdInsert.Id), nil)
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

// skenario 9 -> test delete data subdistrict failed not found
func TestDeleteSubDistrictFailedNotFound(t *testing.T) {
	// tambahkan request dan recorder
	request := httptest.NewRequest(http.MethodDelete, "http://localhost:5000/api/subdistrict/999", nil)
	recorder := httptest.NewRecorder()

	handler_test.ServeHTTP(recorder, request)

	// simpan hasil response
	response := recorder.Result()
	body, _ := io.ReadAll(response.Body)
	responseBody := map[string]any{}
	json.Unmarshal(body, &responseBody)

	// cek hasil test
	assert.Equal(t, 404, response.StatusCode)
	assert.Equal(t, 404, int(responseBody["status_code"].(float64)))
	assert.Equal(t, "not found", responseBody["status"].(string))
}

// skenario 10 -> test get all success
func TestGetAllSubDistrictSuccess(t *testing.T) {
	// buat request dan recorder
	request := httptest.NewRequest(http.MethodGet, "http://localhost:5000/api/subdistricts", nil)
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

// skenario 11 -> test get all failed not found
func TestGetAllSubDistrictsFailed(t *testing.T) {
	// truncate database
	db := NewDatabaseTesting()
	sqlDB, _ := db.DB()
	sqlDB.Exec("SET FOREIGN_KEY_CHECKS = 0")
	sqlDB.Exec("TRUNCATE sub_districts")
	sqlDB.Exec("SET FOREIGN_KEY_CHECKS = 1")

	// buat request dan recorder
	request := httptest.NewRequest(http.MethodGet, "http://localhost:5000/api/subdistricts", nil)
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

// skenario 12 -> test get data subdistrict by Id success
func TestGetSubDistrictByIdSuccess(t *testing.T) {
	// insert data ke database menggunakan repository
	subDistrictRepo := repository.NewSubDistrictRepository(NewDatabaseTesting())
	subdistrict := entity.SubDistrict{
		Name:       "ragunan",
		ZipCode:    "12550",
		DistrictId: 1,
	}
	subDistrictInsert, _ := subDistrictRepo.Insert(context.TODO(), &subdistrict)

	// buat request dan recorder
	request := httptest.NewRequest(http.MethodGet, "http://localhost:5000/api/subdistrict/"+strconv.Itoa(subDistrictInsert.Id), nil)
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

// skenario 13 -> test get data subdistrict by Id Failed not found
func TestGetSubDistrictByIdFailed(t *testing.T) {
	// buat request dan recorder
	request := httptest.NewRequest(http.MethodGet, "http://localhost:5000/api/subdistrict/999", nil)
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
	assert.Equal(t, "record not found", responseBody["message"].(string))
}

// skenario 14 -> test get all subdistricts by district_id success
func TestGetAllSubDistrictsByDistrictIdSuccess(t *testing.T) {
	// buat request dan response
	request := httptest.NewRequest(http.MethodGet, "http://localhost:5000/api/district/1/subdistricts", nil)
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

// skenario 15 -> test get all subdistricts by district_id failed
func TestGetAllSubDistrictsByDistrictIdFailed(t *testing.T) {
	// buat request dan repsponse
	request := httptest.NewRequest(http.MethodGet, "http://localhost:5000/api/district/999/subdistricts", nil)
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
