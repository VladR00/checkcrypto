package handler

import (
	"crypto/internal/common"
	"crypto/internal/storage"
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/jackc/pgx/v5/pgxpool"
)

type HandlerStorage struct {
	Db *pgxpool.Pool
}

func NewHandlerStorage(db *pgxpool.Pool) *HandlerStorage {
	return &HandlerStorage{Db: db}
}

func (s *HandlerStorage) GetPriceHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		DefaultResponse{Type: "Error", Message: "Only GET method allowed"}.Response(w, http.StatusMethodNotAllowed)
		return
	}

	var asset GetAsset

	if err := json.NewDecoder(r.Body).Decode(&asset); err != nil {
		DefaultResponse{Type: "Error", Message: "Decode. Want 'asset':'string'"}.Response(w, http.StatusNotFound)
		return
	}

	if asset.Asset == "" {
		DefaultResponse{Type: "Error", Message: "Decode. Want 'asset':'string'"}.Response(w, http.StatusNotFound)
		return
	}

	request, err := common.GetPriceOnline(asset.Asset)
	if err != nil {
		log.Println(err)
	}

	PriceResponse{Asset: request.Asset, Amount: request.Amount}.Response(w, http.StatusFound)
}

func (s *HandlerStorage) AddAsset(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		DefaultResponse{Type: "Error", Message: "Only POST method allowed"}.Response(w, http.StatusMethodNotAllowed)
		return
	}

	var asset GetAsset

	if err := json.NewDecoder(r.Body).Decode(&asset); err != nil {
		DefaultResponse{Type: "Error", Message: "Decode. Want 'asset':'string'"}.Response(w, http.StatusBadRequest)
		return
	}

	if asset.Asset == "" {
		DefaultResponse{Type: "Error", Message: "Decode. Want 'asset':'string'"}.Response(w, http.StatusBadRequest)
		return
	}

	strg := storage.NewStorage(s.Db)

	if err := strg.AddAsset(asset.Asset); err != nil {
		log.Println(err)
	}

	DefaultResponse{Type: "Message", Message: fmt.Sprintf("%s add to watch.", asset.Asset)}.Response(w, http.StatusFound)
}

func (s *HandlerStorage) RemoveAsset(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodDelete {
		DefaultResponse{Type: "Error", Message: "Only DELETE method allowed"}.Response(w, http.StatusMethodNotAllowed)
		return
	}

	var asset GetAsset

	if err := json.NewDecoder(r.Body).Decode(&asset); err != nil {
		DefaultResponse{Type: "Error", Message: "Decode. Want 'asset':'string'"}.Response(w, http.StatusBadRequest)
		return
	}

	if asset.Asset == "" {
		DefaultResponse{Type: "Error", Message: "Decode. Want 'asset':'string'"}.Response(w, http.StatusBadRequest)
		return
	}

	strg := storage.NewStorage(s.Db)

	if err := strg.RemoveAsset(asset.Asset); err != nil {
		log.Println(err)
	}

	DefaultResponse{Type: "Message", Message: fmt.Sprintf("%s removed from watch.", asset.Asset)}.Response(w, http.StatusFound)
}
