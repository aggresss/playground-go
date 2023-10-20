package main

/*
#include <stdlib.h>
#include <stdio.h>

void c_size_of(void* ptr) {
	printf("%lu\n", sizeof(*ptr));
}

*/
import "C"
import (
	"fmt"
	"unsafe"
)

func sizeofInterface(a any) {
	fmt.Println("sizeof interface", unsafe.Sizeof(a))
}

func sizeofGeneric[T any](a T) {
	fmt.Println("sizeof generic", unsafe.Sizeof(a))
}

func go_size_of[T any](p *T) uintptr {
	return unsafe.Sizeof(*p)
}

func main() {
	a := int32(4)
	fmt.Println(unsafe.Sizeof(a))
	C.c_size_of(unsafe.Pointer(&a))
	var p *int32
	fmt.Println(unsafe.Sizeof(*p))
	fmt.Println(go_size_of(p))

	sizeofInterface(a)
	sizeofGeneric(a)
}
