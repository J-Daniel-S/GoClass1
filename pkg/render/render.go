package render

import (
	"bookings-udemy/pkg/config"
	"bookings-udemy/pkg/models"
	"bytes"
	"fmt"
	"html/template"
	"log"
	"net/http"
	"path/filepath"
)

var functions = template.FuncMap{}

var app *config.AppConfig

// NewTemplates sets the config for the template package
func NewTemplates(a *config.AppConfig) {
	app = a
}

//this function will add the elements common to every page on your site
func AddDefaultData(templateData *models.TemplateData) *models.TemplateData {

	return templateData
}

// RenderTemplate renders a template
func RenderTemplate(w http.ResponseWriter, tmpl string, templateData *models.TemplateData) {
	var templateCache map[string]*template.Template

	if app.UseCache {
		// get the template cache from the app config
		templateCache = app.TemplateCache
	} else {
		templateCache, _ = CreateTemplateCache()
	}

	t, ok := templateCache[tmpl]
	if !ok {
		log.Fatal("Could not get template from template cache")
	}

	buf := new(bytes.Buffer)

	templateData = AddDefaultData(templateData)

	_ = t.Execute(buf, templateData)

	_, err := buf.WriteTo(w)
	if err != nil {
		fmt.Println("error writing template to browser", err)
	}

}

// CreateTemplateCache creates a template cache as a map
func CreateTemplateCache() (map[string]*template.Template, error) {

	myCache := map[string]*template.Template{}

	pages, err := filepath.Glob("./templates/*.page.tmpl")
	if err != nil {
		return myCache, err
	}

	for _, page := range pages {
		name := filepath.Base(page)
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
