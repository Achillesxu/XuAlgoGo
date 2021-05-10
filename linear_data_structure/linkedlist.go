// Package linear_data_structure
// Time    : 2021/5/9 10:23 下午
// Author  : xushiyin
// contact : yuqingxushiyin@gmail.com
// LinkedList is a sequence of nodes that have properties and a reference to the next node in the sequence.
// It is a linear data structure that is used to store data.
// The data structure permits the addition and deletion of components from any node next to another node.
// They are not stored contiguously in memory, which makes them different arrays.
package linear_data_structure

import "fmt"

// Node class
type Node struct {
	Property int
	NextNode *Node
}

// LinkedList class
type LinkedList struct {
	HeadNode *Node
}

// AddToHead method of LinkedList class
func (linkedList *LinkedList) AddToHead(p int) {
	var node = Node{}
	node.Property = p
	if linkedList.HeadNode != nil {
		node.NextNode = linkedList.HeadNode
	}
	linkedList.HeadNode = &node
}

// DelHead method of LinkedList class
func (linkedList *LinkedList) DelHead() {
	var node *Node
	node = linkedList.HeadNode
	if node != nil {
		linkedList.HeadNode = node.NextNode
	}
}

// IterateList method iterates over LinkedList
func (linkedList *LinkedList) IterateList() {
	var node *Node
	for node = linkedList.HeadNode; node != nil; node = node.NextNode {
		fmt.Println(node.Property)
	}
}

// LastNode method returns the last Node
func (linkedList *LinkedList) LastNode() *Node {
	var node *Node
	var lastNode *Node
	for node = linkedList.HeadNode; node != nil; node = node.NextNode {
		if node.NextNode == nil {
			lastNode = node
		}
	}
	return lastNode
}

// AddToEnd method adds the node with Property to the end
func (linkedList *LinkedList) AddToEnd(p int) {
	var node = &Node{}
	node.Property = p
	node.NextNode = nil
	var lastNode *Node
	lastNode = linkedList.LastNode()
	if lastNode != nil {
		lastNode.NextNode = node
	} else {
		linkedList.HeadNode = node
	}
}

// DelEnd method delete the last node
func (linkedList *LinkedList) DelEnd() {
	var node *Node

	for node = linkedList.HeadNode; node != nil && node.NextNode != nil; node = node.NextNode {
		if node.NextNode.NextNode == nil {
			break
		}
	}
	if node != nil {
		if node.NextNode != nil {
			node.NextNode = nil
		} else {
			linkedList.HeadNode = nil
		}
	}
}

// NodeWithValue method returns Node given parameter Property
func (linkedList *LinkedList) NodeWithValue(p int) *Node {
	var node *Node
	var nodeWith *Node
	for node = linkedList.HeadNode; node != nil; node = node.NextNode {
		if node.Property == p {
			nodeWith = node
			break
		}
	}
	return nodeWith
}

// AddAfter method adds a node with nodeProperty after node with Property
func (linkedList *LinkedList) AddAfter(nodeProperty int, property int) {
	var node = &Node{}
	node.Property = property
	node.NextNode = nil
	var nodeWith *Node
	nodeWith = linkedList.NodeWithValue(nodeProperty)
	if nodeWith != nil {
		node.NextNode = nodeWith.NextNode
		nodeWith.NextNode = node
	}
}

// AddBefore method adds a node with nodeProperty after node with Property
func (linkedList *LinkedList) AddBefore(nodeProperty int, property int) {

}

// DelNodeWithValue method delete node with vale
func (linkedList *LinkedList) DelNodeWithValue(p int) {
	var node *Node
	var targetNode *Node

	node = linkedList.HeadNode

	if node != nil {
		if node.Property == p {
			linkedList.HeadNode = nil
		} else {
			for ; node.NextNode != nil; node = node.NextNode {
				if node.NextNode.Property == p {
					targetNode = node.NextNode
					break
				}
			}
			if targetNode != nil {
				node.NextNode = targetNode.NextNode
			}
		}
	}
}

// LinkedListLen method delete node with vale
func (linkedList *LinkedList) LinkedListLen() int {
	var node *Node
	cnt := 0
	for node = linkedList.HeadNode; node != nil; node = node.NextNode {
		cnt += 1
	}
	return cnt
}
