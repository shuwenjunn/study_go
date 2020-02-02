package main

import "fmt"

func main() {
	//	goto + label 跳出多层for 循环

	for i := 0; i < 10; i++ {
		for j := 'A'; j < 'Z'; j++ {
			fmt.Println(j)
			if j == 'B' {
				goto SS
			}
		}
	}

SS:
	fmt.Println("over")
}
