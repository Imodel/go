package main

import (
	"fmt"
)

var t int = 7
func Multiply(a, b int, replay *int) {

	*replay = a * b
}

func main() {
	n := 0
	replay := &n
	fmt.Printf("t value is ", t)
	Multiply(3, 4, replay)

	fmt.Printf("replay value is %d", *replay)

}
