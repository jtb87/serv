package main

import (
	"fmt"
	_ "strings"
)

type notifier interface {
	notify()
}

type duration struct {
	a int
}

func (d *duration) notify() {
	fmt.Println("Sending Notification in", *d)
}

type run int

func (r *run) notify() {
	fmt.Println(*r)
}

func interfaceArray() {
	m := duration{a: 1}
	m.notify()
	a := run(18)
	a.notify()
	var ar []notifier
	ar = append(ar, &m)
	ar = append(ar, &a)
	for _, s := range ar {
		s.notify()
	}
}

func StringSplitting() {
	s := "12345678901234567890"
	fmt.Println(len(s))
	split := 7
	spl := s[:split]
	fmt.Println(spl)
	var ar []string
	for s != "" {
		var spl string
		switch {
		case len(s) < split:
			spl = s
			s = ""
		default:
			spl = s[:split]
			s = s[split:]
		}
		ar = append(ar, spl)
	}
	fmt.Println(ar)
}
