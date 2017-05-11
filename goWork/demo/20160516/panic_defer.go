package main

import "fmt"

func f() {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("recover in f ", r)
		}
	}()

	fmt.Println("calling func g.")
	g(0)
	fmt.Println("Returned normally from g.")
}

func g(a int) {
	if a > 3 {
		fmt.Println("panicing ")
		panic(fmt.Sprintf("%v", a))
	}

	defer fmt.Println("Defer in g ", a)
	fmt.Println("pring in g ", a)
	g(a + 1)
}

func main() {
	f()
}
