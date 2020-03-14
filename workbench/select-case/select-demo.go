package main

import (
	"fmt"
	"time"
)

func chanTest(ch chan int) {

	for {
		select {
		case value, ok := <-ch:
			fmt.Println(value, ok, time.Now())
			if ok == false {
				//select要自己判断退出，如果是for..range 形式，在读取完了关闭的chanel后，退出循环
				fmt.Println("chan closed", time.Now())
				return
			}

		default:
			fmt.Println("chan empty", time.Now())
			//分支的处理会阻塞整个select
			time.Sleep(time.Second * 1)
		}
	}
}

func main() {

	var ch = make(chan int, 100)

	go chanTest(ch)

	ch <- 1
	ch <- 2
	time.Sleep(time.Second * 2)
	ch <- 3
	ch <- 4

	time.Sleep(time.Second)

	close(ch)

	for {
		time.Sleep(time.Second)
	}
}
