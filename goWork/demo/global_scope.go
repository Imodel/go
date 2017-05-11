package main

var a = "G"

var (
	k = 1
	t = 2
)

var p, q float32 = 1.1, 2.2

func main() {
	//var a string
	k, t = 1, 2
	a := "t"
	b := "q"
	n()
	m()
	n()
	print(a, b)
	print(k, t, q, p)
}

func n() {
	print(a)
}

func m() {
	a = "O"
	print(a)
}
