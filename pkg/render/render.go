package render

import (
	"bytes"
	"fmt"
	"html/template"
	"log"
	"net/http"
	"path/filepath"

	"github.com/xvbnm48/go-udemy-course/pkg/config"
)

var functions = template.FuncMap{}
var app *config.AppConfig

// NewTemplates sets the config for the template package
func NewTemplates(a *config.AppConfig) {
	app = a
}

// renderTemplate is a helper function for rendering HTML templates
func RenderTemplate(w http.ResponseWriter, tmpl string) {
	// get the template cache from the app config
	tc := app.TemplateCache

	// tc, err := CreateTemplateCache()
	// if err != nil {
	// 	fmt.Println("Error rendering template:", err)
	// 	log.Fatal(err)
	// 	return
	// }

	t, ok := tc[tmpl]
	if !ok {
		log.Fatal("cannot get template from template cache")
	}

	buff := new(bytes.Buffer)
	_ = t.Execute(buff, nil)
	_, err := buff.WriteTo(w)
	if err != nil {
		fmt.Println("error waiting project")
		log.Fatal(err)
	}

}

func CreateTemplateCache() (map[string]*template.Template, error) {
	myCache := map[string]*template.Template{}

	pages, err := filepath.Glob("./templates/*.page.tmpl")
	if err != nil {
		return myCache, err
	}

	for _, page := range pages {
		name := filepath.Base(page)
		// fmt.Println("page is currently:", page)
		ts, err := template.New(name).Funcs(functions).ParseFiles(page)
		if err != nil {
			return myCache, err
		}
		matches, err := filepath.Glob("./templates/*.layout.tmpl")
		if err != nil {
			return myCache, err
		}
		if len(matches) > 0 {
			ts, err = ts.ParseGlob("./templates/*.layout.tmpl")
			if err != nil {
				return myCache, err
			}
		}
		myCache[name] = ts
	}

	return myCache, nil

}
