package main

/*
#include <stdlib.h>
#include <unistd.h>
#include <stdio.h>

typedef int (callback)(void*, char*, int);

static void call_later(int delay, callback *cb, void* data) {
  sleep(delay);
  char* alpha = "beta";
  int gamma = 5;
  cb(data, alpha, gamma);
}

extern int go_cb(void*, char*, size_t);


// https://zchee.github.io/golang-wiki/cgo/#export-and-definition-in-preamble
// normally you will have to define function or variables
// in another separate C file to avoid the multiple definition
// errors, however, using "static inline" is a nice workaround
// for simple functions like this one.

static int hw_pix_fmt = 5;

static inline int get_hw_pix_fmt() {
	return hw_pix_fmt;
}

static inline void set_hw_pix_fmt(int fmt) {
	hw_pix_fmt = fmt;
}

*/
import "C"
import (
	"errors"
	"fmt"
	"unsafe"
)

type LocalCallback C.callback

// faild: type LocalImitate func(unsafe.Pointer, *C.char, C.int) C.int

// panic: runtime error: cgo argument has Go pointer to unpinned Go pointer

// but stack vavirant address ok.

type CCHAR C.char

func main() {
	data := errors.New("omicron")
	C.call_later(1, (*C.callback)(C.go_cb), PointerStore(&data))
	// ok: C.call_later(3, LocalCallback(C.go_cb))

	C.set_hw_pix_fmt(6)
	fmt.Println(C.get_hw_pix_fmt())

}

//export go_cb
func go_cb(a unsafe.Pointer, b *CCHAR, c uintptr) C.int {
	f := (PointerLoad(a)).(*error)
	fmt.Println(*f, b, c)
	return 0
}
