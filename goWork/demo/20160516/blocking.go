package main

import "fmt"

func f1(ch chan int) {
	fmt.Println(<-ch)
}

func main() {
	ch := make(chan int)
	ch <- 2
	go f1(ch)

}
