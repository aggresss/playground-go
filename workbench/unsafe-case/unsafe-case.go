package main

import (
	"fmt"
	"unsafe"
)

// b2s converts byte slice to a string without memory allocation.
//
// Note it may break if string and/or slice header will change
// in the future go versions.
func b2s(b []byte) string {
	return *(*string)(unsafe.Pointer(&b)) // #nosec
}

func main() {
	var a = []byte("abcde")
	b := b2s(a)
	fmt.Printf("%s\n", b)
}
