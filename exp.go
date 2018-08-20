package main

import (
	_ "encoding/hex"
	"encoding/json"
	"fmt"
	"net/http"
	"unicode/utf16"
)

const maxLengthUCS2 = 70
const maxLengthMessage = 67

func (a *App) InitExperimental2() {
	exp := a.Router.PathPrefix("/exp2").Subrouter()
	// exp.Use(authMiddleWareJWT)
	exp.HandleFunc("/encoding", a.encoding).Methods("GET")
	exp.HandleFunc("/floattest", a.floatTest).Methods("POST")

}

type M struct {
	body string
}

func (a *App) interfaceArray(w http.ResponseWriter, r *http.Request) {
	respondWithJSON(w, 200, "json received and printed")
}

func (a *App) encoding(w http.ResponseWriter, r *http.Request) {
	// s := "The message to be sent"
	s := "??????????????????????????????????????????????????????????????????1?????"
	// b := []byte(s)
	var mL []M
	runes := utf16.Encode([]rune(s))
	if len(runes) > maxLengthUCS2 {
		nstop := len(runes) / maxLengthMessage
		if len(runes)%maxLengthMessage > 0 {
			nstop++
		}
		for n := 0; n < nstop; n++ {
			var stop int
			if len(runes) > maxLengthMessage {
				stop = maxLengthMessage
			} else {
				stop = len(runes)
			}
			m := M{
				body: UCS2ToHex(runes[:stop]),
			}
			mL = append(mL, m)
			runes = runes[stop:]
		}
	} else {
		m := M{
			body: UCS2ToHex(runes),
		}
		mL = append(mL, m)
	}
	fmt.Println(mL)
	for _, v := range mL {
		fmt.Println(v.body)
	}
	respondWithJSON(w, 200, "json received and printed")
}

func UCS2ToHex(s []uint16) string {
	h := ""
	for _, a := range s {
		h += fmt.Sprintf("%04x", a)
	}
	return h
}

type Num struct {
	X int     `json:"x,string"`
	Y float64 `json:"y,string"`
}

// addEntry will add a new entry to the a.DB if the key does not exist yet, returns the referenceId
func (a *App) floatTest(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()
	p := Num{}
	decoder := json.NewDecoder(r.Body)
	if err := decoder.Decode(&p); err != nil {
		respondWithError(w, "Invalid request payload")
		return
	}
	ap := p.Y + 1

	respondWithJSON(w, http.StatusCreated, ap)
}
