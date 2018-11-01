package main

import (
	"fmt"

	"github.com/ObsidianRock/gocism/kindergarden"
)

func main() {

	g, err := kindergarten.NewGarden(`
	VRCGVVRVCGGCCGVRGCVCGCGV
	VRCCCGCRRGVCGCRVVCVGCGCV`, []string{
		"Alice", "Bob", "Charlie", "David", "Eve", "Fred",
		"Ginny", "Harriet", "Ileana", "Joseph", "Kincaid", "Larry"})

	if err != nil {
		fmt.Println("err: ", err.Error())
	}

	fmt.Println(g.Plants("Bob"))
}
