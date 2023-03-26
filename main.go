package main

import (
	"fmt"
	"net/http"
)

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, req *http.Request) {
		fmt.Fprintf(w, "<h1>Hello from Docker container!</h1>")
	})
	http.ListenAndServe(":8081", nil)
}
