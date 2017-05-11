package main

import (
	"fmt"
)

func main() {
	var m map[int]map[int]string

	m = make(map[int]map[int]string)

	a, ok := m[1][1]
	if !ok {
		m[1] = make(map[int]string)
	}

	m[1][1] = "ok"
	a, ok = m[1][1]
	fmt.Println(a, ok)
}
