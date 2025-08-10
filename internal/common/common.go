package common

import (
	"crypto/internal/entity"
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
)

type coinbasePriceResponse struct {
	Data struct {
		Amount string `json:"amount"`
	} `json:"data"`
}

func GetPriceOnline(symbol string) (entity.Price, error) {
	url := fmt.Sprintf("https://api.coinbase.com/v2/prices/%s-%s/spot", symbol, "USD")
	resp, err := http.Get(url)
	if err != nil {
		return entity.Price{}, err
	}
	defer resp.Body.Close()

	var priceResp coinbasePriceResponse
	err = json.NewDecoder(resp.Body).Decode(&priceResp)
	if err != nil {
		return entity.Price{}, err
	}

	price, err := strconv.ParseFloat(priceResp.Data.Amount, 64)
	if err != nil {
		return entity.Price{}, err
	}

	return entity.Price{Asset: symbol, Amount: price}, nil
}
