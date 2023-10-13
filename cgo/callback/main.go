package main

/*
#include <stdlib.h>


#include <unistd.h>

typedef void (*callback)(void*);

static void call_later(int delay, callback cb) {
  sleep(delay);
  char* alpha = "beta";
  int gamma = 5;
  cb(alpha);
}

void go_cb(void*);

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
func go_cb(a unsafe.Pointer) {
	fmt.Println("admin")
}
