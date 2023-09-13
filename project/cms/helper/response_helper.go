package helper

import (
	"cms/model/dto"
	"encoding/json"
	"net/http"
)

// function to mapping status
func TranslateStatus(statusCode int) string {
	status := ""
	if statusCode == http.StatusOK {
		status = "ok"
	} else if statusCode == http.StatusBadRequest {
		status = "bad request"
	} else if statusCode == http.StatusNotFound {
		status = "not found"
	} else if statusCode == http.StatusUnauthorized {
		status = "unauthorized"
	} else if statusCode == http.StatusForbidden {
		status = "forbidden"
	} else if statusCode == http.StatusMethodNotAllowed {
		status = "method not allowed"
	} else {
		status = "internal server error"
	}

	return status
}

// function Response Error
func ResponseError(w http.ResponseWriter, statusCode int, message string) {
	// create object response API
	response := dto.ApiResponse{
		StatusCode: statusCode,
		Status:     TranslateStatus(statusCode),
		Message:    message,
	}

	w.Header().Add("content-type", "application/json")
	w.WriteHeader(statusCode)

	_ = json.NewEncoder(w).Encode(&response)
}

// function Response Success
func ResponseSuccess(w http.ResponseWriter, message string, data any) {
	// create object response API success
	response := dto.ApiResponse{
		StatusCode: http.StatusOK,
		Status:     TranslateStatus(http.StatusOK),
		Message:    message,
		Data:       data,
	}

	w.Header().Add("content-type", "application/json")
	w.WriteHeader(http.StatusOK)

	_ = json.NewEncoder(w).Encode(&response)
}
