// Package basics
// Time    : 2022/8/4 10:15
// Author  : xushiyin
// contact : yuqingxushiyin@gmail.com
package basics

import (
	"fmt"
	"testing"
)

func TestFibonacci(t *testing.T) {
	ch := make(chan int)
	done := make(chan struct{})
	go func() {
		for i := range ch {
			if i > 10 {
				done <- struct{}{}
			}
			fmt.Println(i)
		}
	}()
	Fibonacci(ch, done)
	close(done)
	close(ch)
}
