package main 

import (
	"fmt"
	"competitive-programming-go/datastructures/bst"
)

func getSumOfUncoveredNodesLeft(node *bst.BstNode) int {
	
	if node.Left == nil && node.Right == nil {
		return node.Data
	}
	
	if node.Left != nil {
		return node.Data + getSumOfUncoveredNodesLeft(node.Left)
	}
	return node.Data + getSumOfUncoveredNodesLeft(node.Right)
}

func getSumOfUncoveredNodesRight(node *bst.BstNode) int {
	
	if node.Right == nil && node.Left == nil {
		return node.Data
	}
	
	if node.Right != nil {
		return node.Data + getSumOfUncoveredNodesRight(node.Right)
	}
	return node.Data + getSumOfUncoveredNodesRight(node.Left)
}

func getSumOfUncoveredElements(root *bst.BstNode) int {
	
	leftBoundary, rightBoundary := 0, 0
	
	if root.Left != nil {
		leftBoundary = getSumOfUncoveredNodesLeft(root.Left)
	}
	
	if  root.Right != nil {
		rightBoundary = getSumOfUncoveredNodesRight(root.Right)
	}
	
	return (root.Data + leftBoundary + rightBoundary)
}

func getSumOfTree(root *bst.BstNode) int {
	if root == nil {
		return 0
	}
	return root.Data + getSumOfTree(root.Left) + getSumOfTree(root.Right)
}

func checkIfSameSum(root *bst.BstNode) bool {
	
	uncoveredSum := getSumOfUncoveredElements(root)
	sumOfTree := getSumOfTree(root)
	
	return (sumOfTree == (sumOfTree - uncoveredSum))
}

func main() {
	
	/** Using BST to create a binary tree */
	var root *bst.BstNode
	root = bst.Insert(root, 8)
  	root = bst.Insert(root, 10)
  	root = bst.Insert(root, 14)
  	root = bst.Insert(root, 13)
  	root = bst.Insert(root, 3)
  	root = bst.Insert(root, 1)
  	root = bst.Insert(root, 6)
  	root = bst.Insert(root, 4)
  	root = bst.Insert(root, 7)
  	
  	result := checkIfSameSum(root)
  	
  	if result == true {
  		fmt.Println("Sum of covered and uncovered is same\n")
  	} else {
  		fmt.Println("Sum of covered and uncovered is not the same\n")
  	}
}

