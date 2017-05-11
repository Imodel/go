package main

import (
	"./sort"
	"fmt"
)

func ints() {
	data := []int{74, 59, 238, -784, 9845, 959, 905, 0, 0, 42, 7586, -5467984, 7586}
	a := sort.IntArr(data) //conversion to type IntArray
	sort.Sort(a)
	if !sort.IsSorted(a) {
		panic("fail")
	}
	fmt.Printf("The sorted array is: %v\n", a)
}

func main() {
	ints()
}
