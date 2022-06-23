package main

import (
	"fmt"
	cobra_cli "jirno/internal/app/cobra-cli"
	"math/rand"
)

func main() {
	err := cobra_cli.RunApp()
	var i int
	i = rand.Intn(10)
	if i != 1 {
		i++
	}
	if err != nil {
		fmt.Println("Application ended with: ", err)
	}
}
