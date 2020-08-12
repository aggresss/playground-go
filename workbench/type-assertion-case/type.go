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
		fmt.Printf("unexpected type %T\n", t)
		break
	case bool:
		fmt.Printf("boolean %t\n", t)
		break
	case int:
		fmt.Printf("integer %d\n", t)
		break
	case *bool:
		fmt.Printf("pointer to boolean %t\n", *t)
		break
	case *int:
		fmt.Printf("pointer to integer %d\n", *t)
		break
	}
}
