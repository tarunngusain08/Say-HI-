package auth

import (
	"github.com/dgrijalva/jwt-go"
	"time"
)

type JWT struct {
	secretKey string
}

func NewJWT(secretKey string) *JWT {
	return &JWT{secretKey: secretKey}
}

func (j *JWT) GenerateJWT(username, password string) (string, error) {
	// Replace with your secret key
	secretKey := []byte(j.secretKey)

	// Create the token
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"username": username,
		"password": password,
		"exp":      time.Now().Add(time.Hour * 1).Unix(), // Token expires in 1 hour
	})

	// Sign the token with the secret key
	tokenString, err := token.SignedString(secretKey)
	if err != nil {
		return "", err
	}

	return tokenString, nil
}
