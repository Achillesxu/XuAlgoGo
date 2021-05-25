// Package hetergeneous_data_structure
// Time    : 2021/5/25 7:00 下午
// Author  : xushiyin
// contact : yuqingxushiyin@gmail.com
package hetergeneous_data_structure

// Node struct
type Node struct {
	NextNode *Node
	Property rune
}

// CreateLinkedList create linked list
func CreateLinkedList() *Node {
	headNode := &Node{nil, 'a'}
	curNode := headNode
	for i := 'b'; i <= 'z'; i++ {
		node := &Node{nil, i}
		curNode.NextNode = node
		curNode = node
	}
	return headNode
}

// ReverseLinkedList reverse list
func ReverseLinkedList(nl *Node) *Node {
	curNode := nl
	topNode := (*Node)(nil)
	for {
		if curNode == nil {
			break
		}
		tmpNode := (*Node)(nil)
		tmpNode = curNode.NextNode
		curNode.NextNode = topNode
		topNode = curNode
		curNode = tmpNode
	}
	return topNode
}
