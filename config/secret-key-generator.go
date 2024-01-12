package config

import (
	"crypto/rand"
	"encoding/base64"
)

func GenerateRandomSecretKey() (string, error) {
	key := make([]byte, Config.SecretKeyLength)
	_, err := rand.Read(key)
	if err != nil {
		return "", err
	}
	return base64.URLEncoding.EncodeToString(key), nil
}
