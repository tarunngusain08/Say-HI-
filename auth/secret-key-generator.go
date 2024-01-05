package auth

import (
	"crypto/rand"
	"encoding/base64"
)

func generateRandomSecretKey() (string, error) {
	key := make([]byte, Length)
	_, err := rand.Read(key)
	if err != nil {
		return "", err
	}
	return base64.URLEncoding.EncodeToString(key), nil
}
