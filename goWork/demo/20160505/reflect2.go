package main

import (
	"fmt"
	"reflect"
)

func main() {
	var x float64 = 3.14
	v := reflect.ValueOf(x)

	fmt.Println("setability of v ", v.CanSet())

	v = reflect.ValueOf(&x)

	fmt.Println("type of v is ", v.Type())
	fmt.Println("setability of v ", v.CanSet())

	v = v.Elem()

	v.SetFloat(3.1415)
	fmt.Println(v.Interface())
	fmt.Println(v)

}
