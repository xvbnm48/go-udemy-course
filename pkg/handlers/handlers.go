package handlers

import (
	"net/http"

	"github.com/xvbnm48/go-udemy-course/pkg/render"
)

// Home page handler
func Home(w http.ResponseWriter, r *http.Request) {

	// fmt.Fprintf(w, "Hello, World!")
	render.RenderTemplate(w, "home.page.tmpl")
}

// about page handler
func About(w http.ResponseWriter, r *http.Request) {
	// fmt.Fprintf(w, "About")
	render.RenderTemplate(w, "about.page.tmpl")
}
