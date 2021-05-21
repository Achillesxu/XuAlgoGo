// Package XuAlgoGo
// Time    : 2021/5/6 2:10 下午
// Author  : xushiyin
// contact : yuqingxushiyin@gmail.com
package main

import (
	"container/ring"
	"fmt"
)

// main project entrance
func main() {
	var integers = []int{1, 3, 5, 7}
	var ringList *ring.Ring
	ringList = ring.New(len(integers))
	for i := 0; i < ringList.Len(); i++ {
		ringList.Value = integers[i]
		ringList = ringList.Next()
	}
	ringList = ringList.Next()
	ringList.Do(func(p interface{}) {
		fmt.Println(p.(int))
	})

}
