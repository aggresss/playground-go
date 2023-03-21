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
	t := &T{}
	tType := reflect.TypeOf(t)
	tValue := reflect.ValueOf(t)
	// use cache instead of MethodByName()
	cache := make(map[string]int)
	for i := 0; i < tValue.NumMethod(); i++ {
		if tType.Method(i).IsExported() {
			cache[tType.Method(i).Name] = i
		}
	}
	if i, ok := cache["Geeks"]; ok {
		tValue.Method(i).Call([]reflect.Value{reflect.ValueOf("admin")})
	}
}
