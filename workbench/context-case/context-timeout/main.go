package main

import (
	"context"
	"fmt"
	"time"
)

func main() {
	// ctx, cancel := context.WithCancel(context.Background())
	ctx, cancel := context.WithTimeout(context.Background(), 5000*time.Millisecond)
	go func(ctx context.Context) {
		for {
			select {
			case _, ok := <-ctx.Done():
				fmt.Println("task stopped", ok)
				return
			default:
				fmt.Println("task running")
				time.Sleep(time.Second * 2)
			}
		}
	}(ctx)

	time.Sleep(time.Second * 3)
	fmt.Println("task need stop")
	cancel()
	time.Sleep(time.Second * 3)
}
