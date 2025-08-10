package handler

import (
	common "crypto/internal/common"
	"encoding/json"
	"log"
	"net/http"
)

func GetPriceHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		DefaultResponse{Type: "Error", Message: "Only GET method allowed"}.Response(w, http.StatusMethodNotAllowed)
		return
	}

	var asset GetPrice

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
