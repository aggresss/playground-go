package main

/*
#cgo windows CFLAGS: -DCGO_OS_WINDOWS=1
#cgo darwin CFLAGS: -DCGO_OS_DARWIN=1
#cgo linux CFLAGS: -DCGO_OS_LINUX=1

#if defined(CGO_OS_WINDOWS)
    const char* os = "windows";
#elif defined(CGO_OS_DARWIN)
    const char* os = "darwin";
#elif defined(CGO_OS_LINUX)
    const char* os = "linux";
#else
#    error(unknown os)
#endif

#cgo !windows LDFLAGS: -L. -lsum
#include "sum.h"
*/
import "C"
import "fmt"

func main() {
	a := C.int(10)
	b := C.int(20)
	fmt.Println(C.sum(a, b))
}
