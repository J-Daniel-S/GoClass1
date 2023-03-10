package main

import (
	"net/http"

	"github.com/justinas/nosurf"
)

//it is common practice to name the arg taken by middleware "next"
//most middleware looks very similar to this
//all middleware must take in and return an http.Handler
/*
func writeToConsole(next http.Handler) http.Handler {
	return http.HandlerFunc(func(writer http.ResponseWriter, request *http.Request) {
		fmt.Println("Hit the page");
		next.ServeHTTP(writer, request);
	})
}
*/

//jwt.  Much easier.  Suck it jawoot
//adds CSRF protection on post requests
func noSerf(next http.Handler) http.Handler {
	csrfHandler := nosurf.New(next);

	csrfHandler.SetBaseCookie(http.Cookie{
		HttpOnly: true,
		//for cookies path is always "/" because they apply to the entire site
		Path: "/",
		Secure: app.InProduction,
		SameSite: http.SameSiteLaxMode,
	})
	return csrfHandler;
}

//Loads and saves the session on every request
func SessionLoad(next http.Handler) http.Handler {
	return session.LoadAndSave(next);
}