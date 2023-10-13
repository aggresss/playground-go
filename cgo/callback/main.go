package main

/*
#include <stdlib.h>


#include <unistd.h>

typedef int (*callback)(void*, char*, int);

static void call_later(int delay, callback cb) {
  sleep(delay);
  char* alpha = "beta";
  int gamma = 5;
  cb(NULL, alpha, gamma);
}

int go_cb(void*, char*, int);

*/
import "C"
import (
	"fmt"
	"unsafe"
)

func main() {
	C.call_later(3, C.callback(C.go_cb))
}

//export go_cb
func go_cb(a unsafe.Pointer, b *C.char, c C.int) C.int {
	fmt.Println(a, b, c)
	return 0
}
