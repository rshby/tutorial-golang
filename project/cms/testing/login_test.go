package testing

import (
	"encoding/json"
	"io"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

// function to get token
func GetTokenLogin() string {
	input := map[string]any{
		"username": "rshby",
		"password": "111",
	}

	inputJSON, _ := json.Marshal(&input)
	requestBody := strings.NewReader(string(inputJSON))

	request := httptest.NewRequest(http.MethodPost, "http://localhost:5000/api/login", requestBody)
	request.Header.Add("content-type", "application/json")

	recorder := httptest.NewRecorder()

	// execute api
	routerTest.ServeHTTP(recorder, request)

	// simpan hasil execute
	response := recorder.Result()
	body, _ := io.ReadAll(response.Body)
	responseBody := map[string]any{}

	json.Unmarshal(body, &responseBody)

	// get token
	token := string(responseBody["data"].(map[string]any)["token"].(string))
	return token
}

// skenario 1 -> test login success
func TestLoginSuccess(t *testing.T) {
	// buat request body
	input := map[string]string{
		"username": "reoshby@gmail.com",
		"password": "111",
	}

	// create json string bytes
	inputJSON, _ := json.Marshal(&input)
	requestbody := strings.NewReader(string(inputJSON))

	// buat request dan recorder
	request := httptest.NewRequest(http.MethodPost, "http://localhost:5000/api/login", requestbody)
	request.Header.Add("content-type", "application/json")

	recorder := httptest.NewRecorder()

	routerTest.ServeHTTP(recorder, request)

	// simpan hasil ke response
	response := recorder.Result()
	body, _ := io.ReadAll(response.Body)
	responseBody := map[string]any{}

	// ubah json bytes ke bentuk object map
	json.Unmarshal(body, &responseBody)

	// cek hasil test
	assert.Equal(t, 200, response.StatusCode)
	assert.Equal(t, 200, int(responseBody["status_code"].(float64)))
	assert.Equal(t, strings.ToLower("ok"), strings.ToLower(responseBody["status"].(string)))
}
