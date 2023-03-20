// Golang program to illustrate
// reflect.Call() Function

package main

import (
	"fmt"
	"reflect"
)

type T struct{}

func (t *T) Geeks(foo string) {
	fmt.Println("Geeks:", foo)
}

func main() {
	tType := reflect.TypeOf(&T{})
	tValue := reflect.ValueOf(&T{})
	// use cache instead of MethodByName()
	cache := make(map[string]reflect.Value)
	for i := 0; i < tValue.NumMethod(); i++ {
		if tType.Method(i).IsExported() {
			cache[tType.Method(i).Name] = tValue.Method(i)
		}
	}
	if m, ok := cache["Geeks"]; ok {
		m.Call([]reflect.Value{reflect.ValueOf("admin")})
	}
}
