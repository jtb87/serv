package main

import (
	"github.com/gorilla/mux"
)

// InitializeRoutes defines the routes for the piricing api
func (a *App) InitializeRoutes() {
	api := a.Router.PathPrefix("/api").Subrouter()
	api.HandleFunc("/entry", a.allEntries).Queries("search", "{search}")
	api.HandleFunc("/entry", a.allEntries).Methods("GET")
	api.HandleFunc("/entry", a.addEntry).Methods("POST")
	api.HandleFunc("/entry", a.editEntry).Methods("PUT")
	api.HandleFunc("/entry/{email}", a.viewEntry).Methods("GET")
	api.HandleFunc("/entry/{email}", a.deleteEntry).Methods("DELETE")
}

// NewRouter created a new mux.router and initializes the routes
func (a *App) NewRouter() {
	a.Router = mux.NewRouter().StrictSlash(true)
	a.InitializeRoutes()
}
