package main 

import (
	"fmt"
	"../datastructures"
)


func checkRecursive(root *bst.BstNode, n int) bool {
	if root == nil {
		return false
	}
	if bst.Size(root) == n-bst.Size(root) {
		return true
	}
	return checkRecursive(root.Left, n) || checkRecursive(root.Right, n)
}

func checkBinaryTreeEqualSizedDivision(root *bst.BstNode) bool {
	size := bst.Size(root)
	return checkRecursive(root, size)
}

func main() {
	
	/** Using BST to create a binary tree */
	var root *bst.BstNode
	root = bst.Insert(root, 19)
  	root = bst.Insert(root, 7)
  	root = bst.Insert(root, 3)
  	root = bst.Insert(root, 23)
  	root = bst.Insert(root, 20)
  	root = bst.Insert(root, 30)
  	
  	result := checkBinaryTreeEqualSizedDivision(root)
  	
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
  	
  	result = checkBinaryTreeEqualSizedDivision(root1)
  	
  	if result == true {
  		fmt.Println("Binary Tree can be divided into 2 equal halves.")
  	} else {
  		fmt.Println("Binary Tree cannot be divided into 2 equal halves.")
  	}
    
}

