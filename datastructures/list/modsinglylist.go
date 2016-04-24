package list

import (
	"fmt"
)

type ModSinglyLinkedList struct {
	Head *ModSinglyNode
	size int
}

type ModSinglyNode struct {
	Data interface{}
	Next *ModSinglyNode
	Arbitary *ModSinglyNode
}

func (sll *ModSinglyLinkedList) Length() int {
	return sll.size
}

func (sll *ModSinglyLinkedList) Add(val interface{}) {
	newNode := &ModSinglyNode{val, nil, nil}
	if sll.Head == nil {
		sll.Head = newNode	
	} else {
		temp := sll.Head
		for temp.Next != nil {
			temp = temp.Next
		}
		temp.Next = newNode
	}
	sll.size++
}

func (sll *ModSinglyLinkedList) Empty() bool {
	return (sll.Head == nil && sll.size == 0)
}

func (sll *ModSinglyLinkedList) Remove(val interface{}) bool {

	curr, prev := sll.Head, sll.Head
	if sll.Empty() {
		fmt.Printf("Cannot remove from an empty list")
	} else {
		for (curr != nil) {
			if curr.Data == val {
				if curr == sll.Head { // this is the first element
					sll.Head = curr.Next
				} else {
					prev.Next = curr.Next
					break
				}
			}
			prev = curr
			curr = curr.Next
		}
		return true
	}
	return false
}

func (sll *ModSinglyLinkedList) Contains(val interface{}) bool {

	curr := sll.Head
	for curr != nil {
		if curr.Data == val {
			return true
		}
		curr = curr.Next
	}
	return false
}

// This prints the element of the linked list nodes in the following format.
// Example Node -> 5-10, 2-10, 10-nil
// The first node conatins 5 as its own value with the arbitary node pointing to 10
// Similarly the third node contains a value of 10 and its arbitary node pointing to nil 
func (sll *ModSinglyLinkedList) String() string {
	temp := sll.Head
	str := "[ "
	for temp != nil {
		str += fmt.Sprintf("%v", temp.Data)
		if temp.Arbitary == nil {
			str += fmt.Sprintf("-nil ")
		} else {
			str += fmt.Sprintf("-%v", temp.Arbitary.Data) + " "
		}
		if temp.Next != nil {
				str += ", "
		}
		temp = temp.Next
	}
	str += "]"
	return str
}



