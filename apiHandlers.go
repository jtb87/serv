package main

import (
	"encoding/json"
	"log"
	"net/http"
)

func UserHandler(w http.ResponseWriter, r *http.Request) {
	log.Print("'/user' route")
	respondWithJSON(w, 200, "Now for some SQL injection!")
}

func Index(w http.ResponseWriter, r *http.Request) {
	log.Print("'/index' route")
	respondWithJSON(w, "Index does works, are you pondering what I'm pondering?")
}

func respondWithError(w http.ResponseWriter, code int, message string) {
	respondWithJSON(w, code, map[string]string{"error": message})
}

func respondWithJSON(w http.ResponseWriter, code int, payload interface{}) {
	response, _ := json.Marshal(payload)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)
	w.Write(response)
}
