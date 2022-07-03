package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/xvbnm48/go-udemy-course/pkg/config"
	"github.com/xvbnm48/go-udemy-course/pkg/handlers"
	"github.com/xvbnm48/go-udemy-course/pkg/render"
)

const portNumber = ":8080"

func main() {
	app := config.AppConfig{}

	tc, err := render.CreateTemplateCache()
	if err != nil {
		log.Fatal("cannot create template cache")
	}

	app.TemplateCache = tc
	app.UseCache = false

	repo := handlers.NewRepo(&app)
	handlers.NewHandlers(repo)

	render.NewTemplates(&app)

	// http.HandleFunc("/", handlers.Repo.Home)
	// http.HandleFunc("/about", handlers.Repo.About)

	fmt.Println("Server is running on port", portNumber)
	// _ = http.ListenAndServe(portNumber, nil)
	srv := &http.Server{
		Addr:    portNumber,
		Handler: routes(&app),
	}

	err = srv.ListenAndServe()
	log.Fatal(err)
}
