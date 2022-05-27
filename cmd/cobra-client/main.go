package main

import (
	"fmt"
	"jirno/internal/app/cobra-cli"
)

func main() {
	err := cobra_cli.RunApp()
	if err != nil {
		fmt.Println("Application ended with: ", err)
	}
}

