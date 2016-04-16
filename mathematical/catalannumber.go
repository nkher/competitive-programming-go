package main 

import (
	"fmt"
)

// CatalanNumber() generates the nth catalan number in the series of Catalan Number.
// A catalan number is a number that satisfies the following recurrence relation.
// --> if C0 = 1 then Cn+1 = Summation of (Ci * Cn-i) from i=0 to n, where n >= 0
func CatalanNumber(n int) int {

	result := 0
	if n <= 1 {
		return 1
	}

	for i:=0; i<n; i++ {
		result += CatalanNumber(i) * CatalanNumber(n - i -1)
	}

	return result
}	

func main() {

	// Catalan number starts from 0
	fmt.Println("4th CatalanNumber is : ", CatalanNumber(4))
}