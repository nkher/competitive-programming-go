package main

import (
	"fmt"
)

// This function checks if the string representing a number can be
// partitioned into two halves such that both havles are divisible
// by either of x or y. To make it more clear if the first half
// is divisible by x then the second half must be divisible by y
// and vice versa. O(n) time algorithm.
func CheckNumberPartitionDivisibility(str string, x, y int) {

	if len(str) == 0 {
		fmt.Println("NO")
		return
	}

	length := len(str)
	lr := make([]int, length)
	rl := make([]int, length)

	lr[0] = (int(str[0]) - '0') % x
	for i:=1; i<length; i++ {
		lr[i] = (lr[i-1] * 10 + (int(str[i]) - '0')) % x
	}

	rl[length-1] = (int(str[length-1]) - '0')  % x
	powerTen, previous := 10, int(str[length-1]) - '0'

	for i:=length-2; i>=0; i-- {

		current := ((int(str[i]) - '0') * powerTen)
		rl[i] = (current+previous) % y
		previous = (current+previous)
		powerTen = (powerTen * 10)
	}

	for i:=0; i<length; i++ {

		if lr[i] != 0 {
			continue
		}

		if rl[i+1] == 0 {

			fmt.Println("YES \n")
			for k:=0; k<=i; k++ {
				fmt.Print((int)(str[k] - '0'))
			}

			fmt.Print(", ")

			for k:=i+1; k<length; k++ {
				fmt.Print((int)(str[k] - '0'))
			}

			return
		}
	}

	fmt.Println("NO \n")
}

func main() {

	str := "1234"
	x, y := 12, 34

	CheckNumberPartitionDivisibility(str, x, y)
}


