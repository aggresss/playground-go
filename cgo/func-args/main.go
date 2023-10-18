package main

/*
#include <stdlib.h>
*/
import "C"
import "fmt"

func alpha(x [4]int) {
	for i := 0; i < 4; i++ {
		x[i] = i
	}
}

func beta(x []int) {
	fmt.Printf("beta input len %d\n", len(x))
	for i := 0; i < 4; i++ {
		x[i] = i
	}
}

func main() {
	var gamma [4]int
	fmt.Println(gamma)
	alpha(gamma)
	fmt.Println(gamma)
	beta(gamma[:])
	fmt.Println(gamma)
}

// In Go, arrays are passed to functions as values, not as references.
