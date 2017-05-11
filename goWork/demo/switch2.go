package main 

import (
	"fmt"
)

func main() {
	num := 7

	switch {
		case num >8:
			fmt.Printf("num more than 8")
		case num == 7:
			fmt.Printf("num equals 7")
		default:
			fmt.Printf("num is ?")
	}
}
