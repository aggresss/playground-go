package main

/*
#cgo LDFLAGS: -L. -lsum
#include "sum.h"
*/
import "C"
import "fmt"

func main() {
	a := C.int(10)
	b := C.int(20)
	fmt.Println(C.sum(a, b))
}
