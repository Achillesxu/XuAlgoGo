// Package XuAlgoGo
// Time    : 2021/5/6 2:10 下午
// Author  : xushiyin
// contact : yuqingxushiyin@gmail.com
package main

import (
	nlds "github.com/Achillesxu/XuAlgoGo/non_linear_data_structure"
)

// main project entrance
func main() {
	var tree = &nlds.BinarySearchTree{}
	tree.InsertElement(8, 8)
	tree.InsertElement(3, 3)
	tree.InsertElement(10, 10)
	tree.InsertElement(1, 1)
	tree.InsertElement(6, 6)
	tree.InsertElement(12, 12)
	tree.InsertElement(9, 9)

	tree.String()
	tree.RemoveNode(8)
	tree.String()

}
