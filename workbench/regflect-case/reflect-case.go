package main

import (
	"fmt"
	"reflect"
)

// Person test struct
type Person struct {
	Name string
	Age  int
}

// GetName method
func (p Person) GetName() string {
	return p.Name
}

// SetName method
func (p Person) SetName(s string) {
	p.Name = s
}
func main() {
	a := &Person{"Name", 1}

	t := reflect.TypeOf(*a) //必须取值，否则类型为空
	fmt.Println(t.Name())

	v := reflect.ValueOf(a).Elem() //a需要是引用
	k := v.Type()
	for i := 0; i < v.NumField(); i++ {
		key := k.Field(i)
		val := v.Field(i)
		fmt.Println(key.Name, val.Type(), val.Interface())
	}

	for i := 0; i < v.NumMethod(); i++ {
		key := k.Method(i)
		val := v.Method(i)
		fmt.Println(key.Name, val.Type(), val.Interface())
	}
	v.FieldByName("Name").Set(reflect.ValueOf("Name"))
	fmt.Println(a.Name)
	name := v.MethodByName("GetName").Call([]reflect.Value{})
	fmt.Println(name)
}
