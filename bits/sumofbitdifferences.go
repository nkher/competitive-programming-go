package main 

import (
	"fmt"
)

// GetSumBitDifferencesInAllPairs() takes in an array and returns the sum of all the 
// bit differences of all the pairs that can be formed from the array elements. A bit difference
// in a pair (x, y) is the count of different bits at same positions in binary
// representations of x and y
func GetSumBitDifferencesInAllPairs(arr []int) int {

	sum := 0
	n := len(arr)

	// Travserse all the 32 bits in a number
	for i:=0; i<32; i++ {

		count := 0
		for j:=0; j<n; j++ {

			if (arr[j] & (1 << uint(i)) == 0) { // check if the ith bit is set or not
				count++; 
			}
		}

		// Add the count of differences at the ith bit to the sum
		sum += count * (n - count) * 2
	}
	return sum
}

func main() {

	arr := []int{1, 3, 5}

	fmt.Println("Sum of bit differences in all pairs is : ", GetSumBitDifferencesInAllPairs(arr))
}