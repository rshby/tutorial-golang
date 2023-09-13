package middleware

import (
	"bytes"
	"encoding/json"
	"io"
	"net/http"
	"sync"
	"time"

	"github.com/golang-jwt/jwt"
)

type AuthMiddleware http.Handler

func NewAuthMiddleware(nextHandler http.Handler) AuthMiddleware {
	var a http.HandlerFunc = func(w http.ResponseWriter, r *http.Request) {
		cookie, err := r.Cookie("jwt_token")
		if err != nil {
			response := map[string]any{
				"code":    http.StatusUnauthorized,
				"status":  "unauthorized",
				"message": "anda belum login",
			}

			w.Header().Add("content-type", "application/json")
			w.WriteHeader(http.StatusUnauthorized)

			_ = json.NewEncoder(w).Encode(&response)
			return
		}

		tokenString := cookie.Value
		claims := jwt.StandardClaims{}

		token, _ := jwt.ParseWithClaims(tokenString, claims, func(t *jwt.Token) (interface{}, error) {
			return []byte("123456"), nil
		})

		// ambil body request
		request := map[string]any{}
		err = json.NewDecoder(r.Body).Decode(&request)
		if err != nil {
			panic(err.Error())
		}

		bodyBytes, _ := json.Marshal(&request)
		r.Body.Close()
		r.Body = io.NopCloser(bytes.NewBuffer(bodyBytes))

		if time.Now().Local().After(time.UnixMicro(claims.ExpiresAt * 1000000)) {
			// hapus cookie
			wg := sync.WaitGroup{}
			wg.Add(1)
			go func() {
				defer wg.Done()
				http.SetCookie(w, &http.Cookie{
					Name:    "jwt_token",
					Value:   "",
					MaxAge:  -1,
					Expires: time.Now().Add(-100 * time.Hour),
				})
			}()

			response := map[string]any{
				"code":    http.StatusUnauthorized,
				"status":  "unauthorized",
				"message": "token expired",
			}

			w.Header().Add("content-type", "application/json")
			w.WriteHeader(http.StatusUnauthorized)

			_ = json.NewEncoder(w).Encode(&response)
			wg.Wait()
			return
		}

		if !token.Valid {
			var grp sync.WaitGroup
			grp.Add(1)
			go func() {
				defer grp.Done()
				http.SetCookie(w, &http.Cookie{
					Name:    "jwt_token",
					Value:   "",
					MaxAge:  -1,
					Expires: time.Now().Add(-100 * time.Hour),
				})
			}()

			response := map[string]any{
				"code":    http.StatusUnauthorized,
				"status":  "unauthorized",
				"message": "token tidak valid",
			}

			w.Header().Add("content-type", "application/json")
			w.WriteHeader(http.StatusUnauthorized)

			_ = json.NewEncoder(w).Encode(&response)
			grp.Wait()
			return
		}

		// jika lolos semua pengecekan token
		nextHandler.ServeHTTP(w, r)
	}

	return AuthMiddleware(a)
}
