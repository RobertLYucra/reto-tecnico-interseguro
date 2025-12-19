package service_test

import (
	"testing"

	"go-matrix-processor/internal/service"
)

func TestAuthenticate_Success(t *testing.T) {
	// Arrange
	svc := service.NewAuthService()
	username := "admin"
	password := "admin"

	// Act
	token, err := svc.Authenticate(username, password)

	// Assert
	if err != nil {
		t.Fatalf("Error inesperado: %v", err)
	}
	if token == "" {
		t.Error("Se esperaba un token, se obtuvo string vacío")
	}
}

func TestAuthenticate_Failure(t *testing.T) {
	// Arrange
	svc := service.NewAuthService()
	username := "user"
	password := "wrongpassword"

	// Act
	token, err := svc.Authenticate(username, password)

	// Assert
	if err == nil {
		t.Error("Se esperaba un error de autenticación, se obtuvo nil")
	}
	if token != "" {
		t.Errorf("Se esperaba token vacío, se obtuvo %s", token)
	}
	if err.Error() != "credenciales incorrectas" {
		t.Errorf("Mensaje de error incorrecto: %v", err)
	}
}
