package main

import (
	"fmt"
	"net/http"
)

func handler(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Request Served")
	fmt.Fprintf(w, "Hello HTTP World!")
}

func main() {
	fmt.Println("Server started on port 9999")

	http.HandleFunc("/", handler)
	http.ListenAndServe("localhost:9999", nil)	
}
