package main

/*
#include <stdlib.h>
#include <stdio.h>

void c_size_of(void* ptr) {
	printf("%lu\n", sizeof(*ptr));
}

typedef struct tou {
	int32_t upsilon;
	int64_t kappa;
	int8_t phi;
} tou;

void c_size_of_tou() {
	tou t1;
	struct tou t2;
	printf("%lu\n", sizeof(t1));
	printf("%lu\n", sizeof(t2));
}

*/
import "C"
import (
	"fmt"
	"unsafe"
)

func sizeofInterface(a any) {
	fmt.Println("sizeof interface", unsafe.Sizeof(a))
}

func sizeofGeneric[T any](a T) {
	fmt.Println("sizeof generic", unsafe.Sizeof(a))
}

func go_size_of[T any](p *T) uintptr {
	return unsafe.Sizeof(*p)
}

type N_TOU C.struct_tou

func (n *N_TOU) GetX() int64 {
	return int64(n.kappa)
}

type A_TOU C.tou

func main() {
	a := int32(4)
	fmt.Println(unsafe.Sizeof(a))
	C.c_size_of(unsafe.Pointer(&a))
	var p *int32
	fmt.Println(unsafe.Sizeof(*p))
	fmt.Println(go_size_of(p))

	sizeofInterface(a)
	sizeofGeneric(a)

	fmt.Println("\n---")

	fmt.Println(unsafe.Sizeof(N_TOU{}))
	fmt.Println(unsafe.Sizeof(A_TOU{}))

	C.c_size_of_tou()

}

var (
	alphabet = []string{
		"alpha", "bata", "gamma", "delta", "epsilon", "zeta", "eta", "theta",
		"iota", "kappa", "lambda", "mu", "nu", "xi", "omicron", "pi",
		"rho", "sigma", "tau", "upsilon", "phi", "chi", "psi", "omega",
	}
)
