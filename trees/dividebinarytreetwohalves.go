// +build XXX

package main 

import (
	"fmt"
	"competitive-programming-go/datastructures/bst"
)

/** Solution 1 - Brute Force - O(N^2) */
func CheckRecursive(node *bst.BstNode, n int) bool {
	
	if node == nil {
		return false
	}
	if bst.Size(node) == n-bst.Size(node) {
		return true
	}
	return CheckRecursive(node.Left, n) || CheckRecursive(node.Right, n)
}

func CheckBinaryTreeEqualSizedDivision(root *bst.BstNode) bool {
	size := bst.Size(root)
	return CheckRecursive(root, size)
}

/** Solution 2 - O(N) */

func CheckRecursiveEfficient(node *bst.BstNode, n int, result *bool) int {
	
	if node == nil {
		return 0
	}
	
	// Calculate the sizes of the left and right child
	c := CheckRecursiveEfficient(node.Left, n, result) + 1 + 
			CheckRecursiveEfficient(node.Right, n, result)
	
	// Check if tree can be divided or not
	if c == n-c {
		*result = true 
	}
	
	return c
}

func CheckBinaryTreeEqualSizedDivisionEfficient(root *bst.BstNode) bool {
	size := bst.Size(root)
	result := false
	CheckRecursiveEfficient(root, size, &result)
	return result
}

func main() {
	
	  // Using BST to create a binary tree
	  var root *bst.BstNode
	  root = bst.Insert(root, 19)
  	root = bst.Insert(root, 7)
  	root = bst.Insert(root, 3)
  	root = bst.Insert(root, 23)
  	root = bst.Insert(root, 20)
  	root = bst.Insert(root, 30)
  	
  	fmt.Println("\n--------Using BruteForce Algorithm------------\n")
  	
  	result := CheckBinaryTreeEqualSizedDivision(root)
  	
  	if result == true {
  		fmt.Println("Binary Tree can be divided into 2 equal halves.")
  	} else {
  		fmt.Println("Binary Tree cannot be divided into 2 equal halves.")
  	}
  	
  	var root1 *bst.BstNode
	  root1 = bst.Insert(root1, 10)
  	root1 = bst.Insert(root1, 1)
  	root1 = bst.Insert(root1, 30)
  	root1 = bst.Insert(root1, 25)
  	root1 = bst.Insert(root1, 20)
  	root1 = bst.Insert(root1, 28)
  	root1 = bst.Insert(root1, 40)
  	root1 = bst.Insert(root1, 50)
  	
  	result = CheckBinaryTreeEqualSizedDivision(root1)
  	if result == true {
  		fmt.Println("Binary Tree can be divided into 2 equal halves.")
  	} else {
  		fmt.Println("Binary Tree cannot be divided into 2 equal halves.")
  	}
  	
  	fmt.Println("\n--------Using Efficient Algorithm------------\n")
  	
  	result = CheckBinaryTreeEqualSizedDivisionEfficient(root)
    if result == true {
  		fmt.Println("Binary Tree can be divided into 2 equal halves.")
  	} else {
  		fmt.Println("Binary Tree cannot be divided into 2 equal halves.")
  	}
  	
  	result = CheckBinaryTreeEqualSizedDivisionEfficient(root1)
    if result == true {
  		fmt.Println("Binary Tree can be divided into 2 equal halves.")
  	} else {
  		fmt.Println("Binary Tree cannot be divided into 2 equal halves.\n")
  	}
  	
  	
}

