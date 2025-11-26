package main

import (
	"fmt"
	"log"
	"net/http"
	"strconv"
)

// Define a home handler function which writes a byte slice
func home(w http.ResponseWriter, r *http.Request) {
	_, err := w.Write([]byte("Hello from Snippetbox"))

	if err != nil {
		fmt.Println("Error: ")
	}
}

func snippetView(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(r.PathValue("id")) // Extract the value of the id wildcard from the request (r.PathValue) and convert it to int (Atoi)
	if err != nil || id < 1 {
		http.NotFound(w, r)
		return
	}
	msg := fmt.Sprintf("Display a specific snippet with ID %d", id)
	_, err = w.Write([]byte(msg))
	if err != nil {
		fmt.Println("Error with displaying a specific snippet: ", err)
	}
}

func snippetCreate(w http.ResponseWriter, r *http.Request) {
	_, err := w.Write([]byte("Display a form for creating a snippet"))

	if err != nil {
		fmt.Println("Error creating a snippet:", err)
	}
}

func main() {
	mux := http.NewServeMux()    // Initialize a new servermux
	mux.HandleFunc("/{$}", home) // Register the home function as handler for the "/" URL pattern. {$} - match a single slash, followed by nothing else
	mux.HandleFunc("/snippet/view/{id}", snippetView)
	mux.HandleFunc("/snippet/create", snippetCreate)

	log.Print("Starting server on: 4000") // Print a log message

	err := http.ListenAndServe(":4000", mux) // Start a new web server (TCP network address to listen on, servermux)
	log.Fatal(err)                           // Log err message and terminate the program
}
