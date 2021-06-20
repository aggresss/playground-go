// https://stackoverflow.com/questions/28501976/embedding-in-go-with-pointer-or-with-value

package main

import (
	"fmt"
)

type Base struct {
	Name string
}

func (b Base) PrintName() {
	fmt.Println(b.Name)
}

func (b *Base) PrintNameP() {
	fmt.Println(b.Name)
}

func (b Base) ChangeName(name string) {
	b.Name = name
}

func (b *Base) ChangeNameP(name string) {
	b.Name = name
}

type EmbedsBase struct {
	Base
}

type EmbedsPointerToBase struct {
	*Base
}

func main() {

	fmt.Println("")
	fmt.Println("# embed by value and refrenced by value, not change origianl value")
	b := Base{"Jeff Hardy"}
	eb := EmbedsBase{b}
	eb.PrintName()
	eb.ChangeName("John Cena")
	eb.PrintName()

	fmt.Println("")
	fmt.Println("# embed by value, but refrenced by pointer, changed origianl value")
	b = Base{"Jeff Hardy"}
	ebp := &EmbedsBase{b}
	ebp.PrintNameP()
	ebp.ChangeNameP("John Cena")
	ebp.PrintNameP()

	fmt.Println("")
	fmt.Println("# embed by pointer, but refrenced by value, not chage origianl value")
	b = Base{"Jeff Hardy"}
	epb := EmbedsPointerToBase{&b}
	epb.PrintName()
	epb.ChangeName("John Cena")
	epb.PrintName()

	fmt.Println("")
	fmt.Println("# embed by pointer, and refrenced by pointer, changed origianl value")
	b = Base{"Jeff Hardy"}
	epbp := &EmbedsPointerToBase{&b}
	epbp.PrintNameP()
	epbp.ChangeNameP("John Cena")
	epbp.PrintNameP()
}
