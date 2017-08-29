package main

import (
	_ "fmt"
	_ "github.com/jtb87/config"
)

func main() {
	app := App{}
	app.Init("test_db", "test", "inject")
	defer app.DB.Close()
	// conf, err := config.LoadConfiguration("config.json")
	// if err != nil {
	// 	fmt.Println(err)
	// }
	app.Run()
}
