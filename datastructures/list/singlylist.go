package list

import (
	"fmt"
)

type SinglyLinkedList struct {
	head *SinglyNode
	size int
}

type SinglyNode struct {
	data interface{}
	next *SinglyNode
}

func (sll *SinglyLinkedList) Length() int {
	return sll.size
}

func (sll *SinglyLinkedList) Add(val interface{}) {
	newNode := &SinglyNode{val, nil}
	if sll.head == nil {
		sll.head = newNode	
	} else {
		temp := sll.head
		for temp.next != nil {
			temp = temp.next
		}
		temp.next = newNode
	}
	sll.size++
}

func (sll * SinglyLinkedList) Empty() bool {
	return (sll.head == nil && sll.size == 0)
}

func (sll *SinglyLinkedList) Remove(val interface{}) bool {

	curr, prev := sll.head, sll.head
	if sll.Empty() {
		fmt.Printf("Cannot remove from an empty list")
	} else {
		for (curr != nil) {
			if curr.data == val {
				if curr == sll.head { // this is the first element
					sll.head = curr.next
				} else {
					prev.next = curr.next
					break
				}
			}
			prev = curr
			curr = curr.next
		}
		return true
	}
	return false
}

func (sll *SinglyLinkedList) Contains(val interface{}) bool {

	curr := sll.head
	for curr != nil {
		if curr.data == val {
			return true
		}
		curr = curr.next
	}
	return false
}

func (sll *SinglyLinkedList) String() string {
	temp := sll.head
	str := "[ "
	for temp != nil {
		str += fmt.Sprintf("%v", temp.data) + " "
		temp = temp.next
	}
	str += "]"
	return str
}

