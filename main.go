package main

import (
	"log"
	"net/http"
)

// Define a home handler function which writes a byte slice
func home(w http.ResponseWriter, r *http.Request) {
	_, err := w.Write([]byte("Hello from Snippetbox"))

	if err != nil {
		log.Println("Error writing response:", err)
		return
	}
}

func main() {
	mux := http.NewServeMux() // Initialize new mux server
	mux.HandleFunc("/", home) // response body - / catch all

	log.Print("Starting server on :4000")

	err := http.ListenAndServe(":4000", mux) // passing in TCP network and mux server
	log.Fatal(err)                           // terminates the program and logs the err msg
}
