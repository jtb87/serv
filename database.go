package main

import (
	"fmt"
	"net/http"
)

var db2 Store

func (a *App) initDatabase() {
	db := a.Router.PathPrefix("/db").Subrouter()
	// exp.Use(authMiddleWareJWT)
	// exp.HandleFunc("/encoding", a).Methods("GET")
	db.HandleFunc("/todo", a.TodoGetter).Methods("GET")
}

func InitStore(s Store) {
	db2 = s

}

// type Database struct {
// 	DB Store
// }

type Store interface {
	TodoGetter() []Todo
	// TodoPoster()
}

type dbStore struct {
	db map[int]Todo
}

type Todo struct {
	Id          int    `json:"id"`
	Completed   bool   `json:"completed"`
	Subject     string `json:"subject"`
	DateCreated string `json:"date_created"`
	Tag         string `json:"tag"`
}

func (d *dbStore) TodoGetter() []Todo {
	fmt.Println("so far .. ")
	var b []Todo
	for _, v := range d.db {
		fmt.Println(v)
		b = append(b, v)
	}
	return b
}

func (a *App) TodoGetter(w http.ResponseWriter, r *http.Request) {
	// a.DB
	fmt.Println("so far ...  so good")

	ap := db2.TodoGetter()
	respondWithJSON(w, http.StatusCreated, ap)
}
