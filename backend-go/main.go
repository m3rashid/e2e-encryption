package main

import (
	"fmt"
	"net/http"
	"simple-e2ee/handlers"
)

func main() {
	http.HandleFunc("/", handlers.IndexRoute)
	http.HandleFunc("/exchange-keys", handlers.ExchangeKeys)

	fmt.Println("Server is running on port 8080")
	http.ListenAndServe(":8080", nil)
}
