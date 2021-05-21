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
	// var treeNode *nlds.AVLTreeNode
	// fmt.Println("Tree is empty")
	// var avlTree []byte
	// avlTree, _ = json.MarshalIndent(treeNode, "", " ")
	// fmt.Println(string(avlTree))
	// fmt.Println("\n Add Tree")
	//
	// nlds.InsertNode(&treeNode, nlds.IntegerKey(5))
	// nlds.InsertNode(&treeNode, nlds.IntegerKey(3))
	// nlds.InsertNode(&treeNode, nlds.IntegerKey(8))
	// nlds.InsertNode(&treeNode, nlds.IntegerKey(7))
	// nlds.InsertNode(&treeNode, nlds.IntegerKey(6))
	// nlds.InsertNode(&treeNode, nlds.IntegerKey(10))
	// avlTree, _ = json.MarshalIndent(treeNode, "", " ")
	// fmt.Println(string(avlTree))
	//
	// fmt.Println("\n Delete Tree")
	// nlds.RemoveNode(&treeNode, nlds.IntegerKey(3))
	// nlds.RemoveNode(&treeNode, nlds.IntegerKey(7))
	// avlTree, _ = json.MarshalIndent(treeNode, "", " ")
	// fmt.Println(string(avlTree))
	a := [10]int{1, 2}
	fmt.Println(a)

	linkedList := lds.LinkedList{}
	linkedList.AddToHead(1)
	linkedList.AddToHead(3)
	linkedList.AddToHead(9)
	fmt.Println(linkedList.HeadNode.Property)
	linkedList.AddToEnd(27)

	linkedList.AddAfter(1, 15)
	fmt.Println("")
	linkedList.IterateList()
	linkedList.DelNodeWithValue(15)
	fmt.Println("")
	linkedList.IterateList()
	linkedList.DelNodeWithValue1(1)
	fmt.Println("")
	linkedList.IterateList()
	linkedList.DelNodeWithValueX(27)
	fmt.Println("")
	linkedList.IterateList()

}
