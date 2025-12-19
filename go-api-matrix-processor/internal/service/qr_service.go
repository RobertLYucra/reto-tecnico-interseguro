package service

import (
	"errors"
	"go-matrix-processor/internal/models"

	"gonum.org/v1/gonum/mat"
)

type MatrixService struct{}

func NewMatrixService() *MatrixService {
	return &MatrixService{}
}

// CalculateQR processes the input matrix and returns Q and R matrices
func (s *MatrixService) CalculateQR(inputMatrix models.Matrix) (models.Matrix, models.Matrix, error) {
	if len(inputMatrix) == 0 {
		return nil, nil, errors.New("input matrix cannot be empty")
	}

	rows := len(inputMatrix)
	cols := len(inputMatrix[0])

	// Flatten 2D slice to 1D slice for gonum
	data := make([]float64, 0, rows*cols)
	for _, row := range inputMatrix {
		if len(row) != cols {
			return nil, nil, errors.New("matrix must be rectangular")
		}
		data = append(data, row...)
	}

	// 1. Rotate Matrix 90 degrees counter-clockwise
	// Requirement: Recibe original -> Rota -> Calcula QR
	rotatedMatrix := s.rotateMatrixCounterClockwise(inputMatrix)

	rows = len(rotatedMatrix)
	cols = len(rotatedMatrix[0])
	data = make([]float64, 0, rows*cols)

	// Flatten rotated matrix
	for _, row := range rotatedMatrix {
		data = append(data, row...)
	}

	dense := mat.NewDense(rows, cols, data)

	// Perform QR Decomposition
	var qr mat.QR
	qr.Factorize(dense)

	var q, r mat.Dense
	qr.QTo(&q)
	qr.RTo(&r)

	return denseToSlice(&q), denseToSlice(&r), nil
}

// rotateMatrixCounterClockwise rotates the matrix 90 degrees counter-clockwise
func (s *MatrixService) rotateMatrixCounterClockwise(matrix models.Matrix) models.Matrix {
	rows := len(matrix)
	cols := len(matrix[0])

	// Function: (i, j) -> (cols - 1 - j, i)
	newRows, newCols := cols, rows
	result := make(models.Matrix, newRows)

	for i := 0; i < newRows; i++ {
		result[i] = make([]float64, newCols)
		for j := 0; j < newCols; j++ {
			result[i][j] = matrix[j][newRows-1-i]
		}
	}
	return result
}

func denseToSlice(dense *mat.Dense) models.Matrix {
	rows, cols := dense.Dims()
	result := make(models.Matrix, rows)
	for i := 0; i < rows; i++ {
		result[i] = make([]float64, cols)
		for j := 0; j < cols; j++ {
			result[i][j] = dense.At(i, j)
		}
	}
	return result
}
