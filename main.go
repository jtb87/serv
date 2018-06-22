package main

import (
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
	log.Info("server started")
	app.Run()
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
func (a *App) Run() {
	s := &http.Server{
		Addr:         ":9090",
		Handler:      http.TimeoutHandler(a.Router, time.Second*10, "timeout"),
		ReadTimeout:  10 * time.Second,
		WriteTimeout: 10 * time.Second,
		IdleTimeout:  120 * time.Second,
	}
	log.Fatal(s.ListenAndServe())
}
