package main

import "fmt"

func main() {
	f()
	fmt.Println("e")
}

func f() {
	defer func() { // 必须要先声明defer，否则不能捕获到panic异常
		fmt.Println("c")
		if err := recover(); err != nil {
			fmt.Println(err) // 这里的err是panic传入的内容
		}
		fmt.Println("d")
	}()

	fmt.Println("a")
	fmt.Println("b")
	panic("f")
}
