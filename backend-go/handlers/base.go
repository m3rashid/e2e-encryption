package handlers

import (
	"fmt"
	"io"
	"net/http"
)

func IndexRoute(w http.ResponseWriter, r *http.Request) {
	body, err := io.ReadAll(r.Body)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(string(body))
	w.Write([]byte("Hello World"))
}
