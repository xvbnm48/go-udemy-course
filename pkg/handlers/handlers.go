package handlers

import (
	"net/http"

	"github.com/xvbnm48/go-udemy-course/pkg/config"
	"github.com/xvbnm48/go-udemy-course/pkg/models"
	"github.com/xvbnm48/go-udemy-course/pkg/render"
)

// Repo the repository used by the handlers
var Repo *Repository

// Repository is repository type
type Repository struct {
	App *config.AppConfig
}

// NewRepo creates a new repository
func NewRepo(a *config.AppConfig) *Repository {
	return &Repository{App: a}
}

// NewHandlers sets the repository for the handlers
func NewHandlers(r *Repository) {
	Repo = r
}

// Home page handler
func (m *Repository) Home(w http.ResponseWriter, r *http.Request) {

	// fmt.Fprintf(w, "Hello, World!")
	stringMap := make(map[string]string)
	stringMap["test"] = "hello again"
	render.RenderTemplate(w, "home.page.tmpl", &models.TemplateData{
		StringMap: stringMap,
	})
}

// about page handler
func (m *Repository) About(w http.ResponseWriter, r *http.Request) {
	stringMap := make(map[string]string)
	stringMap["test"] = "hello aruno!"
	render.RenderTemplate(w, "about.page.tmpl", &models.TemplateData{
		StringMap: stringMap,
	})
}
