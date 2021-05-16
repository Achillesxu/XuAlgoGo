// Package linear_data_structure
// Time    : 2021/5/10 8:21 下午
// Author  : xushiyin
// contact : yuqingxushiyin@gmail.com
// please check go standard package container/list, this is very important!!!
package linear_data_structure

// DNode double linked list node
type DNode struct {
	Property     int
	NextNode     *DNode
	PreviousNode *DNode
}

// DoubleLinkedList class
type DoubleLinkedList struct {
	HeadNode *DNode
}

// NodeBetweenValues method of DoubleLinkedList
func (dLinkedList *DoubleLinkedList) NodeBetweenValues(firstProperty, secondProperty int) *DNode {
	var node *DNode
	var nodeWith *DNode
	for node = dLinkedList.HeadNode; node != nil; node = node.NextNode {
		if node.PreviousNode != nil && node.NextNode != nil {
			if node.PreviousNode.Property == firstProperty && node.NextNode.Property == secondProperty {
				nodeWith = node
				break
			}
		}
	}
	return nodeWith
}

// AddToHead method of DoubleLinkedList
func (dLinkedList *DoubleLinkedList) AddToHead(property int) {
	var node = &DNode{}
	node.Property = property
	node.NextNode = nil
	if dLinkedList.HeadNode != nil {
		node.NextNode = dLinkedList.HeadNode
		dLinkedList.HeadNode.PreviousNode = node
	}
	dLinkedList.HeadNode = node
}

// AddAfter method of DoubleLinkedList
func (dLinkedList *DoubleLinkedList) AddAfter(nodeProperty, property int) {
	var node = &DNode{}
	node.Property = property
	node.NextNode = nil
	var nodeWith *DNode
	nodeWith = dLinkedList.NodeWithValue(nodeProperty)
	if nodeWith != nil {
		node.NextNode = nodeWith.NextNode
		node.PreviousNode = nodeWith
		nodeWith.NextNode = node
	}
}

// NodeWithValue method returns DNode given parameter Property
func (dLinkedList *DoubleLinkedList) NodeWithValue(p int) *DNode {
	var node *DNode
	var nodeWith *DNode
	for node = dLinkedList.HeadNode; node != nil; node = node.NextNode {
		if node.Property == p {
			nodeWith = node
			break
		}
	}
	return nodeWith
}
