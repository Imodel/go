package main

import (
	"fmt"
)

func main() {
	callback(1, 2, add)
}
func add(a, b int) {
	fmt.Printf("a + b = %d", a+b)
}

func callback(x, y int, f func(int, int)) {
	f(x, y)

}
