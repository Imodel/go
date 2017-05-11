package main

import (
	"fmt"
	"math"
)

type Square struct {
	side float32
}

type Circle struct {
	radius float32
}

type Sharpe interface {
	Area() float32
}

func (c *Circle) Area() float32 {
	return math.Pi * c.radius * c.radius
}

func (s *Square) Area() float32 {
	return s.side * s.side
}

func main() {
	var areaIntf Sharpe
	sql := new(Square)
	sql.side = 5.0

	areaIntf = sql

	if v, ok := areaIntf.(*Square); ok {
		fmt.Printf("this type is Square Area :%T value is %v\n ", v, v.Area())
	} else if t, ok := areaIntf.(*Circle); ok {
		fmt.Printf("this type is Circle Area value is %T\n", t)
	} else {
		fmt.Println("areaIntf does not contain a variable of type Circle")
	}

	circle := new(Circle)
	circle.radius = 3
	areaIntf = circle
	switch t := areaIntf.(type) {
	case *Square:
		fmt.Printf("this type is Square Area :%T value is %v\n ", t, t.Area())
	case *Circle:
		fmt.Printf("this type is Circle Area value is %T,value is %v\n", t, t.Area())
	default:
		fmt.Printf("unknown this type %T", t)
	}
}
