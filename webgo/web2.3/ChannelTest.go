package main

import (
	"fmt"
)

func sum(a []int, c chan int) {
	sum := 0
	for _, value := range a {
		sum += value
	}
	c <- sum
}

func main() {
	a := []int{1, 2, 3, -9, 5, 6}
	c := make(chan int)
	go sum(a[:len(a)/2], c)
	go sum(a[len(a)/2:], c)
	x, y := <-c, <-c
	fmt.Println(x, y, x+y)
}
