package main

import (
	_ "database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/gorilla/mux"
	"github.com/jmoiron/sqlx"
	"log"
)

type User struct {
	Id       int    `db:"id"`
	Name     string `db:"name"`
	Password string `db:"password"`
	UserName string `db:"username"`
}

type App struct {
	DB     *sqlx.DB
	Router *mux.Router
}

//https://semaphoreci.com/community/tutorials/building-and-testing-a-rest-api-in-go-with-gorilla-mux-and-postgresql
func (a *App) Init(user, password, dbname string) {
	connectionString :=
		fmt.Sprintf("user=%s password=%s dbname=%s", user, password, dbname)
	var err error
	a.DB, err = sqlx.Open("mysql", connectionString)
	if err != nil {
		log.Fatal(err)
	}
}
