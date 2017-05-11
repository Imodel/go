package main

import (
	"./pack1/pack1"
	//"./testPack"
	"fmt"
)

func main() {
	var test1 string
	test1 = pack1.ReturnStr()
	fmt.Printf("pack1 func return string %s\n", test1)
	fmt.Printf("pack1 num var int %d", pack1.Pack1Int)

	//test1 = testPack.ReturnStr()
	//fmt.Print(test1)
}
