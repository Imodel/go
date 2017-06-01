package main

import (
	"fmt"
)

type testInt func(int) bool

func isOdd(i int) bool {
	if i%2 == 0 {
		return true
	}

	return false
}

func isEven(i int) bool {
	if i%2 != 0 {
		return true
	}

	return false
}

func filter(slice []int, f testInt) []int {
	var result []int

	for _, value := range slice {
		if f(value) {
			result = append(result, value)
		}
	}

	return result
}

func main() {
	slice := []int{1, 2, 3, 4, 5, 6}
	fmt.Printf("slice = ", slice)

	odd := filter(slice, isOdd)
	fmt.Printf("odd=", odd)

	even := filter(slice, isEven)
	fmt.Printf("even = ", even)
}
