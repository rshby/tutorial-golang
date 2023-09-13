package testing

import (
	"cms/model/dto"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

// skenario 1 -> create content success
func TestCreateConentSuccess(t *testing.T) {
	// create request body
	input := dto.CreateContentRequest{
		AccountId:  1,
		Title:      "Tutorial Golang Microservice",
		PictureUrl: `https://www.ideamotive.co/hubfs/Building%20a%20Microservice%20in%20Go%20Business%20Guide.png`,
		TextFill:   "lets code",
	}

	inputJSON, _ := json.Marshal(&input)
	requestBody := strings.NewReader(string(inputJSON))

	// create request dan recorder
	request := httptest.NewRequest(http.MethodPost, "http://localhost:5000/api/content", requestBody)
	request.Header.Add("authorization", fmt.Sprintf("Bearer %v", GetTokenLogin()))

	recorder := httptest.NewRecorder()

	// execute api
	routerTest.ServeHTTP(recorder, request)

	// simpan hasil execute api
	response := recorder.Result()
	body, _ := io.ReadAll(response.Body)
	responseBody := map[string]any{}
	json.Unmarshal(body, &responseBody)

	// cek hasil test
	assert.Equal(t, 200, response.StatusCode)
	assert.Equal(t, 200, int(responseBody["status_code"].(float64)))
	assert.Equal(t, "ok", responseBody["status"].(string))
}

// skenario 2 -> create content failed account_id not found
func TestCreateContentFailedAccountNotFound(t *testing.T) {
	// create request body
	input := dto.CreateContentRequest{
		AccountId:  999,
		Title:      "Tutorial Express JS",
		PictureUrl: "https://afteracademy.com/images/getting-started-with-expressjs-banner-96ecb31ec64b6504.jpeg",
		TextFill:   "javascript",
	}

	inputJSON, _ := json.Marshal(&input)
	requestBody := strings.NewReader(string(inputJSON))

	// create request dan recorder
	request := httptest.NewRequest(http.MethodPost, "http://localhost:5000/api/content", requestBody)
	request.Header.Add("authorization", fmt.Sprintf("Bearer %v", GetTokenLogin()))

	recorder := httptest.NewRecorder()

	// execute api
	routerTest.ServeHTTP(recorder, request)

	// simpan hasil execute api
	response := recorder.Result()
	body, _ := io.ReadAll(response.Body)
	responseBody := map[string]any{}
	json.Unmarshal(body, &responseBody)

	// cek hasil test
	assert.Equal(t, 500, response.StatusCode)
	assert.Equal(t, 500, int(responseBody["status_code"].(float64)))
	assert.True(t, strings.Contains(strings.ToLower(responseBody["message"].(string)), strings.ToLower("account_id")))
}
