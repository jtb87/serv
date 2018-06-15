package main

import (
	"log"
	"net/http"
)

func UserHandler(w http.ResponseWriter, r *http.Request) {
	log.Print("'/user' route")
	respondWithJSON(w, 200, "Now for some SQL injection!")
}

func Index(w http.ResponseWriter, r *http.Request) {
	log.Print("'/index' route")
	respondWithJSON(w, 200, "Index works")
}
