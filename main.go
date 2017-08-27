package main

import (
	_ "fmt"
	_ "github.com/jtb87/config"
	"log"
	"net/http"
)

func main() {
	app := App{}
	app.Init("test_db", "test", "inject")
	app.NewRouter() // conf, err := config.LoadConfiguration("config.json")
	// if err != nil {
	// 	fmt.Println(err)
	// }
	// Open_db()
	log.Fatal(http.ListenAndServe(":9090", app.Router))

}

// https://thenewstack.io/make-a-restful-json-api-go/
// https://www.thepolyglotdeveloper.com/2016/12/create-real-time-chat-app-golang-angular-2-websockets/

// func Open_db() {
// 	app := App{}
// 	db, err := sqlx.Open("mysql", "test_db:test@(localhost:3306)")
// 	if err != nil {
// 		log.Panic(err)
// 	}
// 	err = db.Ping()
// 	if err != nil {
// 		log.Panic(err)
// 	}
// 	log.Print(app)
// }
