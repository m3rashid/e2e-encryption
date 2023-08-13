package main

import (
	"net/http"
	"simple-e2ee/handlers"
)

func main() {
	http.HandleFunc("/", handlers.Index)
	http.HandleFunc("/exchange-keys", handlers.ExchangeKeys)

	http.ListenAndServe(":8080", nil)
}
