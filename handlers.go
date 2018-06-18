package main

import (
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"net/http"
	"strings"
)

// viewEntry returns single entry based on parameter in url
func (a *App) viewEntry(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	if e, ok := a.DB[vars["email"]]; ok {
		respondWithJSON(w, http.StatusOK, e)
	} else {
		respondWithError(w, fmt.Sprintf("Resource '%s' Not Found", vars["email"]))
	}
}

// addEntry will add a new entry to the a.DB if the key does not exist yet, returns the referenceId
func (a *App) addEntry(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()
	p := PhoneBookEntry{}
	decoder := json.NewDecoder(r.Body)
	if err := decoder.Decode(&p); err != nil {
		respondWithError(w, "Invalid request payload")
		return
	}
	// validate input
	if err := p.validate(); err != nil {
		respondWithError(w, err.Error())
		return
	}
	if _, ok := a.DB[p.Email]; ok {
		respondWithError(w, "Resource with this email address already exists")
		return
	}
	a.DB[p.Email] = p
	respondWithJSON(w, http.StatusCreated, map[string]string{"referenceID": p.Email})
}

// editEntry edit entry in database returns resourceId
func (a *App) editEntry(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()
	p := PhoneBookEntry{}
	decoder := json.NewDecoder(r.Body)
	if err := decoder.Decode(&p); err != nil {
		respondWithError(w, "Invalid request payload")
		return
	}
	if _, ok := a.DB[p.Email]; ok {
		a.DB[p.Email] = p
		respondWithJSON(w, http.StatusOK, map[string]string{"referenceID": p.Email})
	} else {
		respondWithError(w, fmt.Sprintf("Resource '%s' Not Found", p.Email))
	}
}

// deleteEntry deletes a resource from the database
func (a *App) deleteEntry(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	if _, ok := a.DB[vars["email"]]; ok {
		delete(a.DB, vars["email"])
		respondWithJSON(w, http.StatusOK, "")
	} else {
		respondWithError(w, "Resource Not Found")
	}
}

// allEntries will return a list of all the entries or of all filtered entries in the database
func (a *App) allEntries(w http.ResponseWriter, r *http.Request) {
	pB := []PhoneBookEntry{}
	vars := mux.Vars(r)
	// checks if search parameter exists and searches through database
	if s, ok := vars["search"]; ok {
		for k, v := range a.DB {
			// check if key starts with searchterm
			if strings.HasPrefix(k, s) {
				pB = append(pB, v)
			}
		}
	} else {
		for _, p := range a.DB {
			pB = append(pB, p)
		}
	}
	respondWithJSON(w, http.StatusOK, pB)
}
