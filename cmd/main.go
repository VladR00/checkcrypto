package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strconv"
)

const (
	coinbaseCryptoAPI = "https://api.coinbase.com/v2/prices/%s-%s/spot"
	currency          = "USD"
)

type coinbasePriceResponse struct {
	Data struct {
		Amount string `json:"amount"`
	} `json:"data"`
}

func main() {
	symbol := "LTC"
	price, err := getCoinbasePrice(symbol)
	if err != nil {
		log.Println(err)
	}

	log.Printf("Asset: %s; Amount: %.2f $", symbol, price)
}

func getCoinbasePrice(symbol string) (float64, error) {
	url := fmt.Sprintf(coinbaseCryptoAPI, symbol, currency)
	resp, err := http.Get(url)
	if err != nil {
		return 0, err
	}
	defer resp.Body.Close()

	var priceResp coinbasePriceResponse
	err = json.NewDecoder(resp.Body).Decode(&priceResp)
	if err != nil {
		return 0, err
	}

	price, err := strconv.ParseFloat(priceResp.Data.Amount, 64)
	if err != nil {
		return 0, err
	}

	return price, nil
}
