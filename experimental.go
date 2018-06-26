package main

import (
	"fmt"
	"net/http"
	"time"
)

func (a *App) InitExperimental() {
	exp := a.Router.PathPrefix("/exp").Subrouter()
	exp.HandleFunc("/select", a.selectHandler).Methods("GET")
	exp.HandleFunc("/context", a.contextHandler).Methods("GET")
	exp.HandleFunc("/interface", a.interfaceHandler).Methods("GET")

}

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
	tch := make(chan bool, 1)
	go func() {
		time.Sleep(time.Second * 3)
		tch <- true
	}()
	// if timeout occurs before channel tch returns the request will stop
	select {
	case <-r.Context().Done():
		fmt.Println("request has stopped")
		return
	case <-tch:
		fmt.Println("context point1")
		respondWithJSON(w, 200, "request completed")
	}
}

func (a *App) interfaceHandler(w http.ResponseWriter, r *http.Request) {
	v := type1{
		"Gregorius",
	}
	InterfaceExample(&v)
	v2 := type2{
		1245,
		"TMIB",
	}
	InterfaceExample(&v2)
	respondWithJSON(w, 200, "all good")
}

type I interface {
	Method1() string
}

func InterfaceExample(i I) {
	fmt.Println(i.Method1())
}

type type1 struct {
	Name string
}

func (t *type1) Method1() string {
	return t.Name

}

type type2 struct {
	Age  int
	Name string
}

func (t *type2) Method1() string {
	tmp := fmt.Sprintf("my name is: %s and I'm %d old", t.Name, t.Age)
	return tmp
}
