package main

import (
	"encoding/json"
	"fmt"
	log "github.com/sirupsen/logrus"
	"net/http"
	"os"
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

// logging request middleware
func logRequest(handler http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		lrw := NewLoggingResponseWriter(w)
		handler.ServeHTTP(lrw, r)
		log.WithFields(log.Fields{
			"path":   r.URL.Path,
			"method": r.Method,
			"status": lrw.statusCode,
		}).Info("http-request")

	})
}

// LoggingResponseWriter stores the status code http requests
type LoggingResponseWriter struct {
	http.ResponseWriter
	statusCode int
}

// NewLoggingResponseWriter logs a http request to the logfiles
func NewLoggingResponseWriter(w http.ResponseWriter) *LoggingResponseWriter {
	return &LoggingResponseWriter{w, http.StatusOK}
}

// WriteHeader
func (lrw *LoggingResponseWriter) WriteHeader(code int) {
	lrw.statusCode = code
	lrw.ResponseWriter.WriteHeader(code)
}

func initLog() {
	log.SetFormatter(&log.JSONFormatter{})
	file, err := os.OpenFile("phonebook.log", os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
	if err != nil {
		log.Info("Failed to log to file! using default stderr")
	}
	log.SetOutput(file)
	log.SetLevel(log.DebugLevel)
}

// validate validates the input of the PhoneBookEntry
func (p *PhoneBookEntry) validate() error {
	if p.Email == "" {
		err := fmt.Errorf("%s cannot be empty", "email")
		return err
	}
	if p.Lastname == "" {
		err := fmt.Errorf("%s cannot be empty", "lastname")
		return err
	}
	if p.Firstname == "" {
		err := fmt.Errorf("%s cannot be empty", "firstname")
		return err
	}
	if p.Phonenumber == "" {
		err := fmt.Errorf("%s cannot be empty", "phonenumber")
		return err
	}
	return nil
}

// initPhoneBook seeds the database
func (a *App) initPhoneBook() {
	phoneBookSeed := []PhoneBookEntry{
		PhoneBookEntry{
			Firstname:   "Neymar",
			Lastname:    "jr",
			Email:       "neymar@gmail.com",
			Phonenumber: "0612345678",
		},
		PhoneBookEntry{
			Firstname:   "Lionel",
			Lastname:    "Messi",
			Email:       "messi@mail.com",
			Phonenumber: "00448523697",
		},
		PhoneBookEntry{
			Firstname:   "Cristiano",
			Lastname:    "Ronaldo",
			Email:       "ronaldo@fifa.com",
			Phonenumber: "+312080407598",
		},
		PhoneBookEntry{
			Firstname:   "Philipe",
			Lastname:    "Coutinho",
			Email:       "coutinho@fifa.com",
			Phonenumber: "+312023107598",
		},
		PhoneBookEntry{
			Firstname:   "Diego",
			Lastname:    "Costa",
			Email:       "costa@fifa.com",
			Phonenumber: "+312080407598",
		},
		PhoneBookEntry{
			Firstname:   "Harry",
			Lastname:    "Kane",
			Email:       "kane@fifa.com",
			Phonenumber: "+44580407598",
		},
	}
	for _, s := range phoneBookSeed {
		a.DB[s.Email] = s
	}
}
