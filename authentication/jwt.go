package authentication

import (
	"fmt"
	"log"
	"os"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"github.com/google/uuid"
	"github.com/joho/godotenv"
)

func CreateJWToken(userID uuid.UUID) (string, error) {
	// Load environment variables from .env file in development
	if env := os.Getenv("ENVIRONMENT"); env != "production" {
		err := godotenv.Load()
		if err != nil {
			return "", fmt.Errorf("[FAIL]: could not load environment variables: %w", err)
		}
	}

	var secretKey = []byte(os.Getenv("JWT_SECRET_KEY"))

	claims := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"sub":    userID,
		"issuer": "go-auth-server",
		"aud":    "user",
		"iat":    time.Now().Unix(),
		"exp":    time.Now().Add(time.Hour).Unix(),
	})

	log.Printf("[SUCCESS]: token claims added: %+v\n", claims)

	tokenString, err := claims.SignedString(secretKey)
	if err != nil {
		return "", fmt.Errorf("[FAIL]: could not sign token: %w", err)
	}

	return tokenString, nil
}

func VerifyToken(tokenString string) (*jwt.Token, error) {
	// Load environment variables from .env file in development
	if env := os.Getenv("ENVIRONMENT"); env != "production" {
		err := godotenv.Load()
		if err != nil {
			return nil, fmt.Errorf("[FAIL]: could not load environment variables: %w", err)
		}
	}

	var secretKey = []byte(os.Getenv("JWT_SECRET_KEY"))

	// Verify token
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		return secretKey, nil
	})

	if err != nil {
		return nil, err
	}

	if !token.Valid {
		return nil, fmt.Errorf("[FAIL]: invalid token sent")
	}

	return token, nil
}
