package main

import (
	"fmt"
)

type number struct {
	i float32
}

type nr number

func main() {
	a := number{5.5}
	b := nr{5.5}
	var c = number(b)

	fmt.Print(a, b, c)

}
