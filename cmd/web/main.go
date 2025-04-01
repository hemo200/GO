package main

import (
	"log"
	"net/http"
)

// this is the main branch
func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("GET /{$}", home)
	mux.HandleFunc("GET /snippet/view/{id}", snippetView)
	mux.HandleFunc("GET /snippet/create", snippetCreate)
	mux.HandleFunc("POST /snippet/create", snippetCreatePost)

	//the port to be changes to 6000
	log.Print("starting server on :6000")
	err := http.ListenAndServe(":6000", mux)
	log.Fatal(err)
}
