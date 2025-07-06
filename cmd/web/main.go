package main

import (
	"flag"
	"log"
	"net/http"
)

// this is the main branch
func main() {
	// this variable is to define the  port of the application
	addr := flag.String("addr", ":5000", "HTTP network adress")
	mux := http.NewServeMux()
	fileserver := http.FileServer(http.Dir("./ui/static/"))
	mux.HandleFunc("GET /{$}", home)
	mux.HandleFunc("GET /snippet/view/{id}", snippetView)
	mux.HandleFunc("GET /snippet/create", snippetCreate)
	mux.HandleFunc("POST /snippet/create", snippetCreatePost)
	mux.Handle("GET /static/", http.StripPrefix("/static", fileserver))

	//the port to be changes to 5000
	log.Printf("starting server on %s", *addr)
	err := http.ListenAndServe(*addr, mux)
	log.Fatal(err)
}
