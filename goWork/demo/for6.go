package main 

import (
	"fmt"
)

func main() {
	LABEL1:
		for i:=1;i<5;i++{
			for j:=1;j<=5;j++{
				if j == 4 {
					continue LABEL1;
				}

				fmt.Printf("i value is %d , j value is %d \n", i,j)

			}
		}
}
