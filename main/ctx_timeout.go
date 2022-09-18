// Package main
// Time    : 2022/9/18 12:30
// Author  : xushiyin
// contact : yuqingxushiyin@gmail.com
package main

import (
	"context"
	"fmt"
	"time"
)

func longRunningTask2(ctx context.Context, timeToRun time.Duration) {
	select {
	case <-time.After(timeToRun):
		fmt.Println("completed before context timed out")
	case <-ctx.Done():
		fmt.Println("bailed because context timed out")
	}
}

const timeout = 5 * time.Second

func main() {
	ctx := context.Background()

	// this will bail because the function takes longer than the context allows
	ctx1, _ := context.WithTimeout(ctx, timeout)
	longRunningTask2(ctx1, 10*time.Second)

	// this will complete because the function completes before the context times out
	ctx2, _ := context.WithTimeout(ctx, timeout)
	longRunningTask2(ctx2, 1*time.Second)
}
