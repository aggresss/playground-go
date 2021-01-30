package main

import (
	"fmt"
	"os"
	"runtime"
)

type A struct {
	Num int
}

func (a *A) F1(c int) {
	fmt.Println(c, a.Num)
}

func (a *A) F2(c int) {
	fmt.Println(c, a)
}

func printStack() {
	var buf [4096]byte
	n := runtime.Stack(buf[:], true)
	fmt.Printf("==> %s\n", string(buf[:n]))
}

func panic_02() {
	var a *A = nil
	a.F1(1)
	a.F2(2)
}

func main() {
	defer func() {
		if e := recover(); e != nil {
			printStack()
			os.Exit(1)
		}
	}()

	panic_02()
}
