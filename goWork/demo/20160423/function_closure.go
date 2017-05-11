package main

import (
	"fmt"
)

func f() func(int) int {
	var x int
	return func(delta int) int {
		x += delta
		return x
	}
}

func main() {
	var g = f()

	fmt.Print(g(1), "-")
	fmt.Print(g(20), "-")
	fmt.Print(g(300))
	fmt.Println()
	var k = f()

	fmt.Print(k(4), "-")
	fmt.Print(k(40), "-")
	fmt.Print(k(400))
}
