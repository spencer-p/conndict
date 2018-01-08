package main

import (
	"flag"
	"fmt"
)

func main() {
	startword := flag.String("startword", "the", "the word to build the tree from")
	flag.Parse()

	defs, err := Definitions(*startword)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	fmt.Println(defs)
}
