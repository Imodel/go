package main

import "fmt"
import "reflect"

type TagType struct {
	field1 string "tag field1 string"
	field2 bool   "tag field2 bool"
	field3 int    "tag field3 int"
}

func tagFlect(tt TagType, ix int) {
	tttype := reflect.TypeOf(tt)
	ixField := tttype.Field(ix)
	fmt.Printf("%v\n", ixField.Tag)
}

func main() {
	tt := TagType{"string", true, 8}
	for i := 0; i < 3; i++ {
		tagFlect(tt, i)
	}
}
