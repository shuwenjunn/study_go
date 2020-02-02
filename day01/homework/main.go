package main

import "fmt"

func main() {
	for i := 9; i > 0; i-- {
		for j := 1; j < 10; j++ {
			// fmt.Println(i * j)
			if j <= i {
				fmt.Printf("%d*%d=%d\t ", i, j, i*j)
			}

		}
		fmt.Println()
	}

	s := "string"
	s="134"
	fmt.Println(s)
}
