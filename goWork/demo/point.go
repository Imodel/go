package main 

import (
	"fmt"
)

func main() {
	i := 5
	fmt.Printf("point i is value %d ", i)
	fmt.Println()
	p := &i
	fmt.Printf("%p is %d", p,*p)
}
