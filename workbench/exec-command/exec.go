// https://golang.org/pkg/os/exec/
package main

import (
	"fmt"
	"io"
	"os"
	"os/exec"
	"os/signal"
	"syscall"
	"time"
)

func main() {
	cmd := exec.Command("./hello")

	stdin, _ := cmd.StdinPipe()
	cmd.Start()

	go func() {
		//defer stdin.Close()
		if stdin != nil {
			io.WriteString(stdin, "values written to stdin are passed to cmd's standard input")
		}
		stat, _ := cmd.Process.Wait()
		fmt.Printf("cmd quit: pid=%d stat=%v", stat.Pid(), stat.String())
	}()

	time.Sleep(time.Second * 60)
	cmd.Process.Kill()
	if stdin != nil {
		stdin.Close()
	}

	ch := make(chan os.Signal, 1)
	signal.Notify(ch, syscall.SIGINT, syscall.SIGTERM, syscall.SIGKILL)
	fmt.Println("Wait signal")
	select {
	case s := <-ch:
		fmt.Println(s.String())
	}
}
