package main

import (
	"fmt"
)

func main() {
	var min, max int

	min, max = maxMin(76, 65)

	fmt.Printf("min value is %d ,max value is %d", min, max)
}

func maxMin(a, b int) (int, int) {
	var min, max int
	if a < b {
		min = a
		max = b
	} else {
		min = b
		max = a
	}

	return min, max
}
