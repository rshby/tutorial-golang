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

// skenario 1 -> test get user by email success
func TestGetUserByEmailSuccess(t *testing.T) {
	// create request dan recorder
	request := httptest.NewRequest(http.MethodGet, "http://localhost:5000/api/user?email=reoshby@gmail.com", nil)
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

// skenario 2 -> test get user by email failed not found
func TestGetUserByEmailFailedNotFound(t *testing.T) {
	// create request dan recorder
	request := httptest.NewRequest(http.MethodGet, "http://localhost:5000/api/user?email=qwert@gmail.com", nil)
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
	assert.Equal(t, 404, response.StatusCode)
	assert.Equal(t, 404, int(responseBody["status_code"].(float64)))
	assert.True(t, strings.Contains(strings.ToLower(responseBody["message"].(string)), strings.ToLower("not found")))
}
