package main

import (
	"encoding/json"
	"fmt"
	"html/template"
	"net/http"
)

var templates = template.Must(template.ParseFiles("view.html"))

// strip is http.StripPrefix for funcs instead of http.Handlers
func strip(prefix string, handler func(http.ResponseWriter, *http.Request)) func(http.ResponseWriter, *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		r.URL.Path = r.URL.Path[len(prefix):]
		handler(w, r)
	}
}

// redirect is http.RedirectHandler for funcs instead of http.Handlers
func redirect(url string, status int) func(http.ResponseWriter, *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		http.Redirect(w, r, url, status)
	}
}

// viewHandler is for the main page where a user explores the dictionary
func viewHandler(w http.ResponseWriter, r *http.Request) {
	// URL.Path should contain the start word.
	// Ideally "/view/dictionary" starts with the word "dictionary".
	templates.ExecuteTemplate(w, "view.html", r.URL.Path)
	return
}

// definitionHandler makes the definition getting abstract for the JavaScript
func definitionHandler(w http.ResponseWriter, r *http.Request) {
	jsonErr := func(err error) {
		fmt.Fprintf(w, "{\"error\":\"%s\"}", err)
	}

	// Word should be in URL path again
	word := r.URL.Path

	// Get the def
	defs, err := Definitions(word)
	if err != nil {
		jsonErr(err)
		return
	}

	// Spit out the definition as JSON
	type Response struct {
		Defs []string `json:"definitions"`
	}
	response, err := json.Marshal(Response{defs})
	if err != nil {
		jsonErr(err)
		return
	}

	fmt.Fprintf(w, "%s", response)
}

func serve() {
	// Set handle funcs
	http.HandleFunc("/view/", strip("/view/", viewHandler))
	http.HandleFunc("/", redirect("/view/", http.StatusFound))
	http.HandleFunc("/def/", strip("/def/", definitionHandler))

	// Serve forever
	http.ListenAndServe(":8000", nil)
}
