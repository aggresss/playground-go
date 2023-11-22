package main

import (
	"flag"
	"fmt"
	"os"
	"strconv"
)

type alpha struct {
	i int64
}

var _ flag.Value = &alpha{}

func (a *alpha) Set(s string) error {
	i, err := strconv.ParseInt(s, 10, 64)
	if err != nil {
		return err
	}
	a.i = i
	return nil
}

func (a *alpha) String() string {
	return strconv.FormatInt(int64(a.i), 10)
}

type argments struct {
	intFlag    int
	boolFlag   bool
	stringFlag string
	alphaFlag  alpha
}

func (a *argments) parseArgs(args []string) error {
	fs := flag.NewFlagSet(args[0], flag.ExitOnError)

	fs.IntVar(&a.intFlag, "intflag", 1, "int flag value")
	fs.BoolVar(&a.boolFlag, "boolflag", false, "bool flag value")
	fs.StringVar(&a.stringFlag, "stringflag", "default", "string flag")
	fs.Var(&a.alphaFlag, "alphaflag", "alpha flag")

	fs.Usage = func() {
		fmt.Fprintf(os.Stderr, "flags: %s", args[0])
		fs.PrintDefaults()
		os.Exit(1)
	}

	if err := fs.Parse(os.Args[1:]); err != nil {
		return err
	}

	return nil
}

func main() {
	args := &argments{}
	if err := args.parseArgs(os.Args); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	fmt.Println(args)
}

// go run ./... -intflag 0x22 -stringflag admin -alphaflag 5
