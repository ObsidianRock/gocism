package main

import (
	"fmt"
	"robot-name"
)

func main() {

	r := robotname.New()

	n1 := r.Name()

	n2 := r.Name()

	fmt.Println(n2, n1)

}
