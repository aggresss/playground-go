package main

/*
#include <stdlib.h>
#include <unistd.h>

typedef int* intp;

typedef int (callback)(void*, char*, int);

static void call_later(int delay, callback *cb, void* data) {
  sleep(delay);
  char* alpha = "beta";
  int gamma = 5;
  cb(data, alpha, gamma);
}

int go_cb(void*, char*, int);

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
}

//export go_cb
func go_cb(a unsafe.Pointer, b *CCHAR, c C.int) C.int {
	f := (PointerLoad(a)).(*error)
	fmt.Println(*f, b, c)
	return 0
}
