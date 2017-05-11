package main

import "fmt"
import "time"

func main() {
	ch := make(chan int, 10)
	go func() {
		time.Sleep(10 * time.Second)
		x := <-ch
		fmt.Println("received ", x)
	}()
	fmt.Println("sending ", 10)
	ch <- 10
	fmt.Println("send ", 10)
}
