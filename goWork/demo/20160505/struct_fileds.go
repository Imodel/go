package main

import (
	"fmt"
)

type struct1 struct {
	Name        string
	age, iphone int
}

func main() {
	s := new(struct1)
	s.Name = "test"
	s.age = 18
	s.iphone = 132133445

	fmt.Printf("name is %s, age is %d,iphone is %d\n", s.Name, s.age, s.iphone)

	ms := &struct1{"xk", 7, 182}

	fmt.Printf("name is %s, age is %d,iphone is %d\n", ms.Name, ms.age, ms.iphone)

	var ts struct1
	ts = struct1{"kkk", 8, 135}

	fmt.Printf("name is %s, age is %d,iphone is %d\n", ts.Name, ts.age, ts.iphone)

	ps := struct1{Name: "texk", iphone: 131, age: 9}

	fmt.Printf("name is %s, age is %d,iphone is %d\n", ps.Name, ps.age, ps.iphone)

}
