package main

/*
#include <stdlib.h>
#include <stdio.h>

void my_size_of(void* ptr) {
	printf("%lu\n", sizeof(*ptr));
}

*/
import "C"
import (
	"fmt"
	"unsafe"
)

func main() {
	a := int16(4)
	fmt.Println(unsafe.Sizeof(a))
	C.my_size_of(unsafe.Pointer(&a))

	var p *int16
	fmt.Println(unsafe.Sizeof(*p))
}
