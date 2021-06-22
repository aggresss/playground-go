package main

import (
	"context"
	"fmt"
	"time"
)

func ctx_fn(ctx context.Context, fn func(ctx context.Context) error) error {
	newctx, cancel := context.WithCancel(ctx)
	go func() {
		<-newctx.Done()
		cancel()
	}()
	if err := fn(newctx); err != nil {
		cancel()
		return err
	}
	return nil
}

func main() {
	ctx := context.Background()
	alphaCtx, alphaCancel := context.WithCancel(ctx)
	defer alphaCancel()
	betaCtx, betaCancel := context.WithCancel(alphaCtx)
	defer betaCancel()

	go func() {
		select {
		case <-time.After(time.Second):
			alphaCancel()
		}
	}()

	<-betaCtx.Done()
	fmt.Println("beta ctx done")

}
