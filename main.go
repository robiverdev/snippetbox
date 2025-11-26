package main

import (
	"log"
	"net/http"
)

// Define a home handler function which writes a byte slice
func home(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Hello from Snippetbox"))
}

func main() {
	mux := http.NewServeMux() // Initialize a new servermux
	mux.HandleFunc("/", home) // Register the home function as handler for the "/" URL pattern

	log.Print("Starting server on: 4000") // Print a log message

	err := http.ListenAndServe(":4000", mux) // Start a new web server (TCP network address to listen on, servermux)
	log.Fatal(err)                           // Log err message and terminate the program
}
