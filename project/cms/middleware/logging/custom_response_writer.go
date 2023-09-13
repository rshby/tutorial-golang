package middleware

import (
	"cms/helper"
	"net/http"
)

// create object CustomResponseWriter
type CustomResponseWriter struct {
	ResponseWriter http.ResponseWriter
	StatusCode     int
	StatusMessage  string
	WroteHeader    bool
}

// function provider to create new object CustomResponseWriter
func NewCustomResponseWriter(rw http.ResponseWriter) *CustomResponseWriter {
	return &CustomResponseWriter{
		ResponseWriter: rw,
	}
}

// method implement WriteHeader
func (cr *CustomResponseWriter) WriteHeader(statusCode int) {
	if cr.WroteHeader {
		return
	}

	cr.StatusCode = statusCode
	cr.ResponseWriter.WriteHeader(statusCode)
	cr.StatusMessage = helper.TranslateStatus(cr.StatusCode)
	cr.WroteHeader = true
}

// method implement Header
func (cr *CustomResponseWriter) Header() http.Header {
	return cr.ResponseWriter.Header()
}

// method implement Write
func (cr *CustomResponseWriter) Write(in []byte) (int, error) {
	return cr.ResponseWriter.Write(in)
}

// method get status_code
func (cr *CustomResponseWriter) GetStatusCode() int {
	return cr.StatusCode
}

// method get status
func (cr *CustomResponseWriter) Status() string {
	return cr.StatusMessage
}
