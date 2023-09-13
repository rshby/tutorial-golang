package middleware

import (
	"cms/model/dto"
	"cms/service"
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/gorilla/mux"
)

// create object loggermiddlerare
type LoggerMiddlewareObj struct {
	LoggerService *service.LoggerService
	Next          *mux.Router
}

// function provider
func NewLoggerMiddleware(service *service.LoggerService, next *mux.Router) *LoggerMiddlewareObj {
	return &LoggerMiddlewareObj{
		LoggerService: service,
		Next:          next,
	}
}

// method
func (l *LoggerMiddlewareObj) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	log.Printf("ini middleware logger")
	newResponseWriter := NewCustomResponseWriter(w)
	start := time.Now()

	// lanjutkan ke router selanjutnya
	l.Next.ServeHTTP(newResponseWriter, r)

	// proses insert data logger ke database
	loggerInput := dto.LoggerInsertRequest{
		IpAddress:  r.RemoteAddr,
		UrlPath:    r.URL.String(),
		Method:     r.Method,
		Duration:   fmt.Sprintf("%v\n", time.Since(start)),
		StatusCode: newResponseWriter.GetStatusCode(),
		Status:     newResponseWriter.Status(),
	}

	w.Header().Add("hasil_middleware", "ok")

	// proses insert ke Database Logger
	_, _ = l.LoggerService.Insert(r.Context(), &loggerInput)
}
