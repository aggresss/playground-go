package main

import (
	"context"
	"fmt"
	"math/rand"
	"time"
)

func main() {
	ctx, ctxCancel := context.WithTimeout(context.TODO(), time.Millisecond*100)
	defer ctxCancel()
	if err := request01(ctx); err != nil {
		fmt.Println(err)
	}
}

func request01(ctx context.Context) error {
	select {
	case <-ctx.Done():
		return ctx.Err()
	case <-time.After(time.Millisecond * time.Duration(rand.Intn(100))):
		return request02(ctx)
	}
}

func request02(ctx context.Context) error {
	select {
	case <-ctx.Done():
		return ctx.Err()
	case <-time.After(time.Millisecond * time.Duration(rand.Intn(100))):
		return nil
	}
}
