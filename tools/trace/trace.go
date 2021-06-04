package main

import (
	"fmt"
	"os"
	"os/signal"
	"runtime/trace"
	"syscall"
)

func main() {
	trace.Start(os.Stderr)
	defer trace.Stop()

	ch := make(chan os.Signal, 1)
	signal.Notify(ch, os.Interrupt, syscall.SIGTERM)
	select {
	case s := <-ch:
		fmt.Println(s.String())
	}
}
