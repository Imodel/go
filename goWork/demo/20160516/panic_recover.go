package main

import "fmt"

func badCall() {
	panic("badCall panic\n")
}

func test() {
	defer func() {
		if e := recover(); e != nil {
			fmt.Printf("defer panicing %s\n", e)
		}
	}()

	badCall()
	fmt.Print("after badCall\n")
}

func main() {
	fmt.Print("panicing testing\n")
	test()
	fmt.Print("panicing end\n")
}
