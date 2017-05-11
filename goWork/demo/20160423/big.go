package main

import (
	"fmt"
	"math"
	"math/big"
)

func main() {
	im := big.NewInt(math.MaxInt64)
	fmt.Print(im, "\n")
	in := im
	io := big.NewInt(1956)
	fmt.Print(io, "\n")
	ip := big.NewInt(1)
	fmt.Print(ip, "\n")

	ip.Mul(im, in).Add(ip, im).Div(ip, io)
	fmt.Printf("big int :%v\n", ip)

	rm := big.NewRat(math.MaxInt64, 1956)
	rn := big.NewRat(-1956, math.MaxInt64)
	ro := big.NewRat(19, 56)
	rp := big.NewRat(1111, 2222)
	rq := big.NewRat(1, 1)
	rq.Mul(rm, rn).Add(rq, ro).Mul(rq, rp)
	fmt.Printf("Big Rat: %v\n", rq)
}
