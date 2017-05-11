package main 

import (
	"fmt"
)

func main() {
		a :=1
		goto target
		b :=9
	target:
		b = b+a
		fmt.Printf("value is %v", b)

}
