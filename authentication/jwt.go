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
	err := godotenv.Load()
	if err != nil {
		return "", fmt.Errorf("[FAIL]: could not load environment variables: %w", err)
	}

	var secretKey = []byte(os.Getenv("JWT_SECRET_KEY"))
	log.Println(secretKey)

	claims := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"sub":    userID,
		"issuer": "go-auth-server",
		"aud":    "user",
		"exp":    time.Now().Add(time.Hour).Unix(),
		"iat":    time.Now().Unix(),
	})

	log.Printf("Token claims added: %+v\n", claims)

	tokenString, err := claims.SignedString(secretKey)
	if err != nil {
		return "", fmt.Errorf("[FAIL]: could not sign token: %w", err)
	}

	return tokenString, nil
}
