package main

import (
	"fmt"
	"reflect"
	"strconv"
	"strings"
)

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
		var tn string
		if tn = fd.Tag.Get("json"); len(tn) == 0 {
			continue
		}
		var tv string
		switch fd.Type.Kind() {
		case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64:
			tv = strconv.Itoa(int(vo.Field(i).Uint()))
			if tv == "0" {
				continue
			}
		case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
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
