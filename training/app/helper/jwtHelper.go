package helper

import (
	"github.com/golang-jwt/jwt/v5"
	"os"
	"time"
)

func CreateToken(email string, id int) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, &jwt.MapClaims{
		"sub":   id,
		"email": email,
		"exp":   time.Now().Add(5 * time.Hour).Unix(),
	})

	secretKey := os.Getenv("SECRET_KEY")
	tokenString, err := token.SignedString([]byte(secretKey))
	if err != nil {
		return "", err
	}

	// success generate token string
	return tokenString, nil
}
