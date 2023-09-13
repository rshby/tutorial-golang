package helper

import "net/http"

func NotFoundMethod() http.Handler {
	var notfound http.HandlerFunc = func(w http.ResponseWriter, r *http.Request) {
		ResponseError(w, http.StatusNotFound, "not found", "endpoint not exist")
	}

	return notfound
}
