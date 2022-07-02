package main

import (
	"net/http"
)

// Home page
func Home(w http.ResponseWriter, r *http.Request) {

	// fmt.Fprintf(w, "Hello, World!")
	RenderTemplate(w, "home.page.tmpl")
}

// about page
func About(w http.ResponseWriter, r *http.Request) {
	// fmt.Fprintf(w, "About")
	RenderTemplate(w, "about.page.tmpl")
}
