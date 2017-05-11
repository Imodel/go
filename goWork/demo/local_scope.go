package main

import (
	"fmt"
)

var a = "G"

const (
	b = 1
	c
	d
)

func main() {
	n()
	m()
	n()
	fmt.Print(b, c, d)

}

func n() {
	print(a)
}

func m() {

	var a string
	a = "k"
	print(a)
}
