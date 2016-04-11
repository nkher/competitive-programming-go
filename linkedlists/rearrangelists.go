package main

import (
	"fmt"
	"competitive-programming-go/Datastructures/list"
)

// ReverseLinkedList is a function to reverse a linked list data structure in place
// The method follows an iterative algorithm to reverse the linked list and 
// takes in the head of a linked list and returns the head of the reversed linked list 
func ReverseLinkedList(head *list.SinglyNode) *list.SinglyNode {

	var prev, next *list.SinglyNode
	curr := head
	for curr != nil {
		next = curr.Next
		curr.Next = prev
		prev = curr
		curr = next
	}
	head = prev
	return head
}

// RearrangeLinkedList rearranges a given linked list to a specific format (Read Further).
// Input format -  L0 -> L1 -> … -> Ln-1 -> Ln
// Output format - L0 -> Ln -> L1 -> Ln-1 -> L2 -> Ln-2 and so on
// For example if input is 1 -> 2 -> 3 -> 4 -> 5, then Output: 1 -> 5 -> 2 -> 4 -> 3 
func RearrangeLinkedList(ll1 *list.SinglyLinkedList) {

	head := ll1.Head
	// Getting the middle point - Tortoise and Hare method 
	slow, fast := head, head.Next
	for fast != nil && fast.Next != nil {
		slow = slow.Next
		fast = fast.Next.Next
	}

	// Split the linked list into 2 halves 
	firstHalf, secondHalf := head, slow.Next
	slow.Next = nil

	// Reverse the second half of the list 
	secondHalf = ReverseLinkedList(secondHalf)

	// Assing a dummy node 
	dummyNode := &list.SinglyNode{0, nil}
	temp := dummyNode

	// Start merging the alternate nodes
	for firstHalf != nil && secondHalf != nil {

		// Add element from first half 
		if firstHalf != nil {
			temp.Next = firstHalf
			temp = temp.Next
			firstHalf = firstHalf.Next
		}

		// Add element from second half 
		if secondHalf != nil {
			temp.Next = secondHalf
			temp = temp.Next
			secondHalf = secondHalf.Next
		}
	}

	// Reassign the correct node to the head 
	ll1.Head = dummyNode.Next
}

func main() {

	ll1 := new(list.SinglyLinkedList)
	ll1.Add(1)
	ll1.Add(2)
	ll1.Add(3)
	ll1.Add(4)
	ll1.Add(5)
	ll1.Add(6)

	fmt.Println("Original list is : ", ll1)

	RearrangeLinkedList(ll1)

	fmt.Println("List after modification : ", ll1)
}