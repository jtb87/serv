package main

import (
	"github.com/gorilla/mux"
	log "github.com/sirupsen/logrus"
	"net/http"
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
	//
	a.Router.HandleFunc("/", a.healthy).Methods("GET")
	a.Router.HandleFunc("/fail", a.unhealthy).Methods("GET")
}

// NewRouter created a new mux.router and initializes the routes
func (a *App) NewRouter() {
	a.Router = mux.NewRouter().StrictSlash(true)
	a.InitializeRoutes()
	a.InitExperimental()
	a.Router.Use(logRequest)

}

// logging request middleware
func logRequest(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		log.WithFields(log.Fields{
			"path":   r.URL.Path,
			"method": r.Method,
			// "status": lrw.statusCode,
		}).Info("http-request")
		// lrw := NewLoggingResponseWriter(w)
		// call next function/ middleware
		next.ServeHTTP(w, r)
	})
}

// LoggingResponseWriter stores the status code http requests
type LoggingResponseWriter struct {
	http.ResponseWriter
	statusCode int
}

// NewLoggingResponseWriter logs a http request to the logfiles
func NewLoggingResponseWriter(w http.ResponseWriter) *LoggingResponseWriter {
	return &LoggingResponseWriter{w, http.StatusOK}
}

// WriteHeader
func (lrw *LoggingResponseWriter) WriteHeader(code int) {
	lrw.statusCode = code
	lrw.ResponseWriter.WriteHeader(code)
}
