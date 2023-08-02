package main

import (
	"fmt"
	"io"
	"net/http"
)

func main() {
	fmt.Println("Hello World")

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		body, err := io.ReadAll(r.Body)
		if err != nil {
			fmt.Println(err)
		}
		fmt.Println(string(body))
		fmt.Fprintf(w, "Hello World")
		w.Write([]byte("Hello World"))
	})

	http.ListenAndServe(":8080", nil)
}
