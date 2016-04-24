package stack

import (
	"fmt"
)

type StackNode struct {
	val interface{}
	below *StackNode 
}

type Stack struct {
	size int
	top *StackNode
}

func (stack *Stack) Push(val interface{}) {
	stack.top = &StackNode{val, stack.top}
	stack.size++
} 

func (stack *Stack) Pop() interface{} {
	if stack.top != nil && stack.size > 0 {
		val := stack.top.val
		stack.top = stack.top.below
		stack.size--
		return val
	}	
	return nil
}

func (stack *Stack) Peek() interface{} {
	if stack.top != nil && stack.size > 0 {
		return stack.top.val
	}	
	return nil
}

func (stack *Stack) Length() int {
	return stack.size
}

func (stack *Stack) Empty() bool {
	return (stack.top == nil && stack.size == 0)
}

func (stack *Stack) Top() *StackNode {
	return stack.top
}

func (stack *Stack) String() string {
	temp := stack.top
	str := "[ "
	for temp != nil {
		str += fmt.Sprintf("%v", temp.val) + " "
		temp = temp.below
	}
	str += "]"
	return str
}
