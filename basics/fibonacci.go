// Package basics
// Time    : 2022/8/4 10:10
// Author  : xushiyin
// contact : yuqingxushiyin@gmail.com
package basics

import "fmt"

func Fibonacci(ch chan int, done chan struct{}) {
	x, y := 0, 1
	for {
		select {
		case ch <- x:
			x, y = y, x+y
		case <-done:
			fmt.Println("done")
			return
		}
	}
}
