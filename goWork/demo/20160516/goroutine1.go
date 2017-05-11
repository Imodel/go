package main

import "fmt"
import "time"

func longWait() {
	fmt.Println("begin longWait")
	time.Sleep(10 * time.Second)
	fmt.Println("end longWait")
}

func shotWait() {
	fmt.Println("begin shotWait")
	time.Sleep(5 * time.Second)
	fmt.Println("end shotWait")
}

func main() {
	fmt.Println("begin main")
	go longWait()
	go shotWait()
	fmt.Println("About to sleep in main")

	time.Sleep(10 * time.Second)
	fmt.Println("at the end of main")
}
