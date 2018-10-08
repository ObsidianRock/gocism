package main

import (
	"fmt"

	"github.com/obsidianrock/gocism/flatten-array"
)

func main() {

	v := flatten.Flatten([]interface{}{1, []interface{}{2, 3, 4, 5, 6, 7}, 8})

	fmt.Println(v)
}
