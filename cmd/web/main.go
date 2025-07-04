package main

import (
	"log"
	"net/http"
)

// this is the main branch
func main() {
	mux := http.NewServeMux()
	fileserver := http.FileServer(http.Dir("./ui/static/"))
	mux.HandleFunc("GET /{$}", home)
	mux.HandleFunc("GET /snippet/view/{id}", snippetView)
	mux.HandleFunc("GET /snippet/create", snippetCreate)
	mux.HandleFunc("POST /snippet/create", snippetCreatePost)
	mux.Handle("GET /static/", http.StripPrefix("/static", fileserver))

	//the port to be changes to 5000
	log.Print("starting server on :5000")
	err := http.ListenAndServe(":5000", mux)
	log.Fatal(err)
}
