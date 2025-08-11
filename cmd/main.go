package main

import (
	"log"
	"net/http"

	"crypto/internal/handler"
	"crypto/internal/storage"
)

func main() {
	db, err := storage.ConnectPostgreSQL()
	if err != nil {
		log.Panic(err)
	}
	handlerStorage := handler.NewHandlerStorage(db)

	http.HandleFunc("/currency/add/", handlerStorage.AddAsset)
	http.HandleFunc("/currency/remove/", handlerStorage.RemoveAsset)
	http.HandleFunc("/currency/price/", handlerStorage.GetPriceHandler) //curl -X GET http://localhost:8080/currency/price/ "Content-Type: application/json" -d '{"asset":"LTC"}' | jq
	log.Println("Server start at 8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
