package main

import (
	"fmt"
)

type CartesianNode struct {
	Data int	
	Left *CartesianNode
	Right *CartesianNode
}

func buildCartesianTreeUtil(root int, arr []int, parent []int, leftchild []int, rightchild []int) *CartesianNode {

	if root == -1 {
		return nil
	}

	// Create a new node with root's data
	temp := new(CartesianNode)
	temp.Data = arr[root]

	// Recursively construct left and right subtrees
	temp.Left = buildCartesianTreeUtil(leftchild[root], arr, parent, leftchild, rightchild)

	temp.Right = buildCartesianTreeUtil(rightchild[root], arr, parent, leftchild, rightchild)

	return temp
}

// This function builds a cartesian tree from an input array in O(n) time
func BuildCartesianTree(arr []int) *CartesianNode {

	// declare three arrays and fill them with -1
	n := len(arr)
	parent, leftchild, rightchild := make([]int, n), make([]int, n), make([]int, n)
	for i := 0; i < n; i++ {
		parent[i] = -1
		leftchild[i] = -1
		rightchild[i] = -1
	}

	// 'root' and 'last' stores the index of the root and the last processed of the Cartesian Tree.
    // Initially we take root of the Cartesian Tree as the first element of the input array. This can change
    // according to the algorithm
    root, last := 0, 0

    for i:=1; i<n; i++ {

    	last, rightchild[i] = i-1, -1

    	// Scan upward from the node's parent up to the root of the tree until a node is found
        // whose value is greater than the current one. This is the same as Step 2 mentioned in the algorithm
        for arr[last] <= arr[i] && last != root {
        	last = parent[last]
        }

        // arr[i] is the largest element yet; make it new root
        if arr[last] <= arr[i] {

        	parent[root] = i
        	leftchild[i] = root
        	root = i

        } else if rightchild[last] == -1 { // insert

        	rightchild[last] = i
        	parent[i] = last
        	leftchild[i] = -1

        } else { // reconfigure links

        	parent[rightchild[last]] = i
        	leftchild[i] = rightchild[last]
        	rightchild[last] = i
        	parent[i] = last
        }
     }

     // Since the root of the Cartesian Tree has no parent, so we assign it -1
     parent[root] = -1
     return buildCartesianTreeUtil(root, arr, parent, leftchild, rightchild)
}

func Inorder(root *CartesianNode) {
	if root != nil {
		Inorder(root.Left)
		fmt.Print(root.Data, " ")
		Inorder(root.Right)
	}
}

func main() {

	arr := []int{5, 10, 40, 30, 28}

	root := BuildCartesianTree(arr)

	Inorder(root)
	
	fmt.Println()
}