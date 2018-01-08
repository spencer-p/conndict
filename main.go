package main

import (
	"flag"
	"fmt"
)

func main() {
	startword := flag.String("startword", "the", "the word to build the tree from")
	flag.Parse()
	fmt.Println(Definition(*startword))
}
