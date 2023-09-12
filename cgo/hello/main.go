package main

/*
#cgo LDFLAGS: -L. -lsum
#include "sum.h"
*/
import "C"
import (
	"fmt"
	"unsafe"
)

func main() {
	a := C.int(10)
	b := C.int(20)
	var c C.SUM_POINTER
	C.sum(&c, a, b)
	fmt.Println(*c)
	C.free(unsafe.Pointer(c))
}
