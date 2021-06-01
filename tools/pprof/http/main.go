package main

import (
	"fmt"
	"os"
	"os/signal"
	"syscall"
)

func main() {
	debugS, err := Start("127.0.0.1:6061")
	if err != nil {
		fmt.Println(err.Error())
	}

	ch := make(chan os.Signal, 1)
	signal.Notify(ch, os.Interrupt, syscall.SIGTERM)

	select {
	case s := <-ch:
		fmt.Println(s.String())
	}

	if debugS != nil {
		debugS.Close()
	}
}
