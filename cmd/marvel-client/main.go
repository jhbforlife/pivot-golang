package main

import (
	"fmt"
	"log"

	"github.com/jhbforlife/pivot-golang/cmd/marvel-client/marvel"
)

func main() {
	client := marvel.NewCharClient()
	chars, err := client.GetCharsWithLimit(4)
	if err != nil {
		log.Fatal(err)
	}
	for _, char := range chars {
		fmt.Println(char.Name)
		fmt.Println(char.Description)
		fmt.Println()
	}
}
