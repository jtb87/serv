package main

import (
	_ "encoding/hex"
	"fmt"
	"net/http"
)

func (a *App) InitExperimental2() {
	exp := a.Router.PathPrefix("/exp2").Subrouter()
	// exp.Use(authMiddleWareJWT)
	exp.HandleFunc("/encoding", a.encoding).Methods("GET")
}

func (a *App) encoding(w http.ResponseWriter, r *http.Request) {
	// s := "The message to be sent"
	// b := []byte(s)
	b := []byte("2") //0x30
	// src := []byte(s)
	// dst := make([]byte, hex.EncodedLen(len(src)))
	// hex.Encode(dst, src)
	fmt.Println(b)
	fmt.Printf("bytes %s \n", b)
	fmt.Printf("bytes %x \n", b)

	// bin := "546865206d65737361676520746f2062652073656e74"
	// dst2 := fmt.Sprintf("%s", dst)
	// if dst2 == bin {
	// 	fmt.Println("True they are equal")
	// }
	// fmt.Printf("dst print: %x\n", dst)
	// fmt.Printf("values: %v\n", dst)
	const placeOfInterest = `2`

	fmt.Printf("plain string: ")
	fmt.Printf("%s", placeOfInterest)
	fmt.Printf("\n")

	fmt.Printf("quoted string: ")
	fmt.Printf("%+q", placeOfInterest)
	fmt.Printf("\n")

	fmt.Printf("%v \n", placeOfInterest)
	fmt.Printf("hex bytes: ")
	for i := 0; i < len(placeOfInterest); i++ {
		fmt.Printf("%x ", placeOfInterest[i])
	}
	fmt.Printf("\n")

	respondWithJSON(w, 200, "json received and printed")
}

// %s	the uninterpreted bytes of the string or slice
// %q	a double-quoted string safely escaped with Go syntax
// %x	base 16, lower-case, two characters per byte
// %X	base 16, upper-case, two characters per byte
// 0x1B 0x1B
