package main

import (
	ft "fmt"
)

var t = 5

type T struct{}

const (
	a int = 1
	b
	c int = iota
	d
	k = "8"
)

const (
	e, f, g = 5, "6", 7
)

const (
	h = "10"
)

func init() { // initialization of package
}

func main() {
	ft.Println(a, b, c, d, e, f, g, k, h)
}
