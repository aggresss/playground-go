package main

import (
	"fmt"
)

func functionOfSomeType() interface{} {
	return true
}

func main() {
	var t interface{}
	t = functionOfSomeType()
	switch t := t.(type) {
	default:
		fmt.Printf("unexpected type %T", t) // %T prints whatever type t has
		break
	case bool:
		fmt.Printf("boolean %t\n", t) // t has type bool
		break
	case int:
		fmt.Printf("integer %d\n", t) // t has type int
		break
	case *bool:
		fmt.Printf("pointer to boolean %t\n", *t) // t has type *bool
		break
	case *int:
		fmt.Printf("pointer to integer %d\n", *t) // t has type *int
		break
	}
}
