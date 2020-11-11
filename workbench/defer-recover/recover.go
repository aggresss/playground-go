package main

import (
	"fmt"
	"runtime"
	"time"
)

func testFunc() {
	defer func() {
		if e := recover(); e != nil {
			const size = 16 << 10
			buf := make([]byte, size)
			buf = buf[:runtime.Stack(buf, false)]
			fmt.Println("process panic: ", e, fmt.Sprintf("\n%s", buf))
		}
	}()

	var testMap map[string]string

	testMap["foo"] = "bar"
}

func main() {

	testFunc()

	select {
	case <-time.After(time.Second * time.Duration(5)):
		break
	}
}
