package main

import (
	"fmt"
)

func main() {
	map1 := make(map[int]float32)
	map1[1] = 1.1
	map1[2] = 2.2
	map1[3] = 3.3
	map1[4] = 4.4

	for key, value := range map1 {
		fmt.Printf("key is %d  value is %f\n", key, value)
	}

	for key := range map1 {
		fmt.Printf("key is %d ", key)
	}

	fmt.Println("")

	for _, value := range map1 {
		fmt.Printf("value is %f ", value)
	}

}
