package main

import (
	"log"
	"net/http"
)

// Define a home handler function which writes a byte slice
func home(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Hello from Snippetbox"))
}

func snippetView(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Display a specific snippet"))
}

func snippetCreate(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Display a form for creating a snippet"))
}

func main() {
	mux := http.NewServeMux()    // Initialize a new servermux
	mux.HandleFunc("/{$}", home) // Register the home function as handler for the "/" URL pattern. {$} - match a single slash, followed by nothing else
	mux.HandleFunc("/snippet/view", snippetView)
	mux.HandleFunc("/snippet/create", snippetCreate)

	log.Print("Starting server on: 4000") // Print a log message

	err := http.ListenAndServe(":4000", mux) // Start a new web server (TCP network address to listen on, servermux)
	log.Fatal(err)                           // Log err message and terminate the program
}
