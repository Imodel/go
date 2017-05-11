package main

import (
	"./martix"
	"fmt"
)

func main() {
	p1 := martix.NewPerson("xk", 20)
	p2 := martix.NewPerson("xk", 20)
	fmt.Print(&p1, &p2)
}
