package helper

import (
	"crypto/rand"
	"math/big"
)

// create function to generate otp 6 digit
func GenerateOTP(length int) (string, error) {
	seed := "0123456789"
	byteSlice := make([]byte, length)

	for i := 0; i < length; i++ {
		max := big.NewInt(int64(len(seed)))
		num, err := rand.Int(rand.Reader, max)
		if err != nil {
			return "", err
		}

		byteSlice[i] = seed[num.Int64()]
	}

	return string(byteSlice), nil
}
