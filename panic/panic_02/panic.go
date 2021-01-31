package main

import (
	"fmt"
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

func panic_02() {
	var a *A = nil
	a.F1(1)
	a.F2(2)
}

func main() {
	panic_02()
}
