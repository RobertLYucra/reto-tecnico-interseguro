package handlers

import (
	"bytes"
	"encoding/json"
	"go-matrix-processor/internal/models"
	"go-matrix-processor/internal/service"
	utils "go-matrix-processor/internal/shared/response"
	"net/http"
	"os"

	"github.com/gofiber/fiber/v2"
)

type MatrixHandler struct {
	Service *service.MatrixService
}

func NewMatrixHandler(s *service.MatrixService) *MatrixHandler {
	return &MatrixHandler{Service: s}
}

// ProcessMatrix handles the matrix processing request
func (h *MatrixHandler) ProcessMatrix(c *fiber.Ctx) error {
	var input models.InputMatrix
	if err := c.BodyParser(&input); err != nil {
		return c.Status(400).JSON(utils.ErrorResponse("Entrada inválida", err.Error()))
	}

	qMatrix, rMatrix, err := h.Service.CalculateQR(input.Data)
	if err != nil {
		return c.Status(400).JSON(utils.ErrorResponse("Error al procesar matriz", err.Error()))
	}

	payload := models.QROutput{
		Q: qMatrix,
		R: rMatrix,
	}

	// Send to Node.js API
	nodeURL := os.Getenv("NODE_API_URL")
	if nodeURL == "" {
		nodeURL = "http://localhost:3000/stats"
	}

	jsonData, err := json.Marshal(payload)
	if err != nil {
		return c.Status(500).JSON(utils.ErrorResponse("Error interno al serializar datos", ""))
	}

	nodeResponse, err := http.Post(nodeURL, "application/json", bytes.NewBuffer(jsonData))
	if err != nil {
		return c.Status(502).JSON(utils.ErrorResponse("Error al contactar API de Node.js", err.Error()))
	}
	defer nodeResponse.Body.Close()

	if nodeResponse.StatusCode != http.StatusOK {
		return c.Status(nodeResponse.StatusCode).JSON(utils.ErrorResponse("API de Node.js retornó error", ""))
	}

	// Decode Node.js response (which comes with standard structure)
	// Decode Node.js response (which comes with standard structure)
	var nodeJSResponse struct {
		Success bool                 `json:"success"`
		Message string               `json:"message"`
		Data    models.StatsResponse `json:"data"`
		Errors  []string             `json:"errors"`
	}

	if err := json.NewDecoder(nodeResponse.Body).Decode(&nodeJSResponse); err != nil {
		return c.Status(502).JSON(utils.ErrorResponse("Respuesta inválida de Node.js", ""))
	}

	if !nodeJSResponse.Success {
		var errorDetail string
		if len(nodeJSResponse.Errors) > 0 {
			errorDetail = nodeJSResponse.Errors[0]
		}
		return c.Status(502).JSON(utils.ErrorResponse("Error recibido de Node.js", errorDetail))
	}

	// Merge Stats with Q and R matrices for full response
	fullResponse := models.FullResponse{
		StatsResponse: nodeJSResponse.Data,
		Q:             qMatrix,
		R:             rMatrix,
	}

	return c.JSON(utils.SuccessResponse("Matriz procesada correctamente", fullResponse))
}
