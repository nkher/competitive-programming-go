package main

import (
	"fmt"
)


// This is a method that finds the minimum number of merge operation required
// to make an array in the form of a palindrome. Essentially its elements are such that
// they would look like a palindrome. Time Complexity is O(n)
func MinMergeToGetArrayPalindrome(array []int) int {

	numberOfMerges := 0

	// Keep 2 pointers at each end and move inwards
	i, j := 0, len(array)-1
	for i<=j {

		// If the elements at both ends are same just move ahead respectively
		if array[i] == array[j] {
			i++
			j--
		} else if array[i] > array[j] { // Left greater than right, merge right and one before the right

			j--	
			array[j] = array[j] + array[j+1]
			numberOfMerges++

		} else {

			i++
			array[i] = array[i] + array[i-1]
			numberOfMerges++	
		}
	}
	return numberOfMerges 
}

func main() {


	array := []int{1, 4, 5, 9, 1}
	fmt.Println("Number of merges required : ", MinMergeToGetArrayPalindrome(array))

	array = []int{1, 4, 5, 1}
	fmt.Println("Number of merges required : ", MinMergeToGetArrayPalindrome(array))

	array = []int{15, 4, 15}
	fmt.Println("Number of merges required : ", MinMergeToGetArrayPalindrome(array))

	array = []int{11, 14, 15, 99}
	fmt.Println("Number of merges required : ", MinMergeToGetArrayPalindrome(array))


}