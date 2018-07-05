package main

import (
	"encoding/hex"
	"fmt"
	"net/http"
)

func (a *App) InitExperimental2() {
	exp := a.Router.PathPrefix("/exp2").Subrouter()
	// exp.Use(authMiddleWareJWT)
	exp.HandleFunc("/encoding", a.encoding).Methods("GET")
}

func (a *App) encoding(w http.ResponseWriter, r *http.Request) {
	s := "jessee"
	b := []byte(s)

	src := []byte(s)
	dst := make([]byte, hex.EncodedLen(len(src)))
	hex.Encode(dst, src)
	fmt.Printf("bytes %x\n", b)
	fmt.Printf("dst print: %s\n", dst)
	fmt.Printf("values: %v\n", dst)
	respondWithJSON(w, 200, "json received and printed")
}

// %s	the uninterpreted bytes of the string or slice
// %q	a double-quoted string safely escaped with Go syntax
// %x	base 16, lower-case, two characters per byte
// %X	base 16, upper-case, two characters per byte
