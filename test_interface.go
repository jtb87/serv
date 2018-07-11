package main

import (
	"testing"
)

func TestInterface(t *testing.T) {
	s := Str{
		Message: "}aaa{a",
	}
	s.IsValidGSM()
	m := SplitMessage(s)

	fmt.Printf("%+v \n", m)
}
