package main

import (
	"fmt"
	"strconv"
)

const ROW = 3
const COL = 5

// These arrays are used to get row and column numbers of 8 neighboursof a given cell
var rowNum = []int{-1, -1, -1, 0, 0, 1, 1, 1}
var colNum = []int{-1, 0, 1, -1, 1, -1, 0, 1}

// check whether given cell (row, col) is a valid cell or not.
func isValid(row int, col int, prevRow int, prevCol int) bool {

	return (row >= 0) && (row  < ROW) && (col >= 0) && (col < COL) && !(row == prevRow && col == prevCol)
}

// A utility function to do DFS for a 2D boolean matrix. It only considers the 8 neighbours as adjacent vertices
func dfs(mat [][]rune, row int, col int, prevRow int, prevCol int, word string, path string, ind int, n int) {

	// return if current character doesn't match with the next character in the word
	if ind > n || mat[row][col] != rune(word[ind]) {
		return 
	}

	//append current character position to path
	path += string(word[ind]) + "(" + strconv.Itoa(row) + ", " + strconv.Itoa(col)  + ") "

	// current character matches with the last character in the word
	if ind == n {
		fmt.Println(path)
		return
	}

	// Recur for all connected neighbours
	for k:=0; k<8; k++ {

		if isValid(row + rowNum[k], col + colNum[k], prevRow, prevCol) {

			dfs(mat, row+rowNum[k], col+colNum[k], row, col, word, path, ind+1, n)
		}
	}
}

// The main function to find all occurrences of the word in a matrix
func FindWords(mat [][]rune, word string) {

	// traverse through the all cells of given matrix
	for i:=0; i<ROW; i++ {

		for j:=0; j<COL; j++ {

			// occurrence of first character in matrix
			if mat[i][j] == rune(word[0]) {

				dfs(mat, i, j, -1, -1, word, "", 0, len(word)-1)
			}
		}
	}
}

func main() {

	mat := [][]rune{{'B', 'N', 'E', 'Y', 'S'}, 
					{'H', 'E', 'D', 'E', 'S'}, 
					{'S', 'G', 'N', 'D', 'E'}}

	word := "DES"

	FindWords(mat, word)

}





