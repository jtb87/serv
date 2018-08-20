package main

import (
	"fmt"
	"github.com/gobuffalo/packr"
	"os"
)

func Packr() {
	p, err := os.Getwd()
	if err != nil {
		fmt.Println(err)
	}
	box := packr.NewBox(p + "/sql")
	a := box.String("adderPricing.sql")
	fmt.Println(a)
}
