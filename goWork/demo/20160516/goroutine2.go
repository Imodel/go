package main

import "fmt"

import "time"

func SendData(ch chan string) {
	ch <- "xk"
	ch <- "sj"
	ch <- "sl"
	ch <- "wx"
}

func getData(ch chan string) {
	var input string
	for {
		input = <-ch
		fmt.Printf("%s \n", input)
	}
}

func main() {
	ch := make(chan string)
	go SendData(ch)
	go getData(ch)
	time.Sleep(1 * time.Second)
}
