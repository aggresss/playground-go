package main

import (
	"fmt"
	"runtime"
)

func main() {
	defer func() {
		if err := recover(); err != nil {
			const size = 64 << 10
			buf := make([]byte, size)
			buf = buf[:runtime.Stack(buf, false)]
			fmt.Println(string(buf))
		}
	}()
	panic("expect panic")
}
