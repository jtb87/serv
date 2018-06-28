package main

import (
	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
	log "github.com/sirupsen/logrus"
	"net/http"
	"time"
)

// check https://github.com/brainbreaker/rest-and-go/blob/master/store/controller.go

func main() {

	app := App{
		DB: make(map[string]PhoneBookEntry),
	}
	// initialize router
	app.NewRouter()
	// initialize log
	initLog()
	// seed database
	app.initPhoneBook()
	addr := ":9090"
	log.Info("server started")
	log.Printf("server started on: http://localhost%s", addr)
	app.Run(addr)
}

type App struct {
	DB     map[string]PhoneBookEntry
	Router *mux.Router
}

type PhoneBookEntry struct {
	Firstname   string `json:"firstname"`
	Lastname    string `json:"lastname"`
	Email       string `json:"email"`
	Phonenumber string `json:"phonenumber"`
}

// Run starts the server
func (a *App) Run(addr string) {
	allowedHeaders := handlers.AllowedHeaders([]string{"content-type"})
	allowedOrigins := handlers.AllowedOrigins([]string{"*"})
	allowedMethods := handlers.AllowedMethods([]string{"GET", "HEAD", "POST", "PUT", "DELETE", "OPTIONS"})
	s := &http.Server{
		Addr:    addr,
		Handler: http.TimeoutHandler(handlers.CORS(allowedHeaders, allowedOrigins, allowedMethods)(a.Router), time.Second*10, "timeout"),
		// Handler: handlers.CORS(allowedHeaders, allowedOrigins, allowedMethods)(a.Router),
		// ReadTimeout:  10 * time.Second,
		// WriteTimeout: 10 * time.Second,
		IdleTimeout: 120 * time.Second,
	}
	log.Fatal(s.ListenAndServe())
}
