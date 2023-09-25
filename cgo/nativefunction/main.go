package main

/*
// preamble
int Add(int a, int b) {
	return a+b;
}
*/
import "C"
import (
	"fmt"
	"unsafe"
)

func dumpArray(array *C.uchar, length int) []byte {
	buf := make([]byte, length)
	for i := 0; i < length; i++ {
		buf[i] = *(*byte)(unsafe.Pointer(uintptr(unsafe.Pointer(array)) +
			uintptr(i*int(unsafe.Sizeof(*array)))))
	}
	return buf
}

func main() {
	c := C.Add(C.int(1), C.int(2))
	fmt.Println(c)
}
