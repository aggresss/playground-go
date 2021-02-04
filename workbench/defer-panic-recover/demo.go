package main

import "fmt"

func main() {
	f()
	fmt.Println("e")
}

func f() {
	defer func() {
		fmt.Println("c")
		if err := recover(); err != nil {
			fmt.Println(err)
		}
		fmt.Println("d")
	}()

	fmt.Println("a")
	fmt.Println("b")
	panic("f")
}
