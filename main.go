package main

import (
	"fmt"
	"net/http"
)

// Run http server that prints authorization header
func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Println(r.Header.Get("Authorization"))
		fmt.Fprintf(w, "Authorization: %s", r.Header.Get("Authorization"))
	})

	http.ListenAndServe(":8080", nil)
}
