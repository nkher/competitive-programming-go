package main 

import (
	"fmt"
	"math"
)

// The method IsYPowerOfXSimple() runs a simple algorithm to check if 
// y is a power of x. The approach is brute force approach and runs in 
// O(Log of y to base x)
func IsYPowerOfXSimple(x, y int) bool {

	if x == 1 && y == 1 {
		return true
	}

	pow := 1
	for pow < y {
		pow = pow * x
	}

	return pow == y
} 

// The method IsYPowerOfXSimple() runs an efficient algorithm to check if 
// y is a power of x. The idea is to kepp suqaring the power instead of multiplying
// it with x. if y < x^2, x^4, x^8 and so on ... then continue. if at any point they become 
// equal then return true. if x becomes more than y, then do a binary search between the previous 
// power and the current power. Keep a track of the power using a separate variable.
// Time complexity is O(Log Log y)
func IsYPowerOfXEfficient(x, y int) bool {

	// Initialize the power and a variable to keep a track of the power
	pow, i := x, 1

	for pow < y {
		pow = pow * pow
		i *= 2
	}

	// if pow == y return true
	if pow == y {
		return true
	} else { // do a binary search here 

		low, high := i/2, i
		var mid int

		for low < high {

			mid = (high + low) / 2
			val := int(math.Pow(float64(x), float64(mid)))

			if y == val {
				return true
			} else if val > y {
				high = high - 1
			} else {
				low = low + 1
			}
		}
	}

	return false
}

func main() {

	fmt.Println("Check Simple - for x = 2 and y = 128 : ", IsYPowerOfXSimple(2, 128))
	fmt.Println("Check Simple - for x = 2 and y = 30 : ", IsYPowerOfXSimple(2, 30))
	fmt.Println("Check Simple - for x = 10 and y = 1 : ", IsYPowerOfXSimple(10, 1))

	fmt.Println("Check Efficient - for x = 2 and y = 128 : ", IsYPowerOfXEfficient(2, 128))
	fmt.Println("Check Efficient - for x = 2 and y = 30 : ", IsYPowerOfXEfficient(2, 30))
	fmt.Println("Check Efficient - for x = 10 and y = 1 : ", IsYPowerOfXEfficient(10, 1))
}