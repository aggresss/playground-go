package main

import (
	"fmt"
	"os"
	"os/signal"
	"syscall"

	"github.com/aggresss/playground-go/workbench/build-tag/hello"
)

func main() {
	ch := make(chan os.Signal, 1)
	signal.Notify(ch, syscall.SIGINT, syscall.SIGTERM, syscall.SIGKILL)

	fmt.Println(hello.OsName)

	select {
	case s := <-ch:
		fmt.Println(s.String())
	}
}
