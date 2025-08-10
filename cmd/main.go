package main

import (
	"log"
	"net/http"

	"crypto/internal/handler"
)

func main() {
	// http.HandleFunc("/currency/add/", handler.Add)
	// http.HandleFunc("/currency/remove/", handler.Remove)
	http.HandleFunc("/currency/price/", handler.GetPriceHandler) //curl -X GET http://localhost:8080/currency/price/ "Content-Type: application/json" -d '{"asset":"LTC"}' | jq
	log.Println("Server start at 8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
