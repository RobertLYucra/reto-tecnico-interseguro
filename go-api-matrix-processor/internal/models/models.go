package models

type Matrix [][]float64

type InputMatrix struct {
	Data Matrix `json:"data"`
}

type QROutput struct {
	Q Matrix `json:"q"`
	R Matrix `json:"r"`
}

type StatsResponse struct {
	Max         float64 `json:"max"`
	Min         float64 `json:"min"`
	Average     float64 `json:"average"`
	TotalSum    float64 `json:"total_sum"`
	IsQDiagonal bool    `json:"is_q_diagonal"`
	IsRDiagonal bool    `json:"is_r_diagonal"`
}

type FullResponse struct {
	StatsResponse
	Q Matrix `json:"q"`
	R Matrix `json:"r"`
}
