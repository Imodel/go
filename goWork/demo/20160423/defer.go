package main

import "fmt"

func main() {
	function1()
	for i := 1; i < 5; i++ {
		defer fmt.Print(i)
	}
}

func function1() {
	fmt.Printf("In function1 at the top\n")
	defer function2()
	fmt.Printf("In function1 at the bottom!\n")
}

func function2() {
	fmt.Printf("function2: Deferred until the end of the calling function!")
}
