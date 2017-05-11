package main

import (
	"fmt"
)

func main() {

	count := 1

LABEL1:
	for i := 0; i < 9; i++ {
		fmt.Printf("i value is %d \n", i)
		count++
		if count == 3 {
			fmt.Printf("goto>>>>\n")
			goto LABEL1
		}

		if count == 5 {
			fmt.Printf("continue>>>>\n")
			continue LABEL1
		}

		if count == 7 {
			fmt.Printf("break>>>>\n")
			break LABEL1
		}

	}
}
