package main

import (
	"bookings-udemy/pkg/config"
	"bookings-udemy/pkg/handlers"
	"net/http"

	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"

	//"github.com/bmizerany/pat"
)

func routes(app *config.AppConfig) http.Handler {

	/*
	//create a multiplexer (mux, an http handler)
	//pat is a good one.  install using go get github.com/bmizerany/pat
	//evidently third party http handlers are more efficient than standard
	The following is for pat router
	mux := pat.New();

	mux.Get("/", http.HandlerFunc(handlers.Repo.Home));
	mux.Get("/about", http.HandlerFunc(handlers.Repo.About));
	*/

	//we're using chi here
	mux := chi.NewRouter();

	//middleware is easy to build into chi router.  Observe below
	//recoverer prevents crashes due to unhandled exceptions
	mux.Use(middleware.Recoverer);
	//handles csrf foolishness
	mux.Use(noSerf);
	//handles sessions
	mux.Use(SessionLoad);

	mux.Get("/", handlers.Repo.Home);
	mux.Get("/about", handlers.Repo.About);

	return mux;
}