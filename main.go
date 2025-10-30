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
		log.Println("Error writing response:", err)
		return
	}
}

func snippetView(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(r.PathValue("id"))

	if err != nil || id < 1 {
		http.NotFound(w, r)
		return
	}
	msg := fmt.Sprintf("Display a specific snippet with ID - %d", id)
	_, err := w.Write([]byte(msg))
	if err != nil {
		log.Printf("Error writing response: %v", err)
		return
	}
}

func snippetCreate(w http.ResponseWriter, r *http.Request) {
	_, err := w.Write([]byte("Display a form for creating a new snippet"))

	if err != nil {
		fmt.Println("Error creating response:", err)
		return
	}
}

func main() {
	mux := http.NewServeMux()                         // Initialize new mux server
	mux.HandleFunc("/{$}", home)                      // response body - / catch all - /{$} match a single slash, followed by nothing else
	mux.HandleFunc("/snippet/view/{id}", snippetView) // {id} wildcard
	mux.HandleFunc("/snippet/create", snippetCreate)

	log.Print("Starting server on :4000")

	err := http.ListenAndServe(":4000", mux) // passing in TCP network and mux server
	log.Fatal(err)                           // terminates the program and logs the err msg
}
