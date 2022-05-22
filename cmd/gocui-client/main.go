package main

import (
	app "jirno/internal/app/gocui-cli"
	"log"
)

func main() {
	err := app.RunApp()
	if err != nil {
		log.Fatalln(err)
	}
}
