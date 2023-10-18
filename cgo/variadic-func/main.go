package main

/*
#include <stdlib.h>
#include <stdio.h>
#include <stdarg.h>

void gamma(char* delta, ...) {
	char buf[1024] = {0};
	va_list vl;
	va_start (vl, delta);
	vsnprintf(buf, 1024, delta, vl);
	va_end(vl);

	printf("%s\n", buf);
}

void omicron() {
	gamma("psi, %s", "omega");
}

void theta(char* epsilon, char* eta) {
	printf("%s:%s\n", epsilon, eta);
}
*/
import "C"
import (
	"fmt"
	"unsafe"
)

func variadic1(prefix, body string, a ...any) {
	pPtr, pFunc := StringCasting(prefix)
	defer pFunc()
	sPtr, sFunc := StringCasting(fmt.Sprintf(body, a...))
	defer sFunc()
	C.theta(pPtr, sPtr)
}

func StringCasting(str string) (allocPtr *C.char, freeFunc func()) {
	if len(str) == 0 {
		return nil, func() {}
	}
	allocPtr = C.CString(str)
	freeFunc = func() { C.free(unsafe.Pointer(allocPtr)) }
	return allocPtr, freeFunc
}

func main() {
	C.omicron()
	variadic1("alpha", "beta&%s", "xi")
}
