package main 

import (
	"fmt"
)


func main() {
	  x := min(1, 3, 2, 0)
	  fmt.Printf(" min value is %d", x )

	  t := 1

	  if true{
	  	t =2
	  }

	  t = 3
	  fmt.Printf("t value is %d", t)
}


func min(a ...int) int{
	if len(a)==0 {
		return 0
	}

	min :=a[0]

	for _,v := range a{
		if v < min {
			min = v
		}
	} 

	return min 
}