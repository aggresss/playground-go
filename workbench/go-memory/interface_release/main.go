package main

import (
	"fmt"
)

type myInterface interface {
	Print()
}

type myStruct struct {
	seq int
}

func (m *myStruct) Print() {
	fmt.Println(m.seq)
}

func newMyStruct(seq int) *myStruct {
	return &myStruct{
		seq: seq,
	}
}

func main() {
	mSet := make(map[int]myInterface)

	for i := 1; i < 10; i++ {
		mSet[i] = newMyStruct(i)
	}

	test01 := mSet[5]

	delete(mSet, 5)

	test01.Print()
}
