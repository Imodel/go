package main

import (
	"fmt"
)

type B struct {
	thing int
}

func (b *B) change() {
	b.thing = 1

}

func (b B) change2() {
	b.thing = 222
}

func (b B) write() {
	fmt.Print(b.thing)
}
func main() {
	var b B
	b.thing = 3
	b.change2()
	b.write()

	d := B{2}
	d.write()
	d.change()
	d.write()

}
