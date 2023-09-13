package helper

import (
	"encoding/json"
	"lms-mux/model/web"
	"net/http"
)

func ResponseError(w http.ResponseWriter, statusCode int, status string, message string) {
	response := web.ResponseJSON{
		StatusCode: statusCode,
		Status:     status,
		Message:    message,
	}

	w.Header().Add("content-type", "application/json")
	w.WriteHeader(statusCode)

	_ = json.NewEncoder(w).Encode(&response)
}

func ResponseSuccess(w http.ResponseWriter, statusCode int, status string, message string, data any) {
	response := web.ResponseJSON{
		StatusCode: statusCode,
		Status:     status,
		Message:    message,
		Data:       data,
	}

	w.Header().Add("content-type", "application/json")
	w.WriteHeader(statusCode)

	_ = json.NewEncoder(w).Encode(&response)
}
