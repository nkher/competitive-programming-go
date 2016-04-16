package main 

import (
	"fmt"
	"math"
)

// GetAllPerfectSquaresBetweenTwoNumbers() is a method that returns all perfect 
// squares between two numbers. 
// Algorithm - floor(sqrt(b)) - ceil(sqrt(a)) + 1
func GetAllPerfectSquaresBetweenTwoNumbers(a, b int) int {

	if a > b {
		return GetAllPerfectSquaresBetweenTwoNumbers(b, a)
	}

	c := float64(a)
	d := float64(b)
	return int(math.Floor(math.Sqrt(d)) - math.Ceil(math.Sqrt(c))) + 1
}

func main() {

	a, b := 9, 25

	fmt.Println("Number of perfect squares in between a and b are : ", 
				GetAllPerfectSquaresBetweenTwoNumbers(a, b))

	fmt.Println("Number of perfect squares in between a and b are : ", 
				GetAllPerfectSquaresBetweenTwoNumbers(b, a))

	a, b = 10, 100

	fmt.Println("Number of perfect squares in between a and b are : ", 
				GetAllPerfectSquaresBetweenTwoNumbers(b, a))
}