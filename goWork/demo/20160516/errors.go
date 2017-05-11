package main

import "errors"
import "fmt"

var ErrorsNew error = errors.New("errors new")

func main() {
	fmt.Println("%v", ErrorsNew)
}
