package main

import (
	"bytes"
	"exercism/tournament"
	"fmt"
	"io/ioutil"
	"os"
)

func main() {

	r, err := ioutil.ReadFile("tournament.txt")

	if err != nil {
		fmt.Println(err)
		os.Exit(2)
	}

	v := bytes.NewBuffer(r)
	if err != nil {
		fmt.Println(err)
		os.Exit(2)
	}

	err = tournament.Tally(v, os.Stdout)
	if err != nil {
		fmt.Println(err)
		os.Exit(2)
	}
}
