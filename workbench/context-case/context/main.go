package main

import (
	"context"
	"fmt"
	"time"
)

func main() {
	ctx, cancel := context.WithCancel(context.Background())
	go func(ctx context.Context) {
		for {
			select {
			case <-ctx.Done():
				fmt.Println("task stopped")
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