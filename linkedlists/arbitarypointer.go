package main 

import (
	"fmt"
	"competitive-programming-go/Datastructures/list"
)

// This function reverses a linked list in place in O(N) time using the iterative approach.
// It takes in the linked list and returns the head of the reversed linked list
func ReverseLinkedList(head *list.ModSinglyNode) *list.ModSinglyNode {

	curr := head
	var prev, next *list.ModSinglyNode
	for curr != nil {	
		next = curr.Next
		curr.Next = prev
		prev = curr
		curr = next
	}
	return prev
}

// This function takes in a Modified Linked List that contains a Next and an arbitary pointer,
// and returns the same linked list with the arbitary pointers pointing to the node with
// maximum value node that lies on the right. The last node's arbitary pointer points to nil. 
func PopulateArbitaryNodeWithRightMaxNode(sll *list.ModSinglyLinkedList) {

	// Reverse the given linked list
	sll.Head = ReverseLinkedList(sll.Head)

	// Initialize a pointer to the max value
	max := sll.Head

	// Traverse the reversed list
	temp := sll.Head.Next
	for temp != nil {

		temp.Arbitary = max

		if max.Data.(int) < temp.Data.(int) {
			max = temp
		}

		temp = temp.Next
	}

	sll.Head = ReverseLinkedList(sll.Head)
}

func main() {

	sll := new(list.ModSinglyLinkedList)
	sll.Add(5)
	sll.Add(10)
	sll.Add(2)
	sll.Add(3)

	fmt.Println("Linked list before populating the arbitary node : ", sll)

	PopulateArbitaryNodeWithRightMaxNode(sll)

	fmt.Println("Linked list after populating the arbitary node : ", sll)
}