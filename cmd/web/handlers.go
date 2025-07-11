package main

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
	"strconv"
)

// This is the home page with application logger
func (app *application) home(w http.ResponseWriter, r *http.Request) {
	w.Header().Add("Server", "Go")

	files := []string{
		"./ui/html/base.tmpl",
		"./ui/html/pages/home.tmpl",
		"./ui/html/partials/nav.tmpl",
	}
	// This section will add the comment of the section that will check if the template is accessable or not
	ts, err := template.ParseFiles(files...)

	if err != nil {
		app.logger.Error(err.Error(), "method", r.Method, "uri", r.URL.RequestURI())
		http.Error(w, "Internal  Server Error", http.StatusInternalServerError)
		return
	}
	err = ts.ExecuteTemplate(w, "base", nil)
	if err != nil {
		log.Print(err.Error())
		//logger added to deploy stuctured logger
		app.logger.Error(err.Error(), "method", r.Method, "uri", r.URL.RequestURI())
		http.Error(w, "Internal Server  Error", http.StatusInternalServerError)
	}

}

func (app *application) snippetView(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(r.PathValue("id"))
	if err != nil || id < 1 {
		http.NotFound(w, r)
		return
	}

	fmt.Fprintf(w, "Display a specific snippet with ID %d...", id)
}

// This func to create a new snippet
func (app *application) snippetCreate(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Display a form for creating a new snippet..."))
}

func (app *application) snippetCreatePost(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusCreated)
	w.Write([]byte("Save a new snippet..."))
}
