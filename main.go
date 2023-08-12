package main

import (
	"fmt"
	"io"
	"net/http"
	"simple-e2ee/handlers"
)

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		body, err := io.ReadAll(r.Body)
		if err != nil {
			fmt.Println(err)
		}
		fmt.Println(string(body))
		w.Write([]byte("Hello World"))
	})

	http.HandleFunc("/exchange-keys", handlers.ExchangeKeys)

	http.ListenAndServe(":8080", nil)
}
