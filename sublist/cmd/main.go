package main

import (
	"fmt"

	"github.com/ObsidianRock/gocism/sublist"
)

func main() {

	r := sublist.Sublist([]int{1, 2, 5}, []int{0, 1, 2, 3, 1, 2, 5, 6})

	fmt.Println(r)
}
