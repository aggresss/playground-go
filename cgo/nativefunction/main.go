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

uint8_t buf[10] = {0x01,0x02,0x03,0x04,0x05,0x06,0x07,0x08,0x09,0x0A};

*/
import "C"
import (
	"fmt"
	"os"
	"unsafe"

	"golang.org/x/exp/constraints"
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

func PointerOffset[U constraints.Integer](ptr *uint8, offset U) *uint8 {
	// return (*uint8)(unsafe.Pointer(uintptr(unsafe.Pointer(ptr)) + uintptr(offset)))
	return (*uint8)(unsafe.Add(unsafe.Pointer(uintptr(unsafe.Pointer(ptr))), offset))
}

func main() {
	c, err := C.Add(C.int(1), C.int(2))
	fmt.Println(c, err)

	fmt.Println(os.Getenv("PKG_CONFIG_PATH"))

	fmt.Println(cuchar2Bytes((*uint8)(unsafe.Pointer(&C.buf)), 10))

	fmt.Println(cuchar2Bytes(PointerOffset((*uint8)(unsafe.Pointer(&C.buf)), 2), 10-2))
}
