package server

import (
	"net/http"

	"github.com/gorilla/mux"
)

func NewServer(handler *mux.Router) *http.Server {
	return &http.Server{
		Addr:    "localhost:5000",
		Handler: handler,
	}
}
