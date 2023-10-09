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

type Alpha C.struct_Alpha

const (
	OMIGA = int(C.OMIGA)
)

func main() {
	a := C.int(10)
	b := C.int(20)
	var c C.SUM_POINTER
	C.sum(&c, a, b)
	fmt.Println(*c)
	C.free(unsafe.Pointer(c))

	C.print_struct_p((*C.struct_Alpha)(&Alpha{
		beta:  1,
		gamma: 2.0,
	}))

	C.print_struct((C.struct_Alpha)(Alpha{
		beta:  3,
		gamma: 4.0,
	}))

	d := C.return_struct()
	fmt.Println(Alpha(d))

	e := C.return_struct_point()
	fmt.Println(e)
}
