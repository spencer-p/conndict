package main

import (
	"flag"
	"fmt"
	"os"
)

func main() {
	startword := flag.String("get", "", "a word to get to test")
	shouldserve := flag.Bool("serve", false, "run the webserver")
	flag.Parse()

	if *startword != "" {
		defs, err := Definitions(*startword)
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}
		fmt.Println(defs)
	}

	if *shouldserve {
		serve()
	}
}
