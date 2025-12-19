package service

import (
	"errors"
	"os"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

type AuthService struct{}

func NewAuthService() *AuthService {
	return &AuthService{}
}

// Authenticate checks credentials and returns a JWT token if valid
func (s *AuthService) Authenticate(username, password string) (string, error) {
	// Mock Authentication (Check DB in production)
	if username != "admin" || password != "admin" {
		return "", errors.New("credenciales incorrectas")
	}

	// Create Token
	token := jwt.New(jwt.SigningMethodHS256)
	claims := token.Claims.(jwt.MapClaims)
	claims["username"] = username
	// 3 days expiration
	claims["exp"] = time.Now().Add(time.Hour * 72).Unix()

	// Sign Token
	secret := os.Getenv("JWT_SECRET")
	if secret == "" {
		secret = "supersecretkey" // Fallback for dev
	}

	t, err := token.SignedString([]byte(secret))
	if err != nil {
		return "", err
	}

	return t, nil
}
