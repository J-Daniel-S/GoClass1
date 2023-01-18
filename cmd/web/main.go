package main

import (
	"bookings-udemy/pkg/config"
	"bookings-udemy/pkg/handlers"
	"bookings-udemy/pkg/render"
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/alexedwards/scs/v2"
)

//these are placed outside of the main function so that they will be
//accessible anywhere in the package
const portNumber = ":8080"
var app config.AppConfig;
var session *scs.SessionManager;

// main is the main function
func main() {
	//change this to true when in production
	app.InProduction = false;

	//the following uses packages from github.com/alexedwards/scs
	//I assume it has something to do with session management
	session = scs.New();
	//this causes the session to last for 24 hours
	//by default our session information is kept in cookies.
	//Can be stored elsewhere e.g redis
	session.Lifetime = 24 * time.Hour;
	//parameters for cookies set here
	//this causes the session to persist even if the browser is closed
	session.Cookie.Persist = true;
	//this tells the browser how specific we want to be about the site
	session.Cookie.SameSite = http.SameSiteLaxMode;
	//the following requires https to connect when set to true
	session.Cookie.Secure = app.InProduction;

	//this passes the session created above to the config file for use in other packages
	app.Session = session;


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
