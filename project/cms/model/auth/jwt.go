package auth

import "github.com/golang-jwt/jwt/v4"

// create variabel
var JWT_SECRET_KEY = []byte("sangatrahasia")

// create object Claims
type JWTClaim struct {
	Email           string `json:"email,omitempty"`
	ClaimRegistered jwt.RegisteredClaims
}

// method implement
func (c *JWTClaim) Valid() error {
	return c.ClaimRegistered.Valid()
}
