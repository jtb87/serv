package main

import (
	"fmt"
	"github.com/gorilla/mux"
	"log"
	"net/http"
	"time"
)

func main() {
	fmt.Println()
	app := App{}
	// initialize router
	app.NewRouter()
	initPhoneBook()
	fmt.Println("server initialized on port")
	app.Run()

}

type App struct {
	// DB     *sqlx.DB
	Router *mux.Router
}

// Run starts initializes the application
func (a *App) Run() {
	s := &http.Server{
		Addr:           ":9090",
		Handler:        http.TimeoutHandler(a.Router, time.Second*10, "Timeout!"),
		MaxHeaderBytes: 1 << 20,
	}
	log.Fatal(s.ListenAndServe())
}
