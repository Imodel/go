package main 

import (
	"fmt"
	"strconv"
)


func main() {
	var oring string ="5"
	var newS  string

	fmt.Printf("the int size is %d\n", strconv.IntSize)

	an,err := strconv.Atoi(oring)

	if err!=nil {
		fmt.Printf("%s is not int \n", oring)
		return
	}

	fmt.Printf("integer is %d \n", an)
	an = an+5

	newS = strconv.Itoa(an)

	fmt.Printf(" an value is %s", newS)
}