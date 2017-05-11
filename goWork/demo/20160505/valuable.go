package main

import (
	"fmt"
)

type stockPosition struct {
	ticker     string
	sharePrice float32
	count      float32
}

func (s stockPosition) getValue() float32 {
	return s.sharePrice * s.count
}

type car struct {
	make  string
	model string
	price float32
}

func (c car) getValue() float32 {
	return c.price

}

type valuable interface {
	getValue() float32
}

func showValue(asser valuable) {
	fmt.Printf("interface getValue %f\n", asser.getValue())
}

func main() {
	var o valuable = stockPosition{"stock", 1.1, 5}
	showValue(o)

	o = car{"beijing", "ww", 12.5}
	showValue(o)
}
