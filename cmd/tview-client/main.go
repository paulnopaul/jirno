package main

import (
	tview_cli "jirno/internal/app/tview-cli"
	"log"
)

func main() {
	err := tview_cli.RunApp()
	if err != nil {
		log.Fatalln(err)
	}
}
