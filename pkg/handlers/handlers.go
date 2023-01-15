package handlers

import (
	"bookings-udemy/pkg/config"
	"bookings-udemy/pkg/models"
	"bookings-udemy/pkg/render"
	"net/http"
)

// holds data sent from handlers to templates
//the following code was moved to templatedata.go to avoid import errors
/*
type TemplateData struct {
	StringMap map[string]string
	IntMap map[string]int
	FloatMap map[string]float32
	//when you are unsure of what the data type is you use interface
	//the curly brackets are used because it's declared as a type
	Data map[string]interface{}
	//cross site request forgery token - security token
	CSRFToken string
	//flash message to end user, e.g. "success"
	Flash string
	Warning string
	Error string
}
*/

// Repo the repository used by the handlers
var Repo *Repository

// Repository is the repository type
type Repository struct {
	App *config.AppConfig
}

// NewRepo creates a new repository
func NewRepo(a *config.AppConfig) *Repository {
	return &Repository{
		App: a,
	}
}

// NewHandlers sets the repository for the handlers
func NewHandlers(r *Repository) {
	Repo = r
}

// Home is the handler for the home page
func (m *Repository) Home(w http.ResponseWriter, r *http.Request) {
	render.RenderTemplate(w, "home.page.tmpl", 
	//the empty curly brackets denotes an empty template as a placeholder
	&models.TemplateData{})
}

// About is the handler for the about page
func (m *Repository) About(w http.ResponseWriter, r *http.Request) {
	// perform some logic
	stringMap := make(map[string]string)
	stringMap["test"] = "Hello, again"

	// send data to the template
	render.RenderTemplate(w, "about.page.tmpl", &models.TemplateData{
		StringMap: stringMap,
	})
}
