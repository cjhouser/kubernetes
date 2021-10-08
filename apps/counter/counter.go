package main

import (
	"fmt"
	"net/http"
)

func main() {
	counter := 0
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		defer r.Body.Close()
		counter++
		fmt.Fprintf(w, "Count: %d", counter)
	})
	http.ListenAndServe(":8080", nil)
}
