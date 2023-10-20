package main

/*
#include <stdlib.h>
#include <stdio.h>

void checknull(void **v) {
	if (!v) {
		printf("check NULL ok\n");
	}
}

*/
import "C"
import (
	"fmt"
	"reflect"
	"unsafe"
)

// CVoidPointer represents a (void*) type pointer in the C world.
type CVoidPointer any

// VoidPointer returns a unsafe.Pointer from CVoidPointer.
func VoidPointer(a CVoidPointer) unsafe.Pointer {
	if a == nil {
		return unsafe.Pointer(nil)
	}
	return unsafe.Pointer(reflect.ValueOf(a).Pointer())
}

// CVoidPointer represents a (void**) type pointer in the C world.
type CVoidPointerPointer any

// VoidPointer returns a *unsafe.Pointer from CVoidPointerPointer.
func VoidPointerPointer(a CVoidPointerPointer) *unsafe.Pointer {
	if a == nil {
		return (*unsafe.Pointer)(nil)
	}
	return (*unsafe.Pointer)(unsafe.Pointer(reflect.ValueOf(a).Pointer()))
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

	fmt.Println(reflect.ValueOf(nil).Kind())

	C.checknull(VoidPointerPointer(nil))

	// type unsafe.Pointer of unsafe.Pointer(&i) does not match *T (cannot infer T)
	//fmt.Println(GenericPointer(unsafe.Pointer(&i)))

}
