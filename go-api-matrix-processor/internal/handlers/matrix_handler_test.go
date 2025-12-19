package handlers_test

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"os"
	"testing"

	"go-matrix-processor/internal/handlers"
	"go-matrix-processor/internal/models"
	"go-matrix-processor/internal/service"

	"github.com/gofiber/fiber/v2"
)

func TestProcessMatrix_Integration(t *testing.T) {
	// 1. Mock Node.js API Server
	mockNodeServer := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// Verify request from Go (could check body for Q and R)
		if r.Method != http.MethodPost {
			t.Errorf("Se esperaba petición POST, se obtuvo %s", r.Method)
		}

		// Return success response mimicking Node.js API
		response := map[string]interface{}{
			"success": true,
			"message": "Stats calculated from mock",
			"data": map[string]interface{}{
				"max":       5,
				"total_sum": 10,
			},
		}
		json.NewEncoder(w).Encode(response)
	}))
	defer mockNodeServer.Close()

	// 2. Set Env Var to point to Mock Server
	os.Setenv("NODE_API_URL", mockNodeServer.URL)
	defer os.Unsetenv("NODE_API_URL")

	// 3. Setup Fiber App with Handler
	app := fiber.New()
	svc := service.NewMatrixService()
	handler := handlers.NewMatrixHandler(svc)
	app.Post("/process", handler.ProcessMatrix)

	// 4. Create Request
	inputData := models.InputMatrix{
		Data: models.Matrix{
			{1, 2},
			{3, 4},
		},
	}
	body, _ := json.Marshal(inputData)
	req := httptest.NewRequest("POST", "/process", bytes.NewBuffer(body))
	req.Header.Set("Content-Type", "application/json")

	// 5. Execute Request
	resp, err := app.Test(req, -1)
	if err != nil {
		t.Fatalf("La petición falló: %v", err)
	}

	// 6. Assertions
	if resp.StatusCode != http.StatusOK {
		t.Errorf("Se esperaba estado 200, se obtuvo %d", resp.StatusCode)
	}

	var result struct {
		Success bool `json:"success"`
		Data    struct {
			Max float64 `json:"max"`
		} `json:"data"`
	}
	if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
		t.Fatalf("Falló al decodificar respuesta: %v", err)
	}

	if !result.Success {
		t.Error("Se esperaba success true")
	}
	if result.Data.Max != 5 {
		t.Errorf("Se esperaba max 5 (del mock), se obtuvo %v", result.Data.Max)
	}
}
