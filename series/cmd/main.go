package main

import (
	"fmt"

	"github.com/obsidianrock/gocism/series"
)

func main() {

	s := series.All(2, "12345")
	fmt.Println(s)
}
