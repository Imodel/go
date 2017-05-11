package main

import "fmt"

type Ducker interface {
	duce()
	walk()
}

type Bird struct{}

func Duck(du Ducker) {
	for i := 0; i < 3; i++ {
		du.duce()
		du.walk()
	}
}

func (b *Bird) duce() {
	fmt.Println("Bird duce")
}

func (b *Bird) walk() {
	fmt.Println("Bird walk")
}

func main() {
	b := new(Bird)
	Duck(b)
}
