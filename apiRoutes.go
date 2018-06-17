package main

import (
	"github.com/gorilla/mux"
	_ "log"
	_ "net/http"
)

// InitializeRoutes defines the routes for the piricing api
func (a *App) InitializeRoutes() {
	api := a.Router.PathPrefix("/api").Subrouter()
	api.HandleFunc("/entry", a.allEntries).Methods("GET")
	api.HandleFunc("/entry", a.addEntry).Methods("POST")
	api.HandleFunc("/entry", a.editEntry).Methods("PUT")
	api.HandleFunc("/entry/{email}", a.viewEntry).Methods("GET")
	// api.HandleFunc("/entry", a.deleteEntry).Methods("DELETE")
}

// NewRouter created a new mux.router and initializes the routes
func (a *App) NewRouter() {
	a.Router = mux.NewRouter().StrictSlash(true)
	a.InitializeRoutes()
}

// write headers
// w.Header().Set("Content-Type", "application/json")
// w.Header().Set("Content-Length", strconv.Itoa(len(data)))
// w.WriteHeader(http.StatusOK)
