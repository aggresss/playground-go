// Go语言实现了反射，所谓反射就是能检查程序在运行时的状态。
package main

import (
	"fmt"
	"reflect"
)

func main() {
	var x = 3.4
	v := reflect.ValueOf(x)
	fmt.Println("type:", v.Type())
	fmt.Println("kind is float64:", v.Kind() == reflect.Float64)
	fmt.Println("value:", v.Float())
}
