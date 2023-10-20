package main

/*
#include <stdlib.h>
*/
import "C"
import (
	"fmt"
	"reflect"
	"unsafe"
)

func VoidPointer(a any) unsafe.Pointer {
	return unsafe.Pointer(reflect.ValueOf(a).Pointer())
}

// Not support unsafe.poiter as input
func GenericPointer[T any](a *T) unsafe.Pointer {
	return unsafe.Pointer(a)
}

func main() {
	var i = 1

	fmt.Println(VoidPointer(&i))
	fmt.Println(unsafe.Pointer(VoidPointer(&i)))
	fmt.Println(&i)
	fmt.Println(i)

	fmt.Println(GenericPointer(&i))

	// type unsafe.Pointer of unsafe.Pointer(&i) does not match *T (cannot infer T)
	//fmt.Println(GenericPointer(unsafe.Pointer(&i)))

}
