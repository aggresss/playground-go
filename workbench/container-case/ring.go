package main

import (
	"container/ring"
	"fmt"
)

func main() {
	RingFunc()
}

func RingFunc() {
	r1 := ring.New(10) //初始长度10
	for i := 0; i < r1.Len(); i++ {
		r1.Value = i
		r1 = r1.Next()
	}

	r2 := ring.New(10) //初始长度10
	for i := 10; i < r2.Len()+10; i++ {
		r2.Value = i
		r2 = r2.Next()
	}

	r := r1.Link(r2)
	for i := 0; i < 22; i++ {
		fmt.Println(r.Value)
		r = r.Next()
	}
}
