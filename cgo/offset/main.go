package main

/*
#include <stdlib.h>

char *lambda[] = {"Monday", "Tuesday", "Wednesday", "Thursday", "Friday", "Saturday", "Sunday", 0x00};
char upsilon[][10] = {"Monday", "Tuesday", "Wednesday", "Thursday", "Friday", "Saturday", "Sunday", 0x00};

*/
import "C"
import (
	"fmt"
	"unsafe"
)

type HelperInteger interface {
	HelperSingedInteger | HelperUnsingedInteger
}

type HelperSingedInteger interface {
	~int | ~int8 | ~int16 | ~int32 | ~int64
}

type HelperUnsingedInteger interface {
	~uint | ~uint8 | ~uint16 | ~uint32 | ~uint64 | ~uintptr
}

// PointerOffset offset the pointer point.
func PointerOffset[U any, V HelperInteger](ptr *U, offset V) *U {
	if ptr == nil {
		return nil
	}
	return (*U)(unsafe.Pointer(uintptr(unsafe.Pointer(ptr)) +
		uintptr(unsafe.Sizeof(*ptr))*(uintptr(offset))))
}

// TruncSlice return a slice from a sign-terminated array.
func TruncSlice[T any](ptr *T, fn func(T) bool) []T {
	if ptr == nil {
		return nil
	}
	for i := 0; ; i++ {
		if fn(*(*T)(unsafe.Pointer(uintptr(unsafe.Pointer(ptr)) +
			uintptr(unsafe.Sizeof(*ptr))*uintptr(i)))) {
			return unsafe.Slice(ptr, i)
		}
	}
}

// TruncStringSlice returns a string slice from a NULL-terminated *C.char array.
func TruncStringSlice(ptr **C.char) (v []string) {
	if ptr == nil {
		return nil
	}
	for *ptr != nil && **ptr != C.char(0x00) {
		v = append(v, C.GoString(*ptr))
		ptr = (**C.char)(unsafe.Pointer(uintptr(unsafe.Pointer(ptr)) +
			uintptr(unsafe.Sizeof(*ptr))))
	}
	return v
}

func main() {
	alpha := []int64{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}

	beta := &alpha[0]
	fmt.Println(*beta)
	beta = PointerOffset(beta, 2)
	fmt.Println(*beta)

	gamma := TruncSlice(beta, func(i int64) bool {
		return i == 6
	})
	fmt.Println(gamma)

	fmt.Println(TruncStringSlice((**C.char)(&C.lambda[0])))

}

var (
	alphabet = []string{
		"alpha", "bata", "gamma", "delta", "epsilon", "zeta", "eta", "theta",
		"iota", "kappa", "lambda", "mu", "nu", "xi", "omicron", "pi",
		"rho", "sigma", "tau", "upsilon", "phi", "chi", "psi", "omega",
	}
)
