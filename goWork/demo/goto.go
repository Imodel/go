package main 

func main() {
	i :=1
	HERE:
		print(i)
		i++
		if i ==5 {
			return
		}
		goto HERE
}
