package main

import (
	"fmt"
	"time"
)

func main() {
	stop := make(chan bool)
	go func() {
		for {
			select {
			case <-stop:
				fmt.Println("task stopped")
				return
			default:
				fmt.Println("task running")
				time.Sleep(time.Second * 2)
			}
		}
	}()

	time.Sleep(time.Second * 3)
	fmt.Println("task need stop")
	stop <- true
	time.Sleep(time.Second * 3)
}
