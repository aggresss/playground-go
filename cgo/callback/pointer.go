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

func PointerStore(v interface{}) unsafe.Pointer {
	if v == nil {
		return nil
	}
	var ptr unsafe.Pointer = C.malloc(C.size_t(1))
	if ptr == nil {
		panic("allocate memory faild")
	}
	m.Store(ptr, v)
	return ptr
}

func PointerLoad(ptr unsafe.Pointer) (v interface{}) {
	if ptr == nil {
		return nil
	}
	v, _ = m.Load(ptr)
	return v
}

func PointerDelete(ptr unsafe.Pointer) {
	if ptr == nil {
		return
	}
	m.Delete(ptr)
	C.free(ptr)
}
