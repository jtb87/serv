package main

import (
	"fmt"
	"github.com/gorilla/mux"
	"log"
	"net/http"
)

func main() {
	fmt.Println()
	app := App{}
	// initialize router
	app.NewRouter()
	fmt.Println("server initialized")
	// if err != nil {
	// 	fmt.Println(err)
	// }
	app.Run()
}

type App struct {
	// DB     *sqlx.DB
	Router *mux.Router
}

func (a *App) Run() {
	log.Fatal(http.ListenAndServe(":9090", a.Router))
}

// add server configs
type serverConfig struct {
	timeouts     string
	otherconfigs string
}
