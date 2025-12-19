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

// Authenticate verifica las credenciales y retorna un token JWT si son válidas
func (s *AuthService) Authenticate(username, password string) (string, error) {
	// Autenticación simulada (Verificar BD en producción)
	if username != "admin" || password != "admin" {
		return "", errors.New("credenciales incorrectas")
	}

	// Crear Token
	token := jwt.New(jwt.SigningMethodHS256)
	claims := token.Claims.(jwt.MapClaims)
	claims["username"] = username
	// 3 días de expiración
	claims["exp"] = time.Now().Add(time.Hour * 72).Unix()

	// Firmar Token
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
