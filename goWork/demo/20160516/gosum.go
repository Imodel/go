package main

import "fmt"

func sum(x, y int, ch chan int) {
	ch <- x + y
}

func main() {
	ch := make(chan int)

	go sum(12, 13, ch)
	fmt.Println(<-ch)
}
