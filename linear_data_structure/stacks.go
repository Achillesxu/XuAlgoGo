// Package linear_data_structure
// Time    : 2021/5/11 7:47 下午
// Author  : xushiyin
// contact : yuqingxushiyin@gmail.com
package linear_data_structure

import "strconv"

// Element class
type Element struct {
	ElementValue int
}

// String method on Element class
func (element *Element) String() string {
	return strconv.Itoa(element.ElementValue)
}

type Stack struct {
	Elements   []*Element
	ElementCnt int
}

// New return a new stack
func (stack *Stack) New() {
	stack.Elements = make([]*Element, 0)
}

// Push adds a node to the stack
func (stack *Stack) Push(e *Element) {
	stack.Elements = append(stack.Elements, e)
	stack.ElementCnt += 1
}

// Pop removes and returns a node from the stack in last to first order
func (stack *Stack) Pop() *Element {
	if stack.ElementCnt == 0 {
		return nil
	}
	var stackLen = stack.ElementCnt
	var element = stack.Elements[stackLen-1]
	if stackLen > 1 {
		stack.Elements = stack.Elements[:stackLen-1]
	} else {
		stack.Elements = stack.Elements[:0]
	}
	stack.ElementCnt = len(stack.Elements)
	return element
}
