package main

/*
#include <stdlib.h>
#include <stdio.h>

typedef struct {
	char *msg;
} myStruct;

void myFunc(myStruct *st)
{
	printf("Hello %s!\n", st->msg);
}
*/
import "C"
import (
	"runtime"
	"unsafe"
)

func main() {
	st := C.myStruct{C.CString("world")}
	runtime.SetFinalizer(&st, func(t *C.myStruct) {
		C.free(unsafe.Pointer(t.msg))
	})
	C.myFunc(&st)
	runtime.KeepAlive(&st)
}

// panic: runtime error: cgo argument has Go pointer to unpinned Go pointer
