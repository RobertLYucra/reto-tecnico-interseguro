package handlers

import (
	"go-matrix-processor/internal/service"
	utils "go-matrix-processor/internal/shared/response"

	"github.com/gofiber/fiber/v2"
)

type AuthHandler struct {
	Service *service.AuthService
}

func NewAuthHandler(s *service.AuthService) *AuthHandler {
	return &AuthHandler{Service: s}
}

// LoginRequest define el cuerpo para el inicio de sesión
type LoginRequest struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

// Login maneja la autenticación y la generación de tokens
func (h *AuthHandler) Login(c *fiber.Ctx) error {
	var req LoginRequest
	if err := c.BodyParser(&req); err != nil {
		return c.Status(400).JSON(utils.ErrorResponse("Petición inválida", err.Error()))
	}

	token, err := h.Service.Authenticate(req.Username, req.Password)
	if err != nil {
		if err.Error() == "credenciales incorrectas" {
			return c.Status(401).JSON(utils.ErrorResponse("Credenciales incorrectas", ""))
		}
		return c.Status(500).JSON(utils.ErrorResponse("Error generando token", err.Error()))
	}

	return c.JSON(utils.SuccessResponse("Login exitoso", map[string]string{
		"token": token,
	}))
}
