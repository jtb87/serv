package main

import (
	"github.com/gorilla/mux"
	"log"
	"net/http"
)

func (a *App) NewRouter() {
	a.Router = mux.NewRouter().StrictSlash(true)
	api := a.Router.PathPrefix("/api").Subrouter()
	a.Router.PathPrefix("/").Handler(http.FileServer(http.Dir("./dist/")))
	for _, r := range routes {
		api.HandleFunc(r.Pattern, r.HandlerFunc).Methods(r.Method)
	}
}

func (a *App) Run() {
	log.Fatal(http.ListenAndServe(":9090", a.Router))
}
