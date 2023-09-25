package main

/*
// preamble
# include <stdio.h>
# include <errno.h>

int Add(int a, int b) {
	errno = 1;
	return a+b;
}
*/
import "C"
import (
	"fmt"
	"unsafe"
)

// C.GoBytes(unsafe.pointer, C.int) []byte
func dumpArray(array *C.uchar, length int) []byte {
	buf := make([]byte, length)
	for i := 0; i < length; i++ {
		buf[i] = *(*byte)(unsafe.Pointer(uintptr(unsafe.Pointer(array)) +
			uintptr(i*int(unsafe.Sizeof(*array)))))
	}
	return buf
}

func main() {
	c, err := C.Add(C.int(1), C.int(2))
	fmt.Println(c, err)
}
