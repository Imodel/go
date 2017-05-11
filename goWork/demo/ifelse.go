package main

import (
	"fmt"
)

func main() {

	var first int = 10
	//var cond int

	if first < 0 {
		fmt.Printf("this first less than 10 \n")
	} else {
		fmt.Printf("this first more than 0\n")
	}

	if first := 1; first > 5 {
		fmt.Printf("重新定义了first \n")
	} else if first := 5; first < 2 {
		fmt.Printf("first iii %d\n", first)
	} else {
		fmt.Printf("first ii %d\n", first)
	}

	fmt.Printf("first value is %d", first)

}
