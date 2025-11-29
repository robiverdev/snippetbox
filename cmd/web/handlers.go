package main

import (
	"fmt"
	"log"
	"net/http"
	"strconv"
)

func home(w http.ResponseWriter, r *http.Request) {
	w.Header().Add("Server", "Go")
	_, err := w.Write([]byte("Hello from Snippetbox"))
	if err != nil {
		log.Println(err)
	}
}

func snippetView(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(r.PathValue("id"))
	if err != nil || id < 1 {
		http.NotFound(w, r)
		return
	}
	_, err = fmt.Fprintf(w, "Display a specific snippet with ID %d", id)
	if err != nil {
		log.Println(err)
	}
}

func snippetCreate(w http.ResponseWriter, r *http.Request) {
	_, err := w.Write([]byte("Display a form for creating a new snippet"))
	if err != nil {
		log.Println(err)
	}
}

func snippetCreatePost(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusCreated)
	_, err := w.Write([]byte("Save new snippet"))
	if err != nil {
		log.Println(err)
	}
}
