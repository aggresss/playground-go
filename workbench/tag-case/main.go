package main

import (
	"fmt"
	"reflect"
	"strconv"
	"strings"
)

type tho struct {
	Epsilon int    `json:"epsilon"`
	Zeta    string `json:"zeta"`
}

func transJsonTag(val interface{}) string {
	ret := []string{}

	vo := reflect.ValueOf(val)
	if vo.Kind() == reflect.Ptr {
		vo = vo.Elem()
	}
	if vo.Kind() != reflect.Struct {
		return ""
	}
	for i := 0; i < vo.NumField(); i++ {
		fd := vo.Type().Field(i)
		tn := fd.Tag.Get("json")
		var tv string
		switch fd.Type.Kind() {
		case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64,
			reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64:
			tv = strconv.Itoa(int(vo.Field(i).Int()))
			if tv == "0" {
				continue
			}
		case reflect.String:
			tv = vo.Field(i).String()
			if len(tv) == 0 {
				continue
			}
		}
		ret = append(ret, fmt.Sprintf("%s=%s", tn, tv))
	}

	return strings.Join(ret, ";")
}

func main() {
	t := &tho{
		Epsilon: 99,
		Zeta:    "0",
	}
	fmt.Println(transJsonTag(t))
}
