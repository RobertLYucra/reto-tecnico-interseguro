package middleware

import (
	"fmt"
	"os"
	"strings"

	utils "go-matrix-processor/internal/shared/response"

	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v5"
)

// Protected protege las rutas pasando solo JWTs válidos
func Protected() fiber.Handler {
	return func(c *fiber.Ctx) error {
		authHeader := c.Get("Authorization")
		if authHeader == "" {
			return c.Status(401).JSON(utils.ErrorResponse("Falta cabecera de autorización", ""))
		}

		parts := strings.Split(authHeader, " ")
		if len(parts) != 2 || parts[0] != "Bearer" {
			return c.Status(401).JSON(utils.ErrorResponse("Formato de token inválido", "Use: Bearer <token>"))
		}

		tokenString := parts[1]
		secret := os.Getenv("JWT_SECRET")
		if secret == "" {
			secret = "supersecretkey"
		}

		token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
			if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
				return nil, fmt.Errorf("método de firma inesperado: %v", token.Header["alg"])
			}
			return []byte(secret), nil
		})

		if err != nil || !token.Valid {
			return c.Status(401).JSON(utils.ErrorResponse("Token inválido o expirado", err.Error()))
		}

		// Guardar usuario en locals si es necesario
		// claims := token.Claims.(jwt.MapClaims)
		// c.Locals("user", claims["username"])

		return c.Next()
	}
}
