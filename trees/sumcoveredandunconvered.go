package main 

import (
	"fmt"
	"competitive-programming-go/datastructures/bst"
)

func GetSumOfUncoveredNodesLeft(node *bst.BstNode) int {
	
	if node.Left == nil && node.Right == nil {
		return node.Data
	}
	
	if node.Left != nil {
		return node.Data + GetSumOfUncoveredNodesLeft(node.Left)
	}
	return node.Data + GetSumOfUncoveredNodesLeft(node.Right)
}

func GetSumOfUncoveredNodesRight(node *bst.BstNode) int {
	
	if node.Right == nil && node.Left == nil {
		return node.Data
	}
	
	if node.Right != nil {
		return node.Data + GetSumOfUncoveredNodesRight(node.Right)
	}
	return node.Data + GetSumOfUncoveredNodesRight(node.Left)
}

func GetSumOfUncoveredElements(root *bst.BstNode) int {
	
	leftBoundary, rightBoundary := 0, 0
	
	if root.Left != nil {
		leftBoundary = GetSumOfUncoveredNodesLeft(root.Left)
	}
	
	if  root.Right != nil {
		rightBoundary = GetSumOfUncoveredNodesRight(root.Right)
	}
	
	return (root.Data + leftBoundary + rightBoundary)
}

func GetSumOfTree(root *bst.BstNode) int {
	if root == nil {
		return 0
	}
	return root.Data + GetSumOfTree(root.Left) + GetSumOfTree(root.Right)
}

func CheckIfSameSum(root *bst.BstNode) bool {
	
	uncoveredSum := GetSumOfUncoveredElements(root)
	sumOfTree := GetSumOfTree(root)
	
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
  	
  	result := CheckIfSameSum(root)
  	
  	if result == true {
  		fmt.Println("Sum of covered and uncovered is same\n")
  	} else {
  		fmt.Println("Sum of covered and uncovered is not the same\n")
  	}
}

