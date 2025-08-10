package handler

import (
	"encoding/json"
	"net/http"
)

type DefaultResponse struct {
	Type    string `json:"type"`
	Message string `json:"message"`
}

type PriceResponse struct {
	Asset  string  `json:"asset"`
	Amount float64 `json:"amount"`
	Time   int64   `json:"time"`
}

type GetPrice struct {
	Asset string `json:"asset"`
}

func (r DefaultResponse) Response(w http.ResponseWriter, header int) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(header)
	json.NewEncoder(w).Encode(DefaultResponse{Type: r.Type, Message: r.Message})
}

func (r PriceResponse) Response(w http.ResponseWriter, header int) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(header)
	json.NewEncoder(w).Encode(PriceResponse{Asset: r.Asset, Amount: r.Amount})
}
