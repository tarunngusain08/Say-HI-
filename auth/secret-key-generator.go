package auth

import (
	"Say-Hi/config"
	"crypto/rand"
	"encoding/base64"
)

func generateRandomSecretKey() (string, error) {
	key := make([]byte, config.Config.SecretKeyLength)
	_, err := rand.Read(key)
	if err != nil {
		return "", err
	}
	return base64.URLEncoding.EncodeToString(key), nil
}
