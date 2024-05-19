package utils

import (
	"time"

	"github.com/golang-jwt/jwt/v5"
)

func Sign(data map[string]interface{}, SecretKeyEnvName string, ExpiredAt time.Duration) (string, error) {
	expiredAt := time.Now().Add(time.Duration(time.Minute) * ExpiredAt).Unix()

	// GET KEY FROM ENV
	jwtSecretKey := SecretKeyEnvName

	claims := jwt.MapClaims{}
	claims["exp"] = expiredAt
	claims["authorization"] = true

	for k, v := range data {
		claims[k] = v
	}

	to := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	accessToken, err := to.SignedString([]byte(jwtSecretKey))
	if err != nil {
		return accessToken, err
	}

	return accessToken, nil
}
