package main 

import (
	"fmt"
)

// GetNumberOfDyckPaths() is method that returns the number of Dyck Paths in an (n * n)
// grid. A Dyck Path is a path from the bottom left corner of the grid to the top right
// corner of the grid with the condition that the path lies above the diagonal cells. 
func GetNumberOfDyckPaths(n int) int {

	result := 1
	for i:=0; i<n; i++ {
		result *= (2*n - i)
		result /= (i + 1)
	}

	return result/(n+1) // This is 2nCn/(n+1)
}

func main() {

	// Catalan number starts from 0
	fmt.Println("Number of dyck paths in a 5 by 5 grid are is : ", GetNumberOfDyckPaths(5))
}