package main

import (
	_ "encoding/hex"
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
