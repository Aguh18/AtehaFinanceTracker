package utils

import (
	"log"
	"os"

	"github.com/golang-jwt/jwt"
	"github.com/joho/godotenv"
)

func GenerateToken(claims *jwt.MapClaims) (string, error) {
	errenv := godotenv.Load()
	if errenv != nil {
		log.Fatal("Error loading .env file")
	}
	SecretKey := os.Getenv("SecretKey")

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	tokenString, err := token.SignedString([]byte(SecretKey))
	if err != nil {
		return "", err
	}
	return tokenString, nil

}
