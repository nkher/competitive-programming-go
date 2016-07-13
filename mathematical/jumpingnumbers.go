package main

import (
	"fmt"
	"competitive-programming-go/datastructures/queue"
)

// This method prints all the numbers that are jumping numbers
// which are lower than or equal to the given number. A number 
// is called as a Jumping Number if all adjacent digits in 
// it differ by 1. The difference between ‘9’ and ‘0’ is not 
// considered as 1. All single digit numbers are considered 
// as Jumping Numbers. For example 7, 8987 and 4343456 are 
// Jumping numbers but 796 and 89098 are not.
// Time Complexity - O(n), where n are the total number of jumping numbers
func PrintJumpingNumbersLowerAndEqual(x int) {

	fmt.Println(0)
	for i:=1; i<=9 && i <=x; i++ {
		bfs(x, i)
	}
}

func bfs(x, num int) {

	// Create a queue and insert 'num' into it
	queue := new(queue.Queue)
	queue.Insert(num)

	// Do BFS starting from i
	for !queue.Empty() {

		if head, ok := queue.Remove().(int); ok {

			if head <= x {

				fmt.Println(head) // print the head
				lastDigit := head % 10

				// If last digit is 0, append next digit only
				if lastDigit == 0 {

					queue.Insert((head * 10) + (lastDigit + 1))

				} else if lastDigit == 9 {  // If last digit is 9, append previous digit only

					queue.Insert((head * 10) + (lastDigit - 1))

				} else { // If last digit is not 9 nor 0, append next and previous both

					queue.Insert((num * 10) + (lastDigit - 1))
					queue.Insert((num * 10) + (lastDigit + 1))
				}

			}
		} else {
			break
		}
	}
}

func main() {

	fmt.Println("Jumping numbers when x = 50")
	PrintJumpingNumbersLowerAndEqual(50)

	fmt.Println("Jumping numbers when x = 10")
	PrintJumpingNumbersLowerAndEqual(10)

}