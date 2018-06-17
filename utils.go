package main

import (
	"encoding/json"
	"net/http"
)

func respondWithError(w http.ResponseWriter, message string) {
	respondWithJSON(w, http.StatusBadRequest, map[string]string{"error": message})
}

func respondWithJSON(w http.ResponseWriter, code int, payload interface{}) {
	response, _ := json.Marshal(payload)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)
	w.Write(response)
}
