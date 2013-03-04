package main

import (
	"fmt"
	"net/http"
)

func server_config() {
	server := "localhost"
	port   := "9999"

	server_string := server+":"+port
	return server_string
}

func handler(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Request Served")
	fmt.Fprintf(w, "Hello HTTP World!")
}

func main() {
	http.HandleFunc("/", handler)
	http.ListenAndServe(server_config, nil)

	fmt.Println("Server started on port 9999")
}
