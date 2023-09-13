package middleware

import (
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
	"net/http"
	"os"
	"strings"
	"training/app/model/dto"
)

type AuthMiddleware struct {
}

// create function provider
func NewAuthMiddleware() *AuthMiddleware {
	return &AuthMiddleware{}
}

func (a *AuthMiddleware) Auth() gin.HandlerFunc {
	return func(c *gin.Context) {
		// het token from headers
		tokenHeader := strings.Split(c.GetHeader("Authorization"), " ")
		tokenString := tokenHeader[len(tokenHeader)-1]

		// decode token to claims
		_, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
			return []byte(os.Getenv("SECRET_KEY")), nil
		})

		if err != nil {
			c.JSON(http.StatusUnauthorized, &dto.ApiMessage{
				StatusCode: http.StatusUnauthorized,
				Status:     "unauthorized",
				Message:    err.Error(),
			})
			c.Abort()
			return
		}

		// lolos authorization
		c.Next()
	}
}
