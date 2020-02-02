package main

import "fmt"

func main() {
	var a = make([]int, 5, 10)

	for i := 0; i < 10; i++ {
		a = append(a, i)
	}

	fmt.Println(a) //[0 0 0 0 0 0 1 2 3 4 5 6 7 8 9]

}
