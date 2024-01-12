package auth

import (
	"Say-Hi/config"
	"github.com/dgrijalva/jwt-go"
	"time"
)

type JWT struct{}

func NewJWT() *JWT {
	return &JWT{}
}

func (j *JWT) GenerateJWT(username, password string) (string, error) {
	// Replace with your secret key
	secretKey := []byte(config.Config.SecretKey)

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
