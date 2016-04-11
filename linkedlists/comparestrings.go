package main 

import (
	"fmt"
	"competitive-programming-go/Datastructures/list"
)

// CompareStrings function compares 2 strings represented as linked lists. It is similar to the 
// strcmp function in C. The function returns an integer as the return value. 0 signifies that both the
// strings are equal, -1 means that the 1st string is larger and -1 means that the second string is larger
func CompareStrings(sll1 *list.SinglyLinkedList, sll2 *list.SinglyLinkedList) int {

	// First traverse both the lists until the elements of both are equal 
	h1, h2 := sll1.Head, sll2.Head

	for (h1 != nil && h2 != nil && h1.Data == h2.Data) {
		h1 = h1.Next
		h2 = h2.Next
	}

	// Comparethe subsequent characters
	if h1 != nil && h2  != nil {
		var ret int
		if ret = -1; h1.Data.(int32) > h2.Data.(int32) {
			ret = 1
		} 
		return ret
	}

	if h1 != nil && h2 == nil {
		return 1
	}
	if h1 == nil && h2 != nil {
		return -1
	}

	return 0
}

func main() {

	sll1 := new(list.SinglyLinkedList)
	sll1.Add('n')
	sll1.Add('a')
	sll1.Add('m')
	sll1.Add('e')
	sll1.Add('s')
	sll1.Add('j')

	sll2 := new(list.SinglyLinkedList)
	sll2.Add('n')
	sll2.Add('a')
	sll2.Add('m')
	sll2.Add('e')
	sll2.Add('s')
	sll2.Add('a')
	sll2.Add('b')

	fmt.Println(CompareStrings(sll1, sll2))
}