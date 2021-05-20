// Package non_linear_data_structure
// Time    : 2021/5/16 2:15 下午
// Author  : xushiyin
// contact : yuqingxushiyin@gmail.com
// Adelson, Velski, and Landis pioneered the AVL tree data structure and hence it is named after them.
// It consists of height adjusting binary search trees.
// The balance factor is obtained by finding the difference between the heights of the left and right sub-trees.
// Balancing is done using rotation techniques.
// If the balance factor is greater than one,
// rotation shifts the nodes to the opposite of the left or right sub-trees.
// The search, addition, and deletion operations are processed in the order of O(log n).
package non_linear_data_structure

// KeyValue type
type KeyValue interface {
	LessThan(KeyValue) bool
	EqualTo(KeyValue) bool
}

// AVLTreeNode class
type AVLTreeNode struct {
	KeyValue     KeyValue
	BalanceValue int
	LinkedNodes  [2]*AVLTreeNode
}

// opposite method takes a node value and returns the opposite node's value
// because LinkedNodes is a array of AVLTreeNode
func opposite(nodeValue int) int {
	return 1 - nodeValue
}

// singleRotation method rotates the node opposite to the specified sub-tree
func singleRotation(rootNode *AVLTreeNode, nodeValue int) *AVLTreeNode {
	var saveNode *AVLTreeNode
	saveNode = rootNode.LinkedNodes[opposite(nodeValue)]
	rootNode.LinkedNodes[opposite(nodeValue)] = saveNode.LinkedNodes[nodeValue]
	saveNode.LinkedNodes[nodeValue] = rootNode
	return saveNode
}

// singleRotation method rotates the node twice
func doubleRotation(rootNode *AVLTreeNode, nodeValue int) *AVLTreeNode {
	var saveNode *AVLTreeNode
	saveNode = rootNode.LinkedNodes[opposite(nodeValue)].LinkedNodes[nodeValue]
	rootNode.LinkedNodes[opposite(nodeValue)].LinkedNodes[nodeValue] = saveNode.LinkedNodes[opposite(nodeValue)]
	saveNode.LinkedNodes[opposite(nodeValue)] = rootNode.LinkedNodes[opposite(nodeValue)]
	rootNode.LinkedNodes[opposite(nodeValue)] = saveNode
	saveNode = rootNode.LinkedNodes[opposite(nodeValue)]
	rootNode.LinkedNodes[opposite(nodeValue)] = saveNode.LinkedNodes[nodeValue]
	saveNode.LinkedNodes[nodeValue] = rootNode
	return saveNode
}

// adjustBalance method adjusts the balance of the tree
func adjustBalance(rootNode *AVLTreeNode, nodeValue int, balanceValue int) {
	var node *AVLTreeNode
	node = rootNode.LinkedNodes[nodeValue]
	var oppNode *AVLTreeNode
	oppNode = node.LinkedNodes[opposite(balanceValue)]
	switch oppNode.BalanceValue {
	case 0:
		rootNode.BalanceValue = 0
		node.BalanceValue = 0
	case balanceValue:
		rootNode.BalanceValue = -balanceValue
		node.BalanceValue = 0
	default:
		rootNode.BalanceValue = 0
		node.BalanceValue = balanceValue
	}
	oppNode.BalanceValue = 0
}

// BalanceTree method changes the balance factor by a single or double rotation
func BalanceTree(rootNode *AVLTreeNode, nodeValue int) *AVLTreeNode {
	var node *AVLTreeNode
	node = rootNode.LinkedNodes[nodeValue]
	var balance int
	balance = 2*nodeValue - 1
	if node.BalanceValue == balance {
		rootNode.BalanceValue = 0
		node.BalanceValue = 0
		return singleRotation(rootNode, opposite(nodeValue))
	}
	adjustBalance(rootNode, nodeValue, balance)
	return doubleRotation(rootNode, opposite(nodeValue))
}

// insertRNode method inserts the node and balances the tree
func insertRNode(rootNode *AVLTreeNode, key KeyValue) (*AVLTreeNode, bool) {
	if rootNode == nil {
		return &AVLTreeNode{KeyValue: key}, false
	}
	var dir int
	dir = 0
	if rootNode.KeyValue.LessThan(key) {
		dir = 1
	}
	var done bool
	rootNode.LinkedNodes[dir], done = insertRNode(rootNode.LinkedNodes[dir], key)
	if done {
		return rootNode, true
	}
	rootNode.BalanceValue = rootNode.BalanceValue + (2*dir - 1)
	switch rootNode.BalanceValue {
	case 0:
		return rootNode, true
	case 1, -1:
		return rootNode, false
	}
	return BalanceTree(rootNode, dir), true
}

// InsertNode method
func InsertNode(treeNode **AVLTreeNode, key KeyValue) {
	*treeNode, _ = insertRNode(*treeNode, key)
}

// RemoveNode method
func RemoveNode(treeNode **AVLTreeNode, key KeyValue) {
	*treeNode, _ = removeRNode(*treeNode, key)
}

// removeBalance method
func removeBalance(rootNode *AVLTreeNode, nodeValue int) (*AVLTreeNode, bool) {
	var node *AVLTreeNode
	node = rootNode.LinkedNodes[opposite(nodeValue)]
	var balance int
	balance = 2*nodeValue - 1
	switch node.BalanceValue {
	case -balance:
		rootNode.BalanceValue = 0
		node.BalanceValue = 0
		return singleRotation(rootNode, nodeValue), false
	case balance:
		adjustBalance(rootNode, opposite(nodeValue), -balance)
		return doubleRotation(rootNode, nodeValue), false
	}
	rootNode.BalanceValue = -balance
	node.BalanceValue = balance
	return singleRotation(rootNode, nodeValue), true
}

// removeRNode method
func removeRNode(rootNode *AVLTreeNode, key KeyValue) (*AVLTreeNode, bool) {
	if rootNode == nil {
		return nil, false
	}
	if rootNode.KeyValue.EqualTo(key) {
		switch {
		case rootNode.LinkedNodes[0] == nil:
			return rootNode.LinkedNodes[1], false
		case rootNode.LinkedNodes[1] == nil:
			return rootNode.LinkedNodes[0], false
		}

		var heirNode *AVLTreeNode
		heirNode = rootNode.LinkedNodes[0]
		for heirNode.LinkedNodes[1] != nil {
			heirNode = heirNode.LinkedNodes[1]
		}
		rootNode.KeyValue = heirNode.KeyValue
		key = heirNode.KeyValue
	}
	var dir int
	dir = 0
	if rootNode.KeyValue.LessThan(key) {
		dir = 1
	}
	var done bool
	rootNode.LinkedNodes[dir], done = removeRNode(rootNode.LinkedNodes[dir], key)
	if done {
		return rootNode, true
	}
	rootNode.BalanceValue = rootNode.BalanceValue + (1 - 2*dir)
	switch rootNode.BalanceValue {
	case 1, -1:
		return rootNode, true
	case 0:
		return rootNode, false
	}
	return removeBalance(rootNode, dir)
}

type IntegerKey int

func (k IntegerKey) LessThan(k1 KeyValue) bool {
	return k < k1.(IntegerKey)
}
func (k IntegerKey) EqualTo(k1 KeyValue) bool {
	return k == k1.(IntegerKey)
}
