package list

import (
	"fmt"
)

type SinglyLinkedList struct {
	Head *SinglyNode
	size int
}

type SinglyNode struct {
	Data interface{}
	Next *SinglyNode
}

func (sll *SinglyLinkedList) Length() int {
	return sll.size
}

func (sll *SinglyLinkedList) Add(val interface{}) {
	newNode := &SinglyNode{val, nil}
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

func (sll * SinglyLinkedList) Empty() bool {
	return (sll.Head == nil && sll.size == 0)
}

func (sll *SinglyLinkedList) Remove(val interface{}) bool {

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

func (sll *SinglyLinkedList) Contains(val interface{}) bool {

	curr := sll.Head
	for curr != nil {
		if curr.Data == val {
			return true
		}
		curr = curr.Next
	}
	return false
}

func (sll *SinglyLinkedList) String() string {
	temp := sll.Head
	str := "[ "
	for temp != nil {
		str += fmt.Sprintf("%v", temp.Data) + " "
		temp = temp.Next
	}
	str += "]"
	return str
}



