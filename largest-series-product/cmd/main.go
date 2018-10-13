package main

import (
	"fmt"

	"github.com/obsidianRock/gocism/largest-series-product"
)

func main() {

	sum, _ := lsproduct.LargestSeriesProduct("0123456789", 2)
	fmt.Println(sum)
}
