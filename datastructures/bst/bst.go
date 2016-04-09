package bst

import "fmt"

type BstNode struct {
	Data int
	Left *BstNode
	Right *BstNode
}

func Insert(root *BstNode, data int) *BstNode {
	if root == nil {
		return &BstNode{data, nil, nil}
	}
	if data < root.Data {
		root.Left = Insert(root.Left, data)
		return root
	} else {
		root.Right = Insert(root.Right, data)
		return root
	}
}

func Height(root *BstNode) int {
	if root == nil {
		return 0
	}
	leftHeight := Height(root.Left)
	rightHeight := Height(root.Right)
	if leftHeight > rightHeight {
		return leftHeight + 1
	}
	return rightHeight + 1
}

func Size(root *BstNode) int {
	if root == nil {
		return 0
	}
	return Size(root.Left) + 1 + Size(root.Right)
}

func Inorder(root *BstNode) {
	if root != nil {
		Inorder(root.Left)
		fmt.Print(root.Data, " ")
		Inorder(root.Right)
	}
}

func Preorder(root *BstNode) {
	if root != nil {
		fmt.Print(root.Data, " ")
		Preorder(root.Left)
		Preorder(root.Right)
	}
}

func Postorder(root *BstNode) {
	if root != nil {
		Postorder(root.Left)
		Postorder(root.Right)
		fmt.Print(root.Data, " ")
	}
}