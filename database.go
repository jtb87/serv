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

func (a *App) Init(user, password, dbname string) {
	connectionString := fmt.Sprintf("%s:%s@tcp(127.0.0.1:3306)/%s", user, password, dbname)
	var err error
	a.DB, err = sqlx.Open("mysql", connectionString)
	if err != nil {
		log.Fatal(err)
	}
	err = a.DB.Ping()
	if err != nil {
		log.Print(err)
	}
	a.NewRouter()
}
