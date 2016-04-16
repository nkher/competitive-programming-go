package main

import (
	"fmt"
	"sort"
)

// This method returns an array in its reduced form. Reduced form means a new array
// with 0 is replaced with its smallest element, 1 with the second smallest and so on.
// Hence for input -> {20, 10, 30} output -> {1, 0, 2}
// Time Complexity is O(N), where N = size of input array
func GetArrayInReducedForm(array []int) []int {

	// Create a copy of the array
	copyArray := make([]int, len(array))
	copy(copyArray, array)

	// Sort the copyArray
	sort.Ints(copyArray)
	
	// Make a map from int to int (value to index)
	indexMap := make(map[int]int)
	val := 0
	for i := range array {
		indexMap[copyArray[i]] = val
		val = val + 1
	}

	// Fill in the values within the actual array
	for i:=0; i<len(array); i++ {
		array[i] = indexMap[array[i]]
	}

	return array
}

func main() {

	array := []int{5, 10, 40, 30, 20}

	fmt.Println(GetArrayInReducedForm(array))

}