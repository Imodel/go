package main 

import (
	"fmt"
)

func main() {
	var num int = 5

	switch num {
		case 4,5:
			fmt.Printf("%d value is 4 or 5", num)
		case 3:
			fmt.Printf("%d value is 3", num)
		default:
			fmt.Printf("%d value is ?", num)
	}
}

