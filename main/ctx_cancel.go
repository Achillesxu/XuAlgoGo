// Package main
// Time    : 2022/9/18 12:17
// Author  : xushiyin
// contact : yuqingxushiyin@gmail.com
package main

import (
	"context"
	"fmt"
	"time"
)

func longRunningTask(ctx context.Context) {
	fmt.Println("background long running task launched")
	select {
	case <-ctx.Done():
		fmt.Println("long running task bailed because context cancelled")
	}
}

// it will not run goroutine if main thread run too faster

func main() {
	// this will bail when cancelFunc is called
	ctx, cancelFunc := context.WithCancel(context.Background())
	go longRunningTask(ctx)

	time.Sleep(1 * time.Second)

	fmt.Println("background long running task still going")
	time.Sleep(1 * time.Second)

	fmt.Println("going to cancel background task")
	cancelFunc()

	time.Sleep(1 * time.Second)
	fmt.Println("some time has elapsed after cancelling")
}
