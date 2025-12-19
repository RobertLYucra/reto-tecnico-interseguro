package handlers_test

import (
	"bytes"
	"encoding/json"
	"go-matrix-processor/internal/handlers"
	"go-matrix-processor/internal/service"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gofiber/fiber/v2"
)

func TestLogin_Success(t *testing.T) {
	// Setup
	app := fiber.New()
	authService := service.NewAuthService()
	handler := handlers.NewAuthHandler(authService)
	app.Post("/login", handler.Login)

	// Request
	payload := map[string]string{
		"username": "admin",
		"password": "admin",
	}
	body, _ := json.Marshal(payload)
	req := httptest.NewRequest("POST", "/login", bytes.NewBuffer(body))
	req.Header.Set("Content-Type", "application/json")

	// Execute
	resp, err := app.Test(req)
	if err != nil {
		t.Fatalf("Error al ejecutar request: %v", err)
	}

	// Assert Status
	if resp.StatusCode != http.StatusOK {
		t.Errorf("Se esperaba estado 200, se obtuvo %d", resp.StatusCode)
	}

	// Assert Token
	var result struct {
		Success bool `json:"success"`
		Data    struct {
			Token string `json:"token"`
		} `json:"data"`
	}
	if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
		t.Fatalf("Falló al decodificar respuesta: %v", err)
	}

	if !result.Success {
		t.Error("Se esperaba success true")
	}
	if result.Data.Token == "" {
		t.Error("Se esperaba un token, se obtuvo vacío")
	}
}

func TestLogin_InvalidCredentials(t *testing.T) {
	// Setup
	app := fiber.New()
	authService := service.NewAuthService()
	handler := handlers.NewAuthHandler(authService)
	app.Post("/login", handler.Login)

	// Request
	payload := map[string]string{
		"username": "admin",
		"password": "wrongpassword",
	}
	body, _ := json.Marshal(payload)
	req := httptest.NewRequest("POST", "/login", bytes.NewReader(body))
	req.Header.Set("Content-Type", "application/json")

	// Act
	resp, _ := app.Test(req)

	// Assert
	if resp.StatusCode != 401 {
		t.Errorf("Se esperaba 401, se obtuvo %d", resp.StatusCode)
	}
}

func TestLogin_BadRequest(t *testing.T) {
	// Setup
	app := fiber.New()
	authService := service.NewAuthService()
	handler := handlers.NewAuthHandler(authService)
	app.Post("/login", handler.Login)

	// Request (Invalid JSON)
	req := httptest.NewRequest("POST", "/login", bytes.NewReader([]byte("invalid-json")))
	req.Header.Set("Content-Type", "application/json")

	// Act
	resp, _ := app.Test(req)

	// Assert
	if resp.StatusCode != 400 {
		t.Errorf("Se esperaba 400, se obtuvo %d", resp.StatusCode)
	}
}
