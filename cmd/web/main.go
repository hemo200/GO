package main

import (
	"flag"
	"log/slog"
	"net/http"
	"os"
)

// this is the main branch
func main() {
	// this variable is to define the default port of the application
	addr := flag.String("addr", ":4000", "HTTP network address")
	flag.Parse() //To read values from runtime

	//Creation of a logger
	logger := slog.New(slog.NewTextHandler(os.Stdout, nil))

	//A new HTTP handler
	mux := http.NewServeMux()
	//file server handler
	fileserver := http.FileServer(http.Dir("./ui/static/"))

	//To handle HTTP requests
	mux.HandleFunc("GET /{$}", home)
	mux.HandleFunc("GET /snippet/view/{id}", snippetView)
	mux.HandleFunc("GET /snippet/create", snippetCreate)
	mux.HandleFunc("POST /snippet/create", snippetCreatePost)
	mux.Handle("GET /static/", http.StripPrefix("/static", fileserver))

	// To display all application info logs
	logger.Info("Starting server", "addr", *addr)

	err := http.ListenAndServe(*addr, mux)

	//to create error logs when something goes wrong
	logger.Error(err.Error())
	os.Exit(1)
}
