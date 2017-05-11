package main

import (
	"fmt"
)

func main() {
	sm := make([]map[int]string, 5)

	for _, v := range sm {
		v = make(map[int]string)
		v[1] = "ok"
		fmt.Print(v)
	}
	fmt.Println()
	fmt.Println(sm)

	for i := range sm {
		sm[i] = make(map[int]string)
		sm[i][1] = "ok"
		fmt.Print(sm[i])
	}
	fmt.Println()
	fmt.Println(sm)
}
