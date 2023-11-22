package main

import (
	"flag"
	"fmt"
)

func main() {
	var (
		intFlag    int
		boolFlag   bool
		stringFlag string
	)

	flag.IntVar(&intFlag, "intflag", 0, "int flag value")
	flag.BoolVar(&boolFlag, "boolflag", false, "bool flag value")
	flag.StringVar(&stringFlag, "stringflag", "default", "string flag")

	flag.Parse()

	fmt.Println(intFlag, boolFlag, stringFlag)
}

// go run ./... --intflag 12 --boolflag
