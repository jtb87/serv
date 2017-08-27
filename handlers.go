package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

type HelloWorldJSON struct {
	Message string
}

func HelloWorld(w http.ResponseWriter, r *http.Request) {
	log.Print("logged helloworld")
	respondWithJSON(w, 200, "Hello, world!")
}

func Index(w http.ResponseWriter, r *http.Request) {
	log.Print("logged yes2")
	fmt.Fprintf(w, "Hello, Home, seems to be working no?!")
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
