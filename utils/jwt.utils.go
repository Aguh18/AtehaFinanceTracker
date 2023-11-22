package utils

import (
	"fmt"
	
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

func DecodeToken(tokenString string) (jwt.MapClaims, error) {
	token, err := VerifyToken(tokenString)
	if err != nil {
		return nil, err
	}
	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {

		return claims, nil
		
	}
	return nil, err

}

func VerifyToken(tokenString string) (*jwt.Token, error) {
	errenv := godotenv.Load()
	if errenv != nil {
		log.Fatal("Error loading .env file")
	}
	token, err := jwt.Parse(tokenString, func(t *jwt.Token) (interface{}, error) {
		if _, ok := t.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", t.Header["alg"])
		}
		return []byte(os.Getenv("SecretKey")), nil
	})

	if err != nil {
		return nil, err
	}
	return token, nil

}
