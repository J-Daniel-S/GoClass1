package main

import (
	"bookings-udemy/pkg/config"
	"bookings-udemy/pkg/handlers"
	"net/http"

	"github.com/bmizerany/pat"
)

func routes(app *config.AppConfig) http.Handler {
	//create a multiplexer (an http handler)
	//pat is a good one.  install using go get github.com/bmizerany/pat
	mux := pat.New();

	mux.Get("/", http.HandlerFunc(handlers.Repo.Home));
	mux.Get("/about", http.HandlerFunc(handlers.Repo.About));

	return mux;
}