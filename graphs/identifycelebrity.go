package main

import (
	"fmt"
)

// This function finds a celebrity in a given matrix and returns the 
// id of the celebrity. The id here is the index of the celebrity.
// Celebrity is a person who knows no one and is known by every one.
func FindCelebrity(matrix [][]int) int {

	n, i, j := len(matrix), 0, 1

	for j < n {

		if matrix[i][j] == 1 { //  means 'j' is known and 'i' is definitely not the celebrity
			i=j+1
			i=j
		} else { // 'j' is not known and he is not the celebrity
			j=j+1
		}
	}

	return i
}

func main() {

	matrix := [][]int  {{0, 1, 0},
						{0, 0, 0},
						{0, 1, 0}}

	fmt.Println("Celebrity Id : ", FindCelebrity(matrix))

	matrixTwo := [][]int  {{0, 1, 0, 0, 1},
						{0, 0, 1, 0, 0},
						{0, 1, 0, 1, 1},
						{0, 0, 0, 0, 0},
						{0, 1, 0, 1, 1}}

	fmt.Println("Celebrity Id : ", FindCelebrity(matrixTwo))

}