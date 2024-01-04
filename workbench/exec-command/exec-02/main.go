package main

import (
	"fmt"
	"os"
	"os/exec"
	"os/signal"
	"strconv"
	"syscall"
	"time"
)

func main() {
	var (
		concurrent  int = 1
		err         error
		cmds        []*exec.Cmd
		execCommand = []string{"ffmpeg", "-re", "-y",
			"-i", "rtmp://127.0.0.1:1935/live/test?domain=www.test.com",
			"-c:a", "copy", "-c:v", "copy",
			"-f", "flv", "/dev/null",
		}
	)

	// parse flag
	if len(os.Args) > 1 {
		if concurrent, err = strconv.Atoi(os.Args[1]); err != nil {
			fmt.Println("need a integer arg")
			os.Exit(1)
		}
	}

	cmds = make([]*exec.Cmd, concurrent)

	// run command
	for i := 0; i < concurrent; i++ {
		go func(i int) {
			cmds[i] = exec.Command(execCommand[0], execCommand[1:]...)
			cmds[i].Start()
			stat, _ := cmds[i].Process.Wait()
			fmt.Printf("cmd quit: pid=%d stat=%v\n", stat.Pid(), stat.String())
		}(i)
		time.Sleep(time.Millisecond * 100)
	}

	// wait stop
	ch := make(chan os.Signal, 1)
	signal.Notify(ch, syscall.SIGINT, syscall.SIGTERM)
	fmt.Println("Wait signal")
	fmt.Println((<-ch).String())

	// stop command
	for i := 0; i < concurrent; i++ {
		cmds[i].Process.Kill()
		time.Sleep(time.Millisecond * 100)
	}
}
