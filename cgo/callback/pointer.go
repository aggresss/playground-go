package main

// Reference: https://github.com/mattn/go-pointer

/*
#include <stdlib.h>
*/
import "C"
import (
	"sync"
	"unsafe"
)

var (
	m sync.Map
)

func Save(v interface{}) unsafe.Pointer {
	if v == nil {
		return nil
	}
	var ptr unsafe.Pointer = C.malloc(C.size_t(1))
	if ptr == nil {
		panic("can't allocate 'cgo-pointer hack index pointer': ptr == nil")
	}
	m.Store(ptr, v)
	return ptr
}

func Restore(ptr unsafe.Pointer) (v interface{}) {
	if ptr == nil {
		return nil
	}
	v, _ = m.Load(ptr)
	return v
}

func Unref(ptr unsafe.Pointer) {
	if ptr == nil {
		return
	}
	m.Delete(ptr)
	C.free(ptr)
}
