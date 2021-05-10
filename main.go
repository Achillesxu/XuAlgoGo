// Package XuAlgoGo
// Time    : 2021/5/6 2:10 下午
// Author  : xushiyin
// contact : yuqingxushiyin@gmail.com
package main

import (
	"fmt"
	lds "github.com/Achillesxu/XuAlgoGo/linear_data_structure"
)

// main project entrance
func main() {
	var linkedList lds.LinkedList
	linkedList = lds.LinkedList{}
	linkedList.AddToHead(1)
	linkedList.AddToHead(3)
	linkedList.AddToHead(9)

	linkedList.AddBefore(1, 15)

	linkedList.IterateList()
	fmt.Println(linkedList.LinkedListLen())

	// fmt.Println(linkedList.LinkedListLen())
}
