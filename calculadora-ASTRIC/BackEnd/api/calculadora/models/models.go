package models

// History estructura del historial
type History struct {
	Date string `json: "date"`
	Operation string `json: "operation"`
	Result float64 `json: "result"`
}

var Historys = []History{}