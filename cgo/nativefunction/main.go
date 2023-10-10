package main

/*
// preamble
# include <stdio.h>
# include <stdint.h>
# include <stdlib.h>
# include <errno.h>

int Add(int a, int b) {
	errno = 1;
	return a+b;
}

uint8_t buf[10] = {1,2,3,4,5,6,7,8,9,0};

*/
import "C"
import (
	"fmt"
	"os"
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

// Convert *C.uchar to []byte
// use the memory directly without a copy.
func cuchar2Bytes(buf *uint8, size int) []byte {
	// https://pkg.go.dev/unsafe#Slice >=1.17
	return unsafe.Slice(buf, size)

	// https://pkg.go.dev/reflect#SliceHeader <1.17
	// s := reflect.SliceHeader{
	// 	Data: uintptr(unsafe.Pointer(buf)),
	// 	Len:  size,
	// 	Cap:  size,
	// }
	// return *((*[]byte)(unsafe.Pointer(&s)))

}

func main() {
	c, err := C.Add(C.int(1), C.int(2))
	fmt.Println(c, err)

	fmt.Println(os.Getenv("PKG_CONFIG_PATH"))

	fmt.Println(cuchar2Bytes((*uint8)(unsafe.Pointer(&C.buf)), 10))
}
