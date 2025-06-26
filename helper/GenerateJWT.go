package helper

import (
	"encoding/hex"
	"fmt"
	"github.com/golang-jwt/jwt/v5"
	"time"
)

func GenerateJWT(secretKey, subject string, minutes int) (string, error) {
	secretKeyBytes, err := hex.DecodeString(secretKey)
	if err != nil {
		return "", err
	}

	claims := jwt.MapClaims{
		"sub": subject,
		"iat": time.Now().Unix(),
		"exp": time.Now().Add(time.Minute * time.Duration(minutes)).Unix(),
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString(secretKeyBytes)
}

func VerifyJWT(secretKey, refreshToken string) (string, error) {
	secretKeyBytes, err := hex.DecodeString(secretKey)
	if err != nil {
		return "", err
	}

	token, err := jwt.Parse(refreshToken, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		return secretKeyBytes, nil
	})
	if err != nil || !token.Valid {
		return "", err
	}
	claims, ok := token.Claims.(jwt.MapClaims)
	if !ok {
		return "", err
	}
	userID, ok := claims["sub"].(string)
	if !ok || userID == "" {
		return "", err
	}

	return userID, err
}
