package main

import (
	"fmt"
	"io"
	"net/http"
)

// Run http server that prints authorization header
func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Println(r.Header.Get("Authorization"))
		// print payload to log
		b, err := io.ReadAll(r.Body)
		if err != nil {
			fmt.Println(err)
		}
		fmt.Println("Payload:")
		fmt.Println(string(b))

		fmt.Fprintf(w, "Authorization: %s", r.Header.Get("Authorization"))
	})

	http.ListenAndServe(":8080", nil)
}
