package main

import (
	"fmt"
)

type Person struct {
	Name string
	Age  int
}

func newPerson(name string, age int) *Person {
	if age < 0 {
		return nil
	}

	return &Person{name, age}
}

func main() {
	p := newPerson("xk", 20)
	per := newPerson("tt", -1)
	fmt.Print(p, per)
}
