package service_test

import (
	"math"
	"testing"

	"go-matrix-processor/internal/models"
	"go-matrix-processor/internal/service"
)

func TestCalculateQR_IdentityMatrix(t *testing.T) {
	// Arrange
	svc := service.NewMatrixService()
	input := models.Matrix{
		{1, 0},
		{0, 1},
	}

	// Act
	q, r, err := svc.CalculateQR(input)

	// Assert
	if err != nil {
		t.Fatalf("Error inesperado: %v", err)
	}

	// Check Dimensions Q (2x2) and R (2x2)
	if len(q) != 2 || len(q[0]) != 2 {
		t.Errorf("Se esperaba Q de 2x2, se obtuvo %dx%d", len(q), len(q[0]))
	}
	if len(r) != 2 || len(r[0]) != 2 {
		t.Errorf("Se esperaba R de 2x2, se obtuvo %dx%d", len(r), len(r[0]))
	}

	// Verify Reconstruction: Q * R = Rotated(Identity) = [[0, 1], [1, 0]]
	// For [[0, 1], [1, 0]]:
	// Q = [[0, 1], [1, 0]]
	// R = [[1, 0], [0, 1]]

	// So Q[0][0] should be 0 (or close to 0)
	if math.Abs(q[0][0]) > 0.01 {
		t.Errorf("Se esperaba Q[0][0] cercano a 0 (por rotación), se obtuvo %v", q[0][0])
	}
	// Q[0][1] should be 1 or -1 (QR decomposition sign ambiguity)
	if math.Abs(math.Abs(q[0][1])-1.0) > 0.01 {
		t.Errorf("Se esperaba |Q[0][1]| cercano a 1, se obtuvo %v", q[0][1])
	}
}

func TestCalculateQR_InvalidInput(t *testing.T) {
	svc := service.NewMatrixService()

	// Case: Empty
	_, _, err := svc.CalculateQR(models.Matrix{})
	if err == nil {
		t.Error("Se esperaba error para matriz vacía, se obtuvo nil")
	}

	// Case: Non-rectangular
	nonRect := models.Matrix{
		{1, 2},
		{3},
	}
	_, _, err = svc.CalculateQR(nonRect)
	if err == nil {
		t.Error("Se esperaba error para matriz no rectangular, se obtuvo nil")
	}
}
