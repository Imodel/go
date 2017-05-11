package main

import "fmt"

var i = 5
var str = "ABC"

type Person struct {
	name string
	age  int
}

type Any interface{}

func main() {
	var val Any
	val = 5
	fmt.Printf("val has the value:%v\n", val)
	val = str
	fmt.Printf("val has the value:%v\n", val)
	per1 := new(Person)
	per1.age = 15
	per1.name = "sss"

	val = per1
	fmt.Printf("val has the value:%v\n", val)

	switch val.(type) {
	case int:
		fmt.Print("type is int")
	case string:
		fmt.Print("type is string")
	case *Person:
		fmt.Print("type is Person")
	default:
		fmt.Print("type is unkown")
	}
}
