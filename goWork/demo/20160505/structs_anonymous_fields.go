package main

import "fmt"

type inners struct {
	in1 int
	in2 int
}

type outers struct {
	ou1  float32
	out2 string
	int
	inners
}

func main() {
	out := outers{5.5, "out", 7, inners{1, 2}}

	fmt.Print(out.ou1, out.out2, out.int, out.in1, out.in2)

	out1 := new(outers)
	out1.ou1 = 6.6
	out1.out2 = "ddddd"
	out1.int = 8
	out1.in1 = 1
	out1.in2 = 2
	fmt.Print(out1)
}
