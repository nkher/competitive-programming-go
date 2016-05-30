package main

import (
	"fmt"
	"competitive-programming-go/util"
)
	
// This function returns the maximum product of a triplet of integers in the array
func MaxTripletProduct(list []int) int64 {

	n := len(list)

	if n < 3 {
		return -1
	}

	// Initialize Maximum, second maximum and third maximum element
	maxFirst, maxSecond, maxThird := mymath.MinInt64, mymath.MinInt64, mymath.MinInt64

	// Initialize Minimum and second mimimum element
	minFirst, minSecond := mymath.MaxInt64, mymath.MaxInt64

	for i:=0; i<n; i++ {

		// Update Maximum, second maximum and third maximum element
		if list[i] > maxFirst {

			maxThird = maxSecond
			maxSecond = maxFirst
			maxFirst = list[i]

		} else if (list[i] > maxSecond) { // update only second and third element

			maxThird = maxSecond
			maxSecond = list[i]

		} else if (list[i] > maxThird) { // update the third element
			maxThird = list[i]
		}

		// Update the minimum and second minimum element
		if (list[i] < minFirst) {

			minSecond = minFirst
			minFirst = list[i]

		} else if (list[i] < minSecond) { // Update the second minimum element
			minSecond = list[i]
		}

	}

	return mymath.Max( (int64)(minFirst * minSecond * maxFirst) , (int64)(maxFirst * maxSecond * maxThird) )
}


func main() {

	var list = []int{1, -4, 3, -6, 7, 0}
	fmt.Println("Max product in the array is : ", MaxTripletProduct(list))

	list = []int{10, 3, 5, 6, 20}
	fmt.Println("Max product in the array is : ", MaxTripletProduct(list))

	list = []int{-10, -3, -5, -6, -20}
	fmt.Println("Max product in the array is : ", MaxTripletProduct(list))

}
