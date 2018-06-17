package main

import (
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	_ "log"
	"net/http"
)

type PhoneBookEntry struct {
	Id          int    `json:"id"`
	Firstname   string `json:"firstname"`
	Lastname    string `json:"lastname"`
	Email       string `json:"email"`
	Phonenumber int    `json:"phonenumber"` //check if this needs to be string or int?
}

func (e *PhoneBookEntry) validate() {

}

// allEntries will return a list of all the entries in the phonebook
func (a *App) allEntries(w http.ResponseWriter, r *http.Request) {
	pB := []PhoneBookEntry{}
	for _, p := range PhoneBook {
		pB = append(pB, p)
	}
	respondWithJSON(w, http.StatusOK, pB)
}

func (a *App) viewEntry(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	if e, ok := PhoneBook[vars["email"]]; ok {
		respondWithJSON(w, http.StatusOK, e)
	} else {
		respondWithError(w, fmt.Sprintf("Resource '%s' Not Found", vars["email"]))
	}
}

// addEntry will add a new entry to the phonebook if the key does not exist yet, returns the referenceId
func (a *App) addEntry(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()
	p := PhoneBookEntry{}
	decoder := json.NewDecoder(r.Body)
	if err := decoder.Decode(&p); err != nil {
		respondWithError(w, "Invalid request payload")
		return
	}

	if _, ok := PhoneBook[p.Email]; ok {
		respondWithError(w, "Resource with this email address already exists")
		return
	}
	PhoneBook[p.Email] = p
	respondWithJSON(w, http.StatusCreated, map[string]string{"referenceID": p.Email})
}

//
func (a *App) editEntry(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()
	p := PhoneBookEntry{}
	decoder := json.NewDecoder(r.Body)
	if err := decoder.Decode(&p); err != nil {
		respondWithError(w, "Invalid request payload")
		return
	}
	if _, ok := PhoneBook[p.Email]; ok {
		PhoneBook[p.Email] = p
		respondWithJSON(w, http.StatusOK, map[string]string{"referenceID": p.Email})
	} else {
		respondWithError(w, fmt.Sprintf("Resource '%s' Not Found", p.Email))
	}
}

func (a *App) deleteEntry(w http.ResponseWriter, r *http.Request) {

	// PhoneBook[p.Email] = p
	// respondWithJSON(w, 201, fmt.Sprintf(`{"referenceID:%s}`, p.Email))
}

// seeding the phonebook with some records to test
// initPhoneBook()
var PhoneBook = make(map[string]PhoneBookEntry)

func initPhoneBook() {
	phoneBookSeed := []PhoneBookEntry{
		PhoneBookEntry{
			Firstname:   "bear",
			Lastname:    "grylls",
			Email:       "test@gmail.com",
			Phonenumber: 243432,
		},
		PhoneBookEntry{
			Firstname:   "Neymar",
			Lastname:    "jr",
			Email:       "test@mail.com",
			Phonenumber: 243432,
		},
		PhoneBookEntry{
			Firstname:   "Cristiano",
			Lastname:    "Ronaldo",
			Email:       "cr7@fifa.com",
			Phonenumber: 243432,
		},
	}
	for _, s := range phoneBookSeed {
		PhoneBook[s.Email] = s
	}
}
