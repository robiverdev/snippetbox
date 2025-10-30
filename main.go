package main

import (
	"fmt"
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

func snippetView(w http.ResponseWriter, r *http.Request) {
	_, err := w.Write([]byte("Display specific snippet"))

	if err != nil {
		log.Println("Error viewing response:", err)
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
	mux := http.NewServeMux() // Initialize new mux server
	mux.HandleFunc("/", home) // response body - / catch all
	mux.HandleFunc("/snippet/view", snippetView)
	mux.HandleFunc("/snippet/create", snippetCreate)

	log.Print("Starting server on :4000")

	err := http.ListenAndServe(":4000", mux) // passing in TCP network and mux server
	log.Fatal(err)                           // terminates the program and logs the err msg
}
