// Package main
// Time    : 2022/9/18 12:23
// Author  : xushiyin
// contact : yuqingxushiyin@gmail.com
package main

import (
	"context"
	"fmt"
	"time"
)

func longRunningTask1(ctx context.Context, timeToRun time.Duration) {
	select {
	case <-time.After(timeToRun):
		fmt.Println("completed before context deadline passed")
	case <-ctx.Done():
		fmt.Println("bailed because context deadline passed")
	}
}

const duration = 5 * time.Second

func main() {
	ctx := context.Background()

	// this will bail because the function runs longer than the context's deadline allows
	ctx1, _ := context.WithDeadline(ctx, time.Now().Add(duration))
	longRunningTask1(ctx1, 10*time.Second)

	// this will complete because the function completes before the context's deadline arrives
	ctx2, _ := context.WithDeadline(ctx, time.Now().Add(duration))
	longRunningTask1(ctx2, 1*time.Second)
}
