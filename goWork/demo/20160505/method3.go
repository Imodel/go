package main

import (
	"fmt"
	"math"
)

type point struct {
	x, y float64
}

type namePoint struct {
	point
	name string
}

func (p *point) Abs() float64 {
	return math.Sqrt(p.x*p.x + p.y*p.y)
}

func main() {
	p := &namePoint{point{3, 4}, "xxxx"}
	fmt.Print(p.Abs())
	fmt.Print(p)
}
