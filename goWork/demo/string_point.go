package main 

import (
	"fmt"
)

func main() {
	s := "week"
	var p *string
	p = &s
	fmt.Printf("s => %s ", s)
	fmt.Printf("%p is %s ", p ,*p)
	*p = "go"

	fmt.Printf("s => %s p => %s",s,*p ) 
}
