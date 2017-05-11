package main

import (
	"fmt"
)

type Shaper interface {
	Area() float32
}

type Square struct {
	side float32
}

func (s *Square) Area() float32 {
	return s.side * s.side

}

func main() {
	sql := new(Square)
	sql.side = 1.1

	var areaIntf Shaper
	areaIntf = sql

	fmt.Print(areaIntf.Area())
}
