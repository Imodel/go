package main

import (
	"fmt"
)

func Add2() func(b int) int {
	return func(b int) int {
		return b + 2
	}
}

func Add3(a int) func(b int) int {
	fmt.Println(a)
	return func(b int) int {
		return a + b
	}
}

func main() {
	p2 := Add2()
	fmt.Printf("call Add2 for 3 %v\n", p2(3))

	p3 := Add3(3)
	fmt.Printf("call Add3 for 3 %v\n", p3(2))
}
