package auth

import (
	"cms/helper"
	"cms/model/auth"
	"log"
	"net/http"
	"strings"
	"time"

	"github.com/golang-jwt/jwt/v4"
	"github.com/gorilla/mux"
)

// create object auth middleware
type AuthMiddleware struct {
	Next *mux.Router
}

// function provider to create new object AuthMiddleware
func NewAuthMiddleware(next *mux.Router) *AuthMiddleware {
	return &AuthMiddleware{
		Next: next,
	}
}

func (a *AuthMiddleware) Auth(next http.Handler) http.Handler {
	return http.Handler(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		log.Println("masuk ke middlewaer auth")

		// get headers Authorization
		tokenHeaders := strings.Split(r.Header.Get("Authorization"), " ")
		token := tokenHeaders[len(tokenHeaders)-1]

		// if bearer token tidak ada
		if token == "" || token == " " || (strings.Contains(token, "Bearer")) {
			helper.ResponseError(w, http.StatusUnauthorized, "token required!")
			return
		}

		// parsing token ke bentuk claims
		claims := auth.JWTClaim{}
		tokenResult, err := jwt.ParseWithClaims(token, &claims, func(t *jwt.Token) (interface{}, error) {
			return auth.JWT_SECRET_KEY, nil
		})

		// jika ada error ketika parsing token ke bentuk claims
		if err != nil {
			helper.ResponseError(w, http.StatusUnauthorized, err.Error())
			return
		}

		// jika token expired
		if time.Now().Local().After(time.UnixMicro(claims.ClaimRegistered.ExpiresAt.Unix() * 1000000)) {
			helper.ResponseError(w, http.StatusUnauthorized, "token expired")
			return
		}

		if !tokenResult.Valid {
			helper.ResponseError(w, http.StatusUnauthorized, "token not valid")
			return
		}

		// lanjutkan ke handler berikutnya
		next.ServeHTTP(w, r)
	}))
}
