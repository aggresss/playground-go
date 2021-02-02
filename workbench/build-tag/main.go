package main

import (
	"fmt"
	"os"
	"os/signal"
	"syscall"

	"github.com/aggresss/playground-go/workbench/build-tag/osname"
)

var Version string

func main() {
	ch := make(chan os.Signal, 1)
	signal.Notify(ch, syscall.SIGINT, syscall.SIGTERM, syscall.SIGKILL)

	fmt.Println(osname.OsName)
	fmt.Println(Version)

	select {
	case s := <-ch:
		fmt.Println(s.String())
	}
}
