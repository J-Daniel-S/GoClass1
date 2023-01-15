package main

import (
	"bookings-udemy/pkg/config"
	"bookings-udemy/pkg/handlers"
	"bookings-udemy/pkg/render"
	"fmt"
	"log"
	"net/http"
)

const portNumber = ":8080"

// main is the main function
func main() {
	var app config.AppConfig;

	tc, err := render.CreateTemplateCache();
	if err != nil {
		log.Fatal("cannot create template cache")
	};

	app.TemplateCache = tc;
	app.UseCache = false;

	repo := handlers.NewRepo(&app);
	handlers.NewHandlers(repo);

	render.NewTemplates(&app);

	//default http handler.  commented out bc we're using pat
	/*
	http.HandleFunc("/", handlers.Repo.Home)
	http.HandleFunc("/about", handlers.Repo.About)
	*/

	fmt.Println(fmt.Sprintf("Staring application on port %s", portNumber));
	//the following is used when using the default http handler
	//_ = http.ListenAndServe(portNumber, nil)

	serve := &http.Server {
		//Addr is address you java slave
		Addr: portNumber,
		Handler: routes(&app),
	}

	err = serve.ListenAndServe();
	log.Fatal(err);

}
