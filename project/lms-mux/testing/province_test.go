package testing

import (
	"context"
	"encoding/json"
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

// skenario 1 -> test insert province success
func TestInsertProvinceSuccess(t *testing.T) {
	// buat requestBody
	requestInser := map[string]any{
		"name": "DKI Jakarta",
	}
	requestJSON, _ := json.Marshal(&requestInser)
	requestBody := strings.NewReader(string(requestJSON))

	// buat request dan recorder
	request := httptest.NewRequest(http.MethodPost, "http://localhost:5000/api/province", requestBody)
	request.Header.Add("content-type", "application/json")

	recorder := httptest.NewRecorder()

	handler_test.ServeHTTP(recorder, request)

	response := recorder.Result()
	body, _ := io.ReadAll(response.Body)
	responseBody := map[string]any{}
	json.Unmarshal(body, &responseBody)

	// cek hasil test
	assert.Equal(t, 200, response.StatusCode)
	assert.Equal(t, 200, int(responseBody["status_code"].(float64)))
	assert.Equal(t, "ok", responseBody["status"])
	assert.Equal(t, requestInser["name"].(string), responseBody["data"].(map[string]any)["name"].(string))
}

// skenario 2 -> test insert province failed
func TestInsertProvinceFailed(t *testing.T) {
	// siapkan request body
	input := map[string]any{
		"name": "",
	}
	inputJSON, _ := json.Marshal(&input)
	requestBody := strings.NewReader(string(inputJSON))

	// buat request dan recorder
	request := httptest.NewRequest(http.MethodPost, "http://localhost:5000/api/province", requestBody)
	request.Header.Add("content-type", "application/json")

	recorder := httptest.NewRecorder()

	handler_test.ServeHTTP(recorder, request)

	response := recorder.Result()
	body, _ := io.ReadAll(response.Body)
	responseBody := map[string]any{}
	json.Unmarshal(body, &responseBody)

	// cek hasil test
	assert.Equal(t, 400, response.StatusCode)
	assert.Equal(t, 400, int(responseBody["status_code"].(float64)))
	assert.Equal(t, "bad request", responseBody["status"].(string))
}

// skenario 3 -> test update success
func TestUpdateProvinceSuccess(t *testing.T) {
	// buat request body
	input := map[string]any{
		"id":   1,
		"name": "DKI Jakarta",
	}
	inputJSON, _ := json.Marshal(&input)
	requestBody := strings.NewReader(string(inputJSON))

	// buat request dan recorder
	request := httptest.NewRequest(http.MethodPut, "http://localhost:5000/api/province/"+strconv.Itoa(input["id"].(int)), requestBody)
	request.Header.Add("content-type", "application/json")

	recorder := httptest.NewRecorder()

	handler_test.ServeHTTP(recorder, request)

	response := recorder.Result()
	body, _ := io.ReadAll(response.Body)
	responseBody := map[string]any{}
	json.Unmarshal(body, &responseBody)

	// cek hasil test
	assert.Equal(t, 200, response.StatusCode)
	assert.Equal(t, 200, int(responseBody["status_code"].(float64)))
	assert.Equal(t, "ok", responseBody["status"].(string))
	assert.Equal(t, input["name"].(string), responseBody["data"].(map[string]any)["name"].(string))
}

// skenario 4 -> test update failed
func TestUpdateProvinceFailed(t *testing.T) {
	// siapkan request body
	input := map[string]any{
		"id":   1,
		"name": "",
	}
	inputJSON, _ := json.Marshal(&input)
	requestBody := strings.NewReader(string(inputJSON))

	// buat request dan recorder
	request := httptest.NewRequest(http.MethodPut, "http://localhost:5000/api/province/"+strconv.Itoa(input["id"].(int)), requestBody)
	request.Header.Add("content-type", "application/json")

	recorder := httptest.NewRecorder()

	handler_test.ServeHTTP(recorder, request)

	response := recorder.Result()
	body, _ := io.ReadAll(response.Body)
	responseBody := map[string]any{}
	json.Unmarshal(body, &responseBody)

	// cek hasil test
	assert.Equal(t, 400, response.StatusCode)
	assert.Equal(t, 400, int(responseBody["status_code"].(float64)))
	assert.Equal(t, "bad request", responseBody["status"].(string))
}

// skenario 5 -> test delete province success
func TestDeleteProvinceSuccess(t *testing.T) {
	// insert data ke database mengunakan repository
	repoprovince := repository.NewProvinceRepository(NewDatabaseTesting())
	province := entity.Province{
		Name: "Jawa Tengah",
	}
	provinceInsert, _ := repoprovince.Insert(context.TODO(), &province)

	// buat request dan recorder
	request := httptest.NewRequest(http.MethodDelete, "http://localhost:5000/api/province/"+strconv.Itoa(provinceInsert.Id), nil)
	recorder := httptest.NewRecorder()

	handler_test.ServeHTTP(recorder, request)

	response := recorder.Result()
	body, _ := io.ReadAll(response.Body)
	responseBody := map[string]any{}
	json.Unmarshal(body, &responseBody)

	// cek hasil test
	assert.Equal(t, 200, response.StatusCode)
	assert.Equal(t, 200, int(responseBody["status_code"].(float64)))
	assert.Equal(t, "ok", responseBody["status"].(string))
	assert.Equal(t, "success delete data province", responseBody["message"].(string))
}

// skenario 6 -> test delete province failed
func TestDeleteProvinceFailed(t *testing.T) {
	// buat request dan recorder
	request := httptest.NewRequest(http.MethodDelete, "http://localhost:5000/api/province/99999", nil)
	recorder := httptest.NewRecorder()

	handler_test.ServeHTTP(recorder, request)

	// simpan response
	response := recorder.Result()
	body, _ := io.ReadAll(response.Body)
	responseBody := map[string]any{}
	json.Unmarshal(body, &responseBody)

	// cek hasil test
	assert.Equal(t, 404, response.StatusCode)
	assert.Equal(t, 404, int(responseBody["status_code"].(float64)))
	assert.Equal(t, "not found", responseBody["status"].(string))
	assert.Equal(t, "record not found", responseBody["message"].(string))
}

// skenario 7 -> test Get all provinces success
func TestGetAllProvincesSuccess(t *testing.T) {
	// insert data province menggunakan repository
	repoProvince := repository.NewProvinceRepository(NewDatabaseTesting())
	province := entity.Province{
		Name: "Aceh",
	}
	_, _ = repoProvince.Insert(context.TODO(), &province)

	// buat request dan recorder
	request := httptest.NewRequest(http.MethodGet, "http://localhost:5000/api/provinces", nil)
	recorder := httptest.NewRecorder()

	handler_test.ServeHTTP(recorder, request)

	// simpan response
	response := recorder.Result()
	body, _ := io.ReadAll(response.Body)
	responseBody := map[string]any{}
	json.Unmarshal(body, &responseBody)

	// cek hasil test
	assert.Equal(t, 200, response.StatusCode)
	assert.Equal(t, 200, int(responseBody["status_code"].(float64)))
	assert.Equal(t, "ok", responseBody["status"].(string))
}

// skenario 8 -> test get all provinces failed
func TestGetAllProvincesFailed(t *testing.T) {
	db := NewDatabaseTesting()
	sqlDB, _ := db.DB()
	sqlDB.Exec("SET FOREIGN_KEY_CHECKS = 0")
	sqlDB.Exec("TRUNCATE provinces")
	sqlDB.Exec("SET FOREIGN_KEY_CHECKS = 1")

	// buat recorder dan request
	request := httptest.NewRequest(http.MethodGet, "http://localhost:5000/api/provinces", nil)
	recorder := httptest.NewRecorder()

	handler_test.ServeHTTP(recorder, request)

	// simpan response
	response := recorder.Result()
	body, _ := io.ReadAll(response.Body)
	responseBody := map[string]any{}
	json.Unmarshal(body, &responseBody)

	// cek hasil test
	assert.Equal(t, 404, response.StatusCode)
	assert.Equal(t, 404, int(responseBody["status_code"].(float64)))
	assert.Equal(t, "not found", responseBody["status"].(string))
}

// skenario 9 -> test get province by id success
func TestGetProvinceByIdSuccess(t *testing.T) {
	// insert data province menggunakan repository
	repoProvince := repository.NewProvinceRepository(NewDatabaseTesting())
	province := entity.Province{
		Name: "Jawa Timur",
	}
	provinceInsert, _ := repoProvince.Insert(context.TODO(), &province)

	// buat request dan recorder
	request := httptest.NewRequest(http.MethodGet, "http://localhost:5000/api/province/"+strconv.Itoa(provinceInsert.Id), nil)
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

// skenario 10 -> test get province by id failed
func TestGetProvinceByIdFailed(t *testing.T) {
	db := NewDatabaseTesting()
	sqlDB, _ := db.DB()
	sqlDB.Exec("SET FOREIGN_KEY_CHECKS = 0")
	sqlDB.Exec("TRUNCATE provinces")
	sqlDB.Exec("SET FOREIGN_KEY_CHECKS = 1")

	// buat request dan recorder
	request := httptest.NewRequest(http.MethodGet, "http://localhost:5000/api/province/99999", nil)
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
