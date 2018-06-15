package main

import (
	"github.com/gorilla/mux"
	"net/http"
)

// TODO:  add error checks -- check documentation
func (app *App) NewRouter() {
	app.Router = mux.NewRouter().StrictSlash(true)
	api := app.Router.PathPrefix("/api").Subrouter()
	app.Router.PathPrefix("/").Handler(http.FileServer(http.Dir("./dist/")))
	for _, r := range routes {
		api.HandleFunc(r.Pattern, r.HandlerFunc).Methods(r.Method)
	}
}
