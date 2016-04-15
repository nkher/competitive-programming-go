/** 
Reference : http://www.geeksforgeeks.org/select-a-random-node-from-a-singly-linked-list/
*/

package main

import (
	"fmt"
	"math/rand"
	"competitive-programming-go/Datastructures/list"
)

// GetRandomNode is a function to get a random node from the linked list. Each node
// must have an equally likely probability of getting selected. The method uses
// the concept of Reservoir Sampling
func GetRandomNode(ll *list.SinglyLinkedList) *list.SinglyNode {

	// Let the result be the first node's Data
	result := ll.Head
	current := ll.Head
	n := 2

	// Iterate over all the nodes starting from the second and 
	// give each node an equally likely probability of getting selected
	// Start from n=2 and keep moving till the last node
	// At each step 
	//    1. generate a random number in between 0 and (n-1), say j
	//    	 (First step 1 it would be between 0 and 1, for 2nd it would be 0 and 2 and so on)
	//    2. Select the current node as the result only if (j == 0)
	// See more explanation below

	for n=2; current != nil; n++ {
		// change the result with probability of 1/n
		if rand.Intn(n) == 0 {
			result = current
		}
		current = current.Next
	}

	return result
}

/**

Further Explanation:

How does this work?
Let there be total N nodes in list. It is easier to understand from last node.

The probability that last node is result simply 1/N [For last or N’th node, we generate a random number between 0 to N-1 and make last node as result if the generated number is 0 (or any other fixed number]

The probability that second last node is result should also be 1/N.

The probability that the second last node is result 
          = [Probability that the second last node replaces result] X 
            [Probability that the last node doesn't replace the result] 
          = [1 / (N-1)] * [(N-1)/N]
          = 1/N
*/

func main() {

	ll1 := new(list.SinglyLinkedList)
	ll1.Add(1)
	ll1.Add(2)
	ll1.Add(3)
	ll1.Add(4)
	ll1.Add(5)
	ll1.Add(6)

	fmt.Println("Random selected node is : ", GetRandomNode(ll1).Data)
}