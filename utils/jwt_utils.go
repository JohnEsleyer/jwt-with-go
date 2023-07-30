package utils

import (
	"time"

	"github.com/dgrijalva/jwt-go"
)

func GenerateToken(username string, secretKey []byte, expiration time.Duration) string {
	token := jwt.New(jwt.SigningMethodHS256)

	// Set claims (payload)
	claims := token.Claims.(jwt.MapClaims)
	claims["username"] = username
	claims["exp"] = time.Now().Add(expiration).Unix()

	// Sign the token with the secret key
	tokenString, _ := token.SignedString(secretKey)
	return tokenString
}

func ParseToken(tokenString string, secretKey []byte) (*jwt.Token, error) {
	return jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		return secretKey, nil
	})
}
