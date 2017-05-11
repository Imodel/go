package main

import (
	"fmt"
)

func main() {
	result := 0
	for i := 0; i < 10; i++ {
		result = fibncci(i)
		fmt.Printf("fibncci(%d) is: %d\n", i, result)
	}
}
func fibncci(n int) (r int) {
	if n <= 1 {
		return 1 //r=1
	} else {
		return fibncci(n-1) + fibncci(n-2) //r = fibncci(n-1) + fibncci(n-2)效果相同
	}

	return
}
