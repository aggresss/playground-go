package main

import (
	"fmt"
	"strconv"
	"sync"
	"time"
)

var wg sync.WaitGroup

func run(task string) {
	fmt.Println(task, "start ...")
	time.Sleep(time.Second * 2)
	wg.Done()
}

func main() {
	wg.Add(2)
	for i := 1; i < 3; i++ {
		taskName := "task" + strconv.Itoa(i)
		go run(taskName)
	}

	wg.Wait()
	fmt.Println("All task complete.")
}
