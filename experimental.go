package main

import (
	"fmt"
	"net/http"
	"time"
)

func (a *App) selectHandler(w http.ResponseWriter, r *http.Request) {
	time.Sleep(time.Second)

	ab := true
	ba := false
	aa := make(chan bool)
	bb := make(chan bool)

	go func() {
		// time.Sleep(time.Millisecond)
		bb <- ba
	}()
	go func() {
		aa <- ab
	}()
	go func() {
		aa <- ab
	}()
	fmt.Println(<-aa)
	select {
	case <-aa:
		respondWithJSON(w, 200, "a ==  true")
	case <-bb:
		fmt.Println(bb)
		respondWithJSON(w, 200, "b == true")
	}
}

func (a *App) contextHandler(w http.ResponseWriter, r *http.Request) {
	respondWithJSON(w, 200, "so far nothing has happened")
}
