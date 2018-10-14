package main

import (
	"fmt"

	"github.com/obsidianRock/gocism/protein-translation"
)

func main() {

	s, err := protein.FromRNA("AUGUUUUGGUU")

	fmt.Println(s, err)
}
