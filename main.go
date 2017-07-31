package main

import (
	"fmt"
	"github.com/gorilla-mux"
	"net/http"
)

func main() {
	http.HandleFunc("/hello", helloWorld)    // set router
	http.HandleFunc("/", helloHome)          // set router
	err := http.ListenAndServe(":9090", nil) // set listen port
	if err != nil {
		fmt.Println("ListenAndServe: ", err)
	}
}

func helloWorld(w http.ResponseWriter, r *http.Request) {
	fmt.Println("logged yes")
	fmt.Fprintf(w, "Hello, world!")
}

func helloHome(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello, Home, seems to be working no?!")
}
