package helper

import "net/http"

func NotAllowedMethod() http.Handler {
	var notallowed http.HandlerFunc = func(w http.ResponseWriter, r *http.Request) {
		ResponseError(w, http.StatusMethodNotAllowed, "method not allowed", "http method yang digunakan salah")
	}

	return notallowed
}
