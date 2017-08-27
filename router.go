package main

import (
	"github.com/gorilla/mux"
	"net/http"
)

func (a *App) NewRouter() {

	a.Router = mux.NewRouter().StrictSlash(true)
	api := a.Router.PathPrefix("/api").Subrouter()
	a.Router.PathPrefix("/").Handler(http.FileServer(http.Dir("./dist/")))
	for _, route := range routes {
		api.
			Methods(route.Method).
			Path(route.Pattern).
			Name(route.Name).
			Handler(route.HandlerFunc)
	}
}
