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
var Repo *Repository;

// Repository is the repository type
type Repository struct {
	App *config.AppConfig;
}

// NewRepo creates a new repository
func NewRepo(appConfig *config.AppConfig) *Repository {
	return &Repository{
		App: appConfig,
	}
}

// NewHandlers sets the repository for the handlers
func NewHandlers(repo *Repository) {
	Repo = repo;
}

// Home is the handler for the home page
func (m *Repository) Home(writer http.ResponseWriter, request *http.Request) {

	//pulls the remote IP address out of the request
	remoteIP := request.RemoteAddr;
	//stores in the request context the remoteIP address with the key "remote_ip"
	m.App.Session.Put(request.Context(), "remote_ip", remoteIP);

	render.RenderTemplate(writer, "home.page.tmpl", 
	//the empty curly brackets denotes an empty template as a placeholder
	&models.TemplateData{})
}

// About is the handler for the about page
func (m *Repository) About(writer http.ResponseWriter, request *http.Request) {
	// perform some logic
	stringMap := make(map[string]string);
	stringMap["test"] = "Hello, again";

	//gets the remoteIP for the session from the session cookie
	remoteIP := m.App.Session.GetString(request.Context(), "remote_ip");
	//places the ip in the stringmap
	stringMap["remote_ip"] = remoteIP;

	// send data to the template
	render.RenderTemplate(writer, "about.page.tmpl", &models.TemplateData{
		StringMap: stringMap,
	})
}
