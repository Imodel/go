package main

import (
	"fmt"
)

type Person struct {
	Name string
	Age  int
}

func main() {
	a := Person{
		Name: "joe",
		Age:  19,
	}

	fmt.Print(a, "\n")
	A(&a)

	fmt.Print(a, "\n")
}

func A(per *Person) {
	per.Age = 13

	fmt.Print("A>>>>", per, "\n")
}
